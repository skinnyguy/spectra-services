package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
	"github.com/skinnyguy/spectra-services/api/services"
)

func (r *roleMutationResolver) AddRole(ctx context.Context, obj *model.AbstractModel, request model.InputRole) (*model.RoleResponseWithMessage, error) {
	role, err := services.AddRole(ctx, request)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleMutationResolver) UpdateRole(ctx context.Context, obj *model.AbstractModel, id string, name string, description string) (*model.RoleResponseWithMessage, error) {
	role, err := services.UpdateRole(ctx, id, name, description)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleMutationResolver) DeleteRole(ctx context.Context, obj *model.AbstractModel, id string) (*model.RoleResponseWithMessage, error) {
	role, err := services.DeleteRole(ctx, id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleQueryResolver) GetListRole(ctx context.Context, obj *model.AbstractModel, request model.GetPaginationRoleRequest) (*model.RoleMultiResponse, error) {
	roles, err := services.GetListRole(ctx, request)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// RoleMutation returns graph.RoleMutationResolver implementation.
func (r *Resolver) RoleMutation() graph.RoleMutationResolver { return &roleMutationResolver{r} }

// RoleQuery returns graph.RoleQueryResolver implementation.
func (r *Resolver) RoleQuery() graph.RoleQueryResolver { return &roleQueryResolver{r} }

type roleMutationResolver struct{ *Resolver }
type roleQueryResolver struct{ *Resolver }
