package game

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
	"sync"
)

type Word struct {
	text string
	guessed bool
}

type Game struct {
	mutex *sync.Mutex
	clients map[int]chan<- *model.GuessResult

	letters string
	complete bool
	words []Word
	incorrect []string
}

func New(db *sql.DB) (*Game, error) {
	letters := randLetters(7)
	words, err := query(db, letters)
	if err != nil {
		return nil, err
	}

	fmt.Println(letters)
	fmt.Println(words)

	g := &Game{
		letters: letters,
		words: words,
		mutex: &sync.Mutex{},
		clients: make(map[int]chan<- *model.GuessResult),
	}

	return g, nil
}

func (g *Game) Freeze() *model.Game {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return g.freeze()
}

func (g *Game) freeze() *model.Game {
	correct := make([]string, 0)
	for _, word := range g.words {
		if word.guessed {
			correct = append(correct, word.text)
		}
	}

	incorrect := make([]string, len(g.incorrect))
	copy(incorrect, g.incorrect)


	return &model.Game{
		ID: g.letters,
		Letters: g.letters,
		Correct: correct,
		Incorrect: incorrect,
		Total: len(g.words),
	}
}


func (g *Game) Guess(guess string) *model.GuessResult {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	fmt.Println("Checking guess")

	correct := false
	if !g.complete {
		for i, word := range g.words {
			if guess == word.text {
				g.words[i].guessed = true
				g.complete = g.isComplete()
				correct = true
				break
			}
		}
	}

	if (!correct) {
		g.incorrect = append(g.incorrect, guess)
	}

	res := &model.GuessResult{
		Correct: correct,
		Word: guess,
		Game: g.freeze(),
	}

	for _, client := range g.clients {
		client<- res
	}

	return res
}

func (g *Game) Watch(id int) <-chan *model.GuessResult {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	client := make(chan *model.GuessResult)
	g.clients[id] = client

	fmt.Println("added client")
	fmt.Println(id)
	fmt.Println(len(g.clients))

	return client
}

func (g *Game) Unwatch(id int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	client, ok := g.clients[id]
	if ok {
		close(client)
		delete(g.clients, id)
	}

	fmt.Println("deleted client")
	fmt.Println(id)
	fmt.Println(len(g.clients))
}

func (g *Game) isComplete() bool {
	for _, word := range g.words {
		if !word.guessed {
			return false
		}
	}

	return true
}
