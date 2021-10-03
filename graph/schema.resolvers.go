package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MurrayCode/graphQLGo/graph/generated"
	"github.com/MurrayCode/graphQLGo/graph/model"
	"github.com/MurrayCode/graphQLGo/watches"
)

func (r *mutationResolver) CreateWatch(ctx context.Context, input *model.NewWatch) (*model.Watch, error) {
	return db.Save(input), nil
}

func (r *mutationResolver) UpdateWatch(ctx context.Context, id string, input *model.UpdateWatch) (*model.Watch, error) {
	return db.Update(id, input), nil
}

func (r *mutationResolver) DeleteWatch(ctx context.Context, id string) (*model.Watch, error) {
	return db.Delete(id), nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = watches.Connect()
