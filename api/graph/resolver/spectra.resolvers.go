package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated "github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
)

func (r *mutationResolver) Privileges(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *mutationResolver) Organization(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) Privileges(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) Organization(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
