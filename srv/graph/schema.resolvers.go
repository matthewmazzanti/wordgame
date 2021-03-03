package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/matthewmazzanti/wordgame/srv/graph/generated"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
)

func (r *mutationResolver) CreateGame(ctx context.Context) (*model.Game, error) {
	return r.Game.Freeze(), nil
}

func (r *mutationResolver) AddGuess(ctx context.Context, guess string) (*model.GuessResult, error) {
	return r.Game.Guess(guess), nil
}

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return r.Resolver.Game.Freeze(), nil
}

func (r *subscriptionResolver) WatchGame(ctx context.Context) (<-chan *model.GuessResult, error) {
	id := int(time.Now().UnixNano())
	c := r.Game.Watch(id)
	go func() {
		<-ctx.Done()
		r.Game.Unwatch(id)
	}()
	return c, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
