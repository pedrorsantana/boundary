package secret

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/hashicorp/boundary/internal/secret/store"
)

// Clonable provides a cloning interface
type Cloneable interface {
	Clone() interface{}
}

// Repository is the target database repository
type Repository struct {
	reader db.Reader
	writer db.Writer
	kms    *kms.Kms

	// defaultLimit provides a default for limiting the number of results returned from the repo
	defaultLimit int
}

// allocSecretView will allocate a secret view
func allocSecretView() secrets {
	return secrets{
		Secret: &store.Secret{},
	}
}

func allocSecretSessionView() secret_session {
	return secret_session{
		SecretSession: &store.SecretSession{},
	}
}

func allocSecretSessionLogsView() secret_session_log {
	return secret_session_log{
		SecretSessionLog: &store.SecretSessionLog{},
	}
}

// NewRepository creates a new secret Repository. Supports the options: WithLimit
// which sets a default limit on results returned by repo operations.
func NewRepository(r db.Reader, w db.Writer, kms *kms.Kms) (*Repository, error) {
	const op = "secret.NewRepository"
	if r == nil {
		return nil, errors.New(errors.InvalidParameter, op, "nil reader")
	}
	if w == nil {
		return nil, errors.New(errors.InvalidParameter, op, "nil writer")
	}
	if kms == nil {
		return nil, errors.New(errors.InvalidParameter, op, "nil kms")
	}

	return &Repository{
		reader: r,
		writer: w,
		kms:    kms,
	}, nil
}

func (r *Repository) ListSecrets(ctx context.Context, opt ...Option) ([]Secret, error) {
	const op = "secret.(Repository).ListSecrets"
	opts := getOpts(opt...)

	var args []interface{}
	var where []string
	if len(opts.withScopeId) != 0 {
		where, args = append(where, "scope_id in (?)"), append(args, opts.withScopeId)
	}
	if len(opts.withPublicId) != 0 {
		where, args = append(where, "public_id in (?)"), append(args, opts.withPublicId)
	}

	var foundSecrets []*secrets
	err := r.list(ctx, &foundSecrets, strings.Join(where, " and "), args, opt...)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	secretsList := make([]Secret, 0, len(foundSecrets))
	for _, sct := range foundSecrets {
		data, err := base64.StdEncoding.DecodeString(sct.GetCode())
		if err != nil {
			return nil, err
		}
		sct.Code = string(data)
		secretsList = append(secretsList, sct)
	}

	return secretsList, nil
}

func (r *Repository) list(ctx context.Context, resources interface{}, where string, args []interface{}, opt ...Option) error {
	const op = "secret.(Repository).list"
	opts := getOpts(opt...)
	limit := r.defaultLimit
	var dbOpts []db.Option
	if opts.withLimit != 0 {
		// non-zero signals an override of the default limit for the repo.
		limit = opts.withLimit
	}
	dbOpts = append(dbOpts, db.WithLimit(limit))

	if len(opts.withOrder) > 0 {
		dbOpts = append(dbOpts, db.WithOrder(opts.withOrder))
	}

	if err := r.reader.SearchWhere(ctx, resources, where, args, dbOpts...); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (r *Repository) CreateSecretInRepo(ctx context.Context, name string, manager string, publicId string, scopeId string, code string, opt ...Option) (*secrets, error) {
	const op = "secret.(Repository).CreateSecret"

	//databaseWrapper, err := r.kms.GetWrapper(ctx, secret.GetScopeId(), kms.KeyPurposeDatabase)
	//if err != nil {
	// 	return nil, errors.Wrap(err, op, errors.WithMsg("unable to get database wrapper"))
	// }

	secret := allocSecretView()
	secret.Name = name
	secret.ScopeId = scopeId
	secret.Code = code
	secret.Manager = manager
	secret.tableName = "secrets"

	var newSecret *secrets
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(read db.Reader, w db.Writer) error {
			newSecret = secret.Clone()
			newSecret.tableName = "secrets"
			id, err := newSecretId()
			if err != nil {
				return err
			}
			newSecret.PublicId = id

			if err := w.Create(ctx, newSecret); err != nil {
				return err
			}
			return nil
		},
	)

	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return newSecret, nil
}

