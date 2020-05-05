package iam

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/vault/sdk/helper/base62"
	"github.com/hashicorp/watchtower/internal/db"
	"github.com/hashicorp/watchtower/internal/iam/store"
	"google.golang.org/protobuf/proto"
)

// Roles are granted permissions and assignable to User and Groups
type Role struct {
	*store.Role
	tableName string `gorm:"-"`
}

// ensure that Group implements the interfaces of: Resource, ClonableResource, and db.VetForWriter
var _ Resource = (*Role)(nil)
var _ ClonableResource = (*Role)(nil)
var _ db.VetForWriter = (*Role)(nil)

// NewRole creates a new role with a scope (project/organization)
// options include: withDescripion, withFriendlyName
func NewRole(primaryScope *Scope, opt ...Option) (*Role, error) {
	opts := GetOpts(opt...)
	withFriendlyName := opts.withFriendlyName
	withDescription := opts.withDescription
	if primaryScope == nil {
		return nil, errors.New("error the role primary scope is nil")
	}
	if primaryScope.Type != OrganizationScope.String() &&
		primaryScope.Type != ProjectScope.String() {
		return nil, errors.New("roles can only be within an organization or project scope")
	}
	publicId, err := base62.Random(20)
	if err != nil {
		return nil, fmt.Errorf("error generating public id %w for new role", err)
	}
	r := &Role{
		Role: &store.Role{
			PublicId:       publicId,
			PrimaryScopeId: primaryScope.GetPublicId(),
		},
	}
	if withFriendlyName != "" {
		r.FriendlyName = withFriendlyName
	}
	if withDescription != "" {
		r.Description = withDescription
	}
	return r, nil
}

// Clone creates a clone of the Role
func (r *Role) Clone() Resource {
	cp := proto.Clone(r.Role)
	return &Role{
		Role: cp.(*store.Role),
	}
}

// AssignedRoles returns a list of principal roles (Users and Groups) for the Role.
func (role *Role) AssignedRoles(ctx context.Context, r db.Reader) ([]AssignedRole, error) {
	viewRoles := []*assignedRoleView{}
	if err := r.SearchBy(
		ctx,
		&viewRoles,
		"role_id = ? and type in(?, ?)",
		role.Id, UserRoleType.String(), GroupRoleType.String()); err != nil {
		return nil, fmt.Errorf("error getting assigned roles %w", err)
	}

	pRoles := []AssignedRole{}
	for _, vr := range viewRoles {
		switch vr.Type {
		case UserRoleType.String():
			pr := &UserRole{
				UserRole: &store.UserRole{
					Id:             vr.Id,
					CreateTime:     vr.CreateTime,
					UpdateTime:     vr.UpdateTime,
					PublicId:       vr.PublicId,
					FriendlyName:   vr.FriendlyName,
					PrimaryScopeId: vr.PrimaryScopeId,
					RoleId:         vr.RoleId,
					Type:           UserRoleType.String(),
					PrincipalId:    vr.PrincipalId,
				},
			}
			pRoles = append(pRoles, pr)
		case GroupRoleType.String():
			pr := &GroupRole{
				GroupRole: &store.GroupRole{
					Id:             vr.Id,
					CreateTime:     vr.CreateTime,
					UpdateTime:     vr.UpdateTime,
					PublicId:       vr.PublicId,
					FriendlyName:   vr.FriendlyName,
					PrimaryScopeId: vr.PrimaryScopeId,
					RoleId:         vr.RoleId,
					Type:           GroupRoleType.String(),
					PrincipalId:    vr.PrincipalId,
				},
			}
			pRoles = append(pRoles, pr)
		default:
			return nil, fmt.Errorf("error unsupported role type: %s", vr.Type)
		}
	}
	return pRoles, nil
}

// VetForWrite implements db.VetForWrite() interface
func (role *Role) VetForWrite(ctx context.Context, r db.Reader, opType db.OpType) error {
	if role.PublicId == "" {
		return errors.New("error public id is empty string for role write")
	}
	if role.PrimaryScopeId == "" {
		return errors.New("error primary scope id not set for role write")
	}
	// make sure the scope is valid for users
	if err := role.primaryScopeIsValid(ctx, r); err != nil {
		return err
	}
	return nil
}

func (role *Role) primaryScopeIsValid(ctx context.Context, r db.Reader) error {
	ps, err := LookupPrimaryScope(ctx, r, role)
	if err != nil {
		return err
	}
	if ps.Type != OrganizationScope.String() && ps.Type != ProjectScope.String() {
		return errors.New("error primary scope is not an organization or project for role")
	}
	return nil
}

// GetPrimaryScope returns the PrimaryScope for the Role
func (role *Role) GetPrimaryScope(ctx context.Context, r db.Reader) (*Scope, error) {
	return LookupPrimaryScope(ctx, r, role)
}

// ResourceType returns the type of the Role
func (*Role) ResourceType() ResourceType { return ResourceTypeRole }

// Actions returns the  available actions for Role
func (*Role) Actions() map[string]Action {
	return StdActions()
}

// TableName returns the tablename to override the default gorm table name
func (r *Role) TableName() string {
	if r.tableName != "" {
		return r.tableName
	}
	return "iam_role"
}

// SetTableName sets the tablename and satisfies the ReplayableMessage interface
func (r *Role) SetTableName(n string) {
	if n != "" {
		r.tableName = n
	}
}