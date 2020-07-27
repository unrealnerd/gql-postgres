package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/unrealnerd/gql-postgres/graph/generated"
	"github.com/unrealnerd/gql-postgres/graph/model"
)

func (r *queryResolver) Products(ctx context.Context, first *int, after *int) ([]*model.Product, error) {
	return r.getProductsFromInventory(ctx, first, after)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
