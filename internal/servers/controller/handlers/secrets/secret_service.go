package secrets

import (
	"context"
	"encoding/base64"
	stderrors "errors"
	"fmt"
	"time"

	"github.com/hashicorp/boundary/internal/auth"
	"github.com/hashicorp/boundary/internal/gen/controller/api/resources/secrets"
	pbs "github.com/hashicorp/boundary/internal/gen/controller/api/resources/secrets"
	pbsvc "github.com/hashicorp/boundary/internal/gen/controller/api/services"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/hashicorp/boundary/internal/secret"
	"github.com/hashicorp/boundary/internal/servers/controller/common"
	"github.com/hashicorp/boundary/internal/servers/controller/handlers"
	"github.com/hashicorp/boundary/internal/servers/controller/handlers/accounts"
	docker "github.com/hashicorp/boundary/internal/servers/worker/docker"
	smtp "github.com/hashicorp/boundary/internal/smtp"
	"github.com/hashicorp/boundary/internal/types/action"
	"github.com/hashicorp/boundary/internal/types/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	maskManager handlers.MaskManager

	// IdActions contains the set of actions that can be performed on
	// individual resources
	IdActions = action.ActionSet{
		action.Read,
		action.Update,
		action.Delete,
		action.Authenticate,
	}

	// CollectionActions contains the set of actions that can be performed on
	// this collection
	CollectionActions = action.ActionSet{
		action.Create,
		action.List,
	}

	collectionTypeMap = map[resource.Type]action.ActionSet{
		resource.Account: accounts.CollectionActions,
	}
)

func init() {
	// var err error
	// if maskManager, err = handlers.NewMaskManager(&store.Secret{}, &pbs.Secret{}); err != nil {
	// 	panic(err)
	// }
}

// Service handles request as described by the pbsvc.AuthMethodServiceServer interface.
type Service struct {
	pbsvc.UnimplementedSecretServiceServer

	repoFn    common.SecretRepoFactory
	iamRepoFn common.IamRepoFactory
	kms       *kms.Kms
}

// NewService returns a auth method service which handles auth method related requests to boundary.
func NewService(kms *kms.Kms, repoFn common.SecretRepoFactory) (Service, error) {
	if kms == nil {
		return Service{}, stderrors.New("nil kms provided")
	}
	if repoFn == nil {
		return Service{}, fmt.Errorf("nil password repository provided")
	}
	return Service{kms: kms, repoFn: repoFn}, nil
}

var _ pbsvc.SecretServiceServer = Service{}

func (s Service) GetSecret(ctx context.Context, req *pbsvc.GetSecretRequest) (*pbsvc.GetSecretResponse, error) {
	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	result, err := repoFn.ListSecrets(ctx, secret.WithPublicId(req.GetId()), secret.WithLimit(1))
	if err != nil {
		return nil, err
	}

	var resp = &pbsvc.GetSecretResponse{}

	if len(result) > 0 {
		o, err := toProto(result[0], nil)
		if err != nil {
			return nil, err
		}

		opts := []auth.Option{auth.WithType(resource.Secret), auth.WithAction(action.Read), auth.WithScopeId(o.ScopeId)}
		var a = auth.Verify(ctx, opts...)
		if a.Error != nil {
			return nil, a.Error
		}

		// Remove Code for non authorized creators
		opts = []auth.Option{auth.WithType(resource.Secret), auth.WithAction(action.Create), auth.WithScopeId(o.ScopeId)}
		var aC = auth.Verify(ctx, opts...)
		if aC.Error == nil {
			o.Code = ""
		}

		resp.Item = o
	}

	return resp, nil
}

func (s Service) CreateSecret(ctx context.Context, req *pbsvc.CreateSecretRequest) (*pbsvc.CreateSecretResponse, error) {
	opts := []auth.Option{auth.WithType(resource.Secret), auth.WithAction(action.Create), auth.WithScopeId(req.GetItem().GetScopeId())}
	var a = auth.Verify(ctx, opts...)
	if a.Error != nil {
		return nil, a.Error
	}

	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	encodedCode := base64.StdEncoding.EncodeToString([]byte(req.GetItem().GetCode()))

	result, err := repoFn.CreateSecretInRepo(ctx, req.GetItem().Name, req.GetItem().Manager, req.GetItem().GetId(), req.GetItem().GetScopeId(), encodedCode)
	if err != nil {
		return nil, err
	}

	respSecret, err := toProto(result, nil)

	var resp = &pbsvc.CreateSecretResponse{}
	resp.Item = respSecret

	return resp, nil
}

