package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/akhilmhdh/eventzo/pkg/graphql/graph"
	"github.com/akhilmhdh/eventzo/pkg/graphql/models"
)

func (r *queryResolver) GetAnEvent(ctx context.Context, id string) (*models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
