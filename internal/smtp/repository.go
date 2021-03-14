package smtp

import (
	"context"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/kms"
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

// NewRepository creates a new secret Repository. Supports the options: WithLimit
// which sets a default limit on results returned by repo operations.
func NewRepository(r db.Reader, w db.Writer, kms *kms.Kms) (*Repository, error) {
	const op = "smtp.NewRepository"
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

func LookupSmtpServer(ctx context.Context, scopeId string) {
	const op = "smtp.(Repository).LookupSmtpServer"

	var args []interface{}
	var where []string
	if len(scopeId) != 0 {
		where, args = append(where, "scope_id in (?)"), append(args, scopeId)
	}

}
