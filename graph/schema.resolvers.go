package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MurrayCode/graphQLGo/database"
	"github.com/MurrayCode/graphQLGo/graph/generated"
	"github.com/MurrayCode/graphQLGo/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateWatch(ctx context.Context, input *model.NewWatch) (*model.Watch, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Watch(ctx context.Context, id string) (*model.Watch, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Watches(ctx context.Context) ([]*model.Watch, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
