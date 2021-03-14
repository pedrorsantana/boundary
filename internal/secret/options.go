package secret

import "github.com/hashicorp/boundary/internal/db"

// getOpts - iterate the inbound Options and return a struct
func getOpts(opt ...Option) options {
	opts := getDefaultOptions()
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Option - how Options are passed as arguments
type Option func(*options)

// options = how options are represented
type options struct {
	withName            string
	withDescription     string
	withCode            string
	withLimit           int
	withScopeId         string
	withScopeIds        []string
	withScopeName       string
	withUserId          string
	withAuthTokenId     string
	withPublicId        string
	withSecretId        string
	withSecretSessionId string
	withOutput          string
	withOrder           string
	withVersion         uint32
	withRequestedTime   uint32
}

func getDefaultOptions() options {
	return options{
		withName:            "",
		withDescription:     "",
		withLimit:           db.DefaultLimit,
		withCode:            "",
		withScopeId:         "",
		withScopeIds:        nil,
		withScopeName:       "",
		withUserId:          "",
		withPublicId:        "",
		withAuthTokenId:     "",
		withSecretId:        "",
		withSecretSessionId: "",
		withOutput:          "",
		withOrder:           "",
		withVersion:         0,
	}
}

// WithDescription provides an optional description
func WithDescription(desc string) Option {
	return func(o *options) {
		o.withDescription = desc
	}
}

// WithName provides an option to search by a friendly name
func WithName(name string) Option {
	return func(o *options) {
		o.withName = name
	}
}

// WithLimit provides an option to provide a limit.  Intentionally allowing
// negative integers.   If WithLimit < 0, then unlimited results are returned.
// If WithLimit == 0, then default limits are used for results.
func WithLimit(limit int) Option {
	return func(o *options) {
		o.withLimit = limit
	}
}

// WithCode provides an option to specify the default target port.
func WithCode(c string) Option {
	return func(o *options) {
		o.withCode = c
	}
}

// WithScopeId provides an option to search by a scope id
func WithScopeId(scopeId string) Option {
	return func(o *options) {
		o.withScopeId = scopeId
	}
}

// WithScopeId provides an option to search by multiple scope id
func WithScopeIds(scopeIds []string) Option {
	return func(o *options) {
		o.withScopeIds = scopeIds
	}
}

// WithScopeId provides an option to search by a scope name
func WithScopeName(scopeName string) Option {
	return func(o *options) {
		o.withScopeName = scopeName
	}
}

// WithUserId provides an option to search by a user public id
func WithUserId(userId string) Option {
	return func(o *options) {
		o.withUserId = userId
	}
}

// WithPublicId provides an optional public id
func WithPublicId(id string) Option {
	return func(o *options) {
		o.withPublicId = id
	}
}

// withAuthTokenId provides an auth_token id
func withAuthTokenId(authTokenId string) Option {
	return func(o *options) {
		o.withAuthTokenId = authTokenId
	}
}

// WithSecretId provides an secret id
func WithSecretId(secretId string) Option {
	return func(o *options) {
		o.withSecretId = secretId
	}
}

// WithSecretSessionId provides an secret id
func WithSecretSessionId(secretSessionId string) Option {
	return func(o *options) {
		o.withSecretSessionId = secretSessionId
	}
}

// withOutput provides an output from secretsession logs
func WithOutput(output string) Option {
	return func(o *options) {
		o.withOutput = output
	}
}

// WithVersion provides an version
func WithVersion(version uint32) Option {
	return func(o *options) {
		o.withVersion = version
	}
}

// WithVersion provides an version
func WithRequestedTime(requested_time uint32) Option {
	return func(o *options) {
		o.withRequestedTime = requested_time
	}
}

// WithOrder provides a order by query
func WithOrder(order string) Option {
	return func(o *options) {
		o.withOrder = order
	}
}
