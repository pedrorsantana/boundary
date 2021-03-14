package secret

import (
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/db/timestamp"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/secret/store"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultSecretTableName = "secrets"
)

type secrets struct {
	*store.Secret
	tableName string `gorm:"-"`
}

type secret_session struct {
	*store.SecretSession
	tablename string `gorm:"-"`
}

type secret_session_log struct {
	*store.SecretSessionLog
	tablename string `gorm:"-"`
}

// Secret is a commmon interface
type Secret interface {
	GetPublicId() string
	GetScopeId() string
	GetCode() string
	GetName() string
	GetManager() string
	GetCreateTime() *timestamp.Timestamp
}

// SecretSession is a commmon interface
type SecretSession interface {
	GetPublicId() string
	GetScopeId() string
	GetUserId() string
	GetSecretId() string
	GetDescription() string
	GetRequestedTime() uint32
	GetCreateTime() *timestamp.Timestamp
	GetUpdatedTime() *timestamp.Timestamp
}

// SecretSession is a commmon interface
type SecretSessionLog interface {
	GetPublicId() string
	GetOutput() string
	GetCreateTime() *timestamp.Timestamp
}

// Clone creates a clone of the Secret
func (s *secrets) Clone() *secrets {
	cp := proto.Clone(s.Secret)
	return &secrets{
		Secret: cp.(*store.Secret),
	}
}

// Clone creates a clone of the Secret
func (s *secret_session) Clone() *secret_session {
	cp := proto.Clone(s.SecretSession)
	return &secret_session{
		SecretSession: cp.(*store.SecretSession),
	}
}

// Clone creates a clone of the Secret
func (s *secret_session_log) Clone() *secret_session_log {
	cp := proto.Clone(s.SecretSessionLog)
	return &secret_session_log{
		SecretSessionLog: cp.(*store.SecretSessionLog),
	}
}

func newSecretId() (string, error) {
	const op = "authtoken.newSecretId"
	id, err := db.NewPublicId("srct")
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return id, nil
}

func newSecretSessionId() (string, error) {
	const op = "authtoken.newSecretId"
	id, err := db.NewPublicId("srss")
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return id, nil
}

func newSecretSessionLogId() (string, error) {
	const op = "authtoken.newSecretId"
	id, err := db.NewPublicId("sssl")
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return id, nil
}
