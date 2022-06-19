package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/generated"
	"github.com/cbnits/gqlgen-gqlgen-dummy/graph/model"
)

func (r *mutationResolver) UpsertEmployee(ctx context.Context, input model.EmployeeRequest) (*model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EmployeeDbGet(ctx context.Context, input model.EmployeeGetterRequest) (*model.EmployeeGetterResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