func (r *Repository) DeleteSecret(ctx context.Context, publicId string, opt ...Option) (int, error) {
	const op = "secret.(Repository).DeleteSecret"
	opts := getOpts(opt...)
	if publicId == "" {
		return db.NoRowsAffected, errors.New(errors.InvalidParameter, op, "missing public id")
	}

	deleteSecret := allocSecretView()
	deleteSecret.PublicId = publicId
	deleteSecret.ScopeId = opts.withScopeId

	var rowsDeleted int
	var deleteResource *secrets
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(_ db.Reader, w db.Writer) error {
			deleteResource = deleteSecret.Clone()
			deleteResource.tableName = "secrets"
			rowsDeleted, err := w.Delete(
				ctx,
				deleteResource,
				db.WithWhere("public_id = ?", deleteResource.GetPublicId()),
			)
			if err != nil {
				return errors.Wrap(err, op)
			}
			if rowsDeleted > 1 {
				// return err, which will result in a rollback of the delete
				return errors.New(errors.MultipleRecords, op, "more than 1 resource would have been deleted")
			}
			return nil
		},
	)
	if err != nil {
		return db.NoRowsAffected, errors.Wrap(err, op)
	}
	return rowsDeleted, nil
}

func (r *Repository) CreateSecretRequestInRepo(ctx context.Context, opt ...Option) (*secret_session, error) {
	const op = "secret.(Repository).CreateSecretRequestInRepo"
	opts := getOpts(opt...)

	sess := allocSecretSessionView()
	sess.ScopeId = opts.withScopeId
	sess.Description = opts.withDescription
	sess.UserId = opts.withUserId
	sess.SecretId = opts.withSecretId
	sess.RequestedTime = opts.withRequestedTime

	var newSecretSession *secret_session
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(read db.Reader, w db.Writer) error {
			newSecretSession = sess.Clone()
			newSecretSession.tablename = "secret_sessions"
			id, err := newSecretSessionId()
			if err != nil {
				return err
			}
			newSecretSession.PublicId = id

			if err := w.Create(ctx, newSecretSession); err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return newSecretSession, nil
}

func (r *Repository) SetSecretRequestLogInRepo(ctx context.Context, opt ...Option) (*secret_session_log, error) {
	const op = "secret.(Repository).SetSecretRequestLogInRepo"
	opts := getOpts(opt...)

	logs := allocSecretSessionLogsView()
	logs.Output = opts.withOutput
	logs.SecretSessionId = opts.withSecretSessionId

	var newSecretSessionLog *secret_session_log
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(read db.Reader, w db.Writer) error {
			newSecretSessionLog = logs.Clone()
			newSecretSessionLog.tablename = "secret_session_logs"
			id, err := newSecretSessionLogId()
			if err != nil {
				return err
			}
			newSecretSessionLog.PublicId = id

			if err := w.Create(ctx, newSecretSessionLog); err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return newSecretSessionLog, nil
}

func (r *Repository) LookupSecretSession(ctx context.Context, secret_id string, user_id string, opt ...Option) (SecretSession, error) {
	const op = "secret.(Repository).LookupSecretSession"

	var args []interface{}
	var where []string

	where, args = append(where, "secret_id in (?)"), append(args, secret_id)
	where, args = append(where, "user_id in (?)"), append(args, user_id)

	var foundLogs []*secret_session
	err := r.list(ctx, &foundLogs, strings.Join(where, " and "), args, opt...)
	if err != nil {
		return nil, err
	}

	if len(foundLogs) == 0 {
		return nil, nil
	}

	return foundLogs[0].SecretSession, nil
}

func (r *Repository) LookupSecretSessionLogs(ctx context.Context, secret_session_id string, opt ...Option) (SecretSessionLog, error) {
	const op = "secret.(Repository).LookupSecretSessionLogs"
	var args []interface{}
	var where []string

	where, args = append(where, "secret_session_id in (?)"), append(args, secret_session_id)

	var foundLogs []*secret_session_log
	err := r.list(ctx, &foundLogs, strings.Join(where, " and "), args, opt...)
	if err != nil {
		return nil, err
	}

	if len(foundLogs) == 0 {
		return nil, nil
	}

	return foundLogs[0].SecretSessionLog, nil
}

func (r *Repository) ListSecretSessions(ctx context.Context, secret_id string, opt ...Option) ([]SecretSession, error) {
	const op = "secret.(Repository).ListSecretSessions"
	opts := getOpts(opt...)

	if len(opts.withScopeId) == 0 {
		return nil, fmt.Errorf("scope_id is nil.")
	}

	if len(secret_id) == 0 {
		return nil, fmt.Errorf("secret_id is nil.")
	}

	var args []interface{}
	var where []string

	where, args = append(where, "scope_id in (?)"), append(args, opts.withScopeId)
	where, args = append(where, "secret_id in (?)"), append(args, secret_id)

	var foundSessions []*secret_session
	err := r.list(ctx, &foundSessions, strings.Join(where, " and "), args, opt...)
	if err != nil {
		return nil, err
	}

	sessList := make([]SecretSession, 0, len(foundSessions))
	for _, sct := range foundSessions {
		sessList = append(sessList, sct)
	}

	return sessList, nil
}
