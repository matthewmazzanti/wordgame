package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matthewmazzanti/wordgame/srv/graph/generated"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
)

func (r *mutationResolver) CreateGame(ctx context.Context) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddGuess(ctx context.Context, guess string) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
