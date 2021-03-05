package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/matthewmazzanti/wordgame/srv/game"
	"github.com/matthewmazzanti/wordgame/srv/graph/generated"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
)

func (r *mutationResolver) SetUser(ctx context.Context) (*model.User, error) {
	user := ctx.Value("user").(*model.User)
	w := ctx.Value("writer").(http.ResponseWriter)

	/*
	fmt.Println(user)
	fmt.Println(w)
	*/

	if user == nil {
		user = &model.User{
			ID:   game.RandID(),
			Name: "Matthew",
		}

		r.Users[user.ID] = user

		http.SetCookie(w, &http.Cookie{
			Name:     "user-id",
			Value:    user.ID,
			SameSite: http.SameSiteNoneMode,
		})

		fmt.Printf("No cookie set! Set to: %s\n", user.ID)
	}

	return user, nil
}

func (r *mutationResolver) NewGame(ctx context.Context) (*model.Game, error) {
	user := ctx.Value("user").(*model.User)
	if user == nil {
		return nil, fmt.Errorf("No user found")
	}

	game, err := game.New(r.DB)
	if err != nil {
		return nil, err
	}

	r.Games[game.ID] = game

	return game.Freeze(), nil
}

func (r *mutationResolver) AddGuess(ctx context.Context, id string, guess string) (*model.GuessResult, error) {
	user := ctx.Value("user").(*model.User)
	if user == nil {
		return nil, fmt.Errorf("No user found")
	}

	game, ok := r.Games[id]
	if !ok {
		return nil, fmt.Errorf("No game with id %s", id)
	}

	return game.Guess(guess), nil
}

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	user := ctx.Value("user").(*model.User)
	if user == nil {
		return nil, fmt.Errorf("No user found")
	}

	game, ok := r.Games[id]
	if !ok {
		return nil, fmt.Errorf("No game with id %s", id)
	}

	return game.Freeze(), nil
}

func (r *subscriptionResolver) WatchGame(ctx context.Context, id string) (<-chan *model.GuessResult, error) {
	user := ctx.Value("user").(*model.User)
	if user == nil {
		return nil, fmt.Errorf("No user found")
	}

	game, ok := r.Games[id]
	if !ok {
		return nil, fmt.Errorf("No game with id %s", id)
	}

	watchID := int(time.Now().UnixNano())
	channel := game.Watch(watchID)
	go func() {
		<-ctx.Done()
		game.Unwatch(watchID)
	}()

	return channel, nil
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
