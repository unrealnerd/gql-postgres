package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/unrealnerd/gql-postgres/graph/model"
	"github.com/unrealnerd/gql-postgres/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ... this struct is what link to the generated code
type Resolver struct {
	Products []*model.Product
}

func (r *queryResolver) getProductsFromInventory(ctx context.Context, first *int, after *int) ([]*model.Product, error) {
	repo := &repo.InventoryRepo{}

	if after == nil {
		after = new(int)
	}

	if first == nil {
		first = new(int)
	}

	return repo.GetProducts(*first, *after)
}