func (s Service) DeleteSecret(ctx context.Context, req *pbsvc.DeleteSecretRequest) (*pbsvc.DeleteSecretResponse, error) {
	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	sct, err := repoFn.ListSecrets(ctx, secret.WithPublicId(req.Id), secret.WithLimit(1))
	if err != nil {
		return nil, err
	}
	if len(sct) == 0 {
		return nil, nil
	}

	opts := []auth.Option{auth.WithType(resource.Secret), auth.WithAction(action.Delete), auth.WithScopeId(sct[0].GetScopeId())}
	var a = auth.Verify(ctx, opts...)
	if a.Error != nil {
		return nil, a.Error
	}
	_, err = repoFn.DeleteSecret(ctx, sct[0].GetPublicId(), secret.WithScopeId(sct[0].GetScopeId()))
	if err != nil {
		return nil, err
	}

	return &pbsvc.DeleteSecretResponse{}, nil
}

func (s Service) ListSecrets(ctx context.Context, req *pbsvc.ListSecretsRequest) (*pbsvc.ListSecretsResponse, error) {
	opts := []auth.Option{auth.WithType(resource.Secret), auth.WithAction(action.List), auth.WithScopeId(req.ScopeId)}
	var a = auth.Verify(ctx, opts...)
	if a.Error != nil {
		return nil, a.Error
	}

	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	ul, err := repoFn.ListSecrets(ctx, secret.WithScopeId(req.ScopeId))
	if err != nil {
		return nil, err
	}

	var outUl []*secrets.Secret
	for _, u := range ul {
		o, err := toProto(u, nil)
		if err != nil {
			return nil, handlers.ApiErrorWithCodeAndMessage(codes.Internal, "Unable to convert value to proto: %v.", err)
		}
		outUl = append(outUl, o)
	}

	var resp = &pbsvc.ListSecretsResponse{}
	resp.Items = outUl

	return resp, nil
}

func (s Service) RequestSecretAccess(ctx context.Context, req *pbsvc.RequestSecretAccessRequest) (*pbsvc.RequestSecretAccessResponse, error) {
	opts := []auth.Option{auth.WithType(resource.SecretSession), auth.WithAction(action.Create), auth.WithScopeId(req.GetItem().ScopeId)}
	var a = auth.Verify(ctx, opts...)
	if a.Error != nil {
		return nil, a.Error
	}

	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	if req.GetItem() == nil {
		return nil, stderrors.New("item not found in request.")
	}

	//TODO: Issue de segurança, preciso pegar esse valor de outro lugar
	md, ok := metadata.FromIncomingContext(ctx)
	if ok != true {
		return nil, stderrors.New("Can't find hostname.")
	}

	result, err := repoFn.CreateSecretRequestInRepo(ctx,
		secret.WithDescription(req.GetItem().Description),
		secret.WithSecretId(req.GetItem().SecretId),
		secret.WithScopeId(req.GetItem().ScopeId),
		secret.WithUserId(a.UserId),
		secret.WithRequestedTime(req.GetItem().RequestedTime),
	)
	if err != nil {
		return nil, err
	}

	sct, err := repoFn.ListSecrets(ctx, secret.WithLimit(1), secret.WithRequestedTime(req.GetItem().RequestedTime), secret.WithPublicId(req.GetItem().SecretId))
	if err != nil {
		return nil, err
	}

	if len(sct) == 0 {
		return nil, stderrors.New("error finding secret by id")
	}

	approveURL := fmt.Sprintf("http://%s/v1/secrets/authorize-session?secret_id=%s&id=%s&user_id=%s&hash=%s", md.Get("x-forwarded-host")[0], sct[0].GetPublicId(), result.GetPublicId(), a.UserId, "HASH_R")

	go smtp.SendEmail(sct[0].GetManager(),
		fmt.Sprintf("Request access in %s", sct[0].GetName()),
		fmt.Sprintf("Hello, \n The user %s is requesting access to %s during %d hour(s), if you want approve please click in the link below. \n\n Request Description: %s \n\n Approve URL: %s", a.UserId, sct[0].GetName(), req.GetItem().RequestedTime, result.GetDescription(), approveURL))

	var resp = &pbsvc.RequestSecretAccessResponse{}

	o, err := toProtoSession(result, nil)
	if err != nil {
		return nil, err
	}
	resp.Item = o

	return resp, nil
}

// AuthorizeSecretSession was requested by secret manager for authorize the session.
func (s Service) AuthorizeSecretSession(ctx context.Context, req *pbsvc.AuthorizeSecretSessionRequest) (*pbsvc.AuthorizeSecretSessionResponse, error) {
	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	sct, err := repoFn.ListSecrets(ctx, secret.WithPublicId(req.SecretId), secret.WithLimit(1))
	if err != nil {
		return nil, err
	}

	if len(sct) == 0 {
		return nil, nil
	}

	sctS, err := repoFn.LookupSecretSession(ctx, req.SecretId, req.GetUserId())
	if err != nil {
		return nil, err
	}

	resp := &pbsvc.AuthorizeSecretSessionResponse{}

	maxDate := sctS.GetCreateTime().Timestamp.AsTime().Add(time.Hour * time.Duration(sctS.GetRequestedTime()))
	if maxDate.Before(time.Now()) {
		resp.Status = "Expired"
		return resp, nil
	}

	go func() {
		execResult, err := docker.ReadCodeAndExecute(sct[0].GetCode())
		if err != nil {
			execResult = err.Error()
		}
		encodedExecResult := base64.StdEncoding.EncodeToString([]byte(execResult))
		repoFn.SetSecretRequestLogInRepo(context.Background(), secret.WithSecretSessionId(req.GetId()), secret.WithSecretId(req.SecretId), secret.WithOutput(encodedExecResult))
	}()

	resp.Status = "Executing..."
	return resp, nil
}

func (s Service) ListSecretSessionEvents(ctx context.Context, req *pbsvc.ListSecretSessionEventsRequest) (*pbsvc.ListSecretSessionEventsResponse, error) {
	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}

	found, err := repoFn.ListSecretSessions(ctx, req.Include[1], secret.WithScopeId(req.Include[3]))
	if err != nil {
		return nil, err
	}

	var outFound []*secrets.SecretSession
	for _, u := range found {
		o := &secrets.SecretSession{
			//TODO: Fix it
			Id:          u.GetPublicId(),
			ScopeId:     u.GetScopeId(),
			UserId:      u.GetUserId(),
			Description: "",
			SecretId:    u.GetSecretId(),
			CreatedTime: u.GetCreateTime().GetTimestamp(),
			UpdatedTime: u.GetUpdatedTime().GetTimestamp(),
		}
		outFound = append(outFound, o)
	}

	var resp = &pbsvc.ListSecretSessionEventsResponse{}
	resp.Items = outFound

	return resp, nil
}

func (s Service) GetSecretSession(ctx context.Context, req *pbsvc.GetSecretSessionRequest) (*pbsvc.GetSecretSessionResponse, error) {
	opts := []auth.Option{auth.WithType(resource.SecretSession), auth.WithAction(action.Read), auth.WithScopeId(req.Include)}
	var a = auth.Verify(ctx, opts...)
	if a.Error != nil {
		return nil, a.Error
	}

	repoFn, err := s.repoFn()
	if err != nil {
		return nil, err
	}
	if req.Include == "" {
		return nil, stderrors.New("Argument scope_id not found")
	}

	sctSession, err := repoFn.LookupSecretSession(ctx, req.GetId(), a.UserId)
	if err != nil {
		return nil, err
	}

	var resp = &pbsvc.GetSecretSessionResponse{}
	if sctSession == nil {
		resp.Item = nil
		return resp, nil
	}

	ssLogs, err := repoFn.LookupSecretSessionLogs(ctx, sctSession.GetPublicId(), secret.WithOrder("create_time DESC"))
	if err != nil {
		return nil, err
	}
	if sctSession == nil {
		resp.Item = nil
		return resp, nil
	}

	o, err := toProtoSessionLog(ssLogs, nil)
	if err != nil {
		return nil, err
	}

	o.Id = req.Id
	resp.Item = o

	return resp, nil
}

func toProto(in secret.Secret, m []*secret.Secret) (*pbs.Secret, error) {
	out := pbs.Secret{
		Id:          in.GetPublicId(),
		ScopeId:     in.GetScopeId(),
		Manager:     in.GetManager(),
		Name:        in.GetName(),
		Code:        in.GetCode(),
		CreatedTime: in.GetCreateTime().GetTimestamp(),
	}
	return &out, nil
}

func toProtoSession(in secret.SecretSession, m []*secret.SecretSession) (*pbs.SecretSession, error) {
	out := pbs.SecretSession{
		Id:          in.GetPublicId(),
		ScopeId:     in.GetScopeId(),
		Description: in.GetDescription(),
		CreatedTime: in.GetCreateTime().GetTimestamp(),
		UpdatedTime: in.GetUpdatedTime().GetTimestamp(),
	}

	return &out, nil
}

func toProtoSessionLog(in secret.SecretSessionLog, m []*secret.SecretSession) (*pbs.SecretSessionLog, error) {
	decodedOutput, err := base64.StdEncoding.DecodeString(in.GetOutput())
	if err != nil {
		return nil, err
	}
	out := pbs.SecretSessionLog{
		Id:     in.GetPublicId(),
		Output: string(decodedOutput),
	}

	return &out, nil
}
