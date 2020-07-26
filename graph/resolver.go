package graph

import (
	"context"

	"github.com/unrealnerd/gql-postgres/graph/model"
	"github.com/unrealnerd/gql-postgres/graph/generated"
	"github.com/unrealnerd/gql-postgres/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ... this struct is what link to the generated code
type Resolver struct {
	Products []*model.Product
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	repo := &repo.InventoryRepo{}
	quotes := repo.Find("1 = $1", 1)// dummy query

	return quotes, nil
}