package graph

import (
	"context"

	"github.com/Daviiap/maxsum/graph/generated"
	"github.com/Daviiap/maxsum/graph/model"
	"github.com/Daviiap/maxsum/graph/resolvers"
)

func (r *mutationResolver) MaxSum(ctx context.Context, list []int) (*model.MaxSumResponse, error) {
	return resolvers.MaxSumResolver(list), nil
}

func (r *queryResolver) Test(ctx context.Context) (string, error) {
	res := "running"
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
