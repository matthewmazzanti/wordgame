package game

import (
	"log"
	"database/sql"
	"math/rand"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
)

type Word struct {
	text string
	guessed bool
}

type guessMsg struct {
	text string
	out chan<- *model.Game
}

type addWatchMsg struct {
	id int
	out chan<- *model.Game
}

type delWatchMsg struct {
	id int
}

type Game struct {
	letters string
	words []Word
	complete bool
	in chan interface{}
	outs map[int]chan<- *model.Game
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
		in: make(chan interface{}),
		outs: make(map[int]chan<- *model.Game),
	}

	go g.update()

	return g, nil
}

func (g *Game) Freeze() *model.Game {
	guessed := make([]string, 0)
	for _, word := range g.words {
		if word.guessed {
			guessed = append(guessed, word.text)
		}
	}

	fmt.Println(guessed)

	return &model.Game{
		ID: g.letters,
		Letters: g.letters,
		Guessed: guessed,
	}
}


func (g *Game) Guess(guess string) *model.Game {
	if g.complete {
		return g.Freeze()
	}

	// Channel to recieve result on
	out := make(chan *model.Game)

	g.in<- guessMsg{
		text: guess,
		out: out,
	}

	fmt.Println("waiting for response")
	res := <-out
	fmt.Println("got response")
	close(out)
	return res
}

func (g *Game) update() {
	for !g.complete {
		fmt.Println("waiting for event")
		input := <-g.in
		switch event := input.(type) {
		case guessMsg:
			g.addGuess(event.text)
			state := g.Freeze()
			event.out<- state
			for _, out := range g.outs {
				fmt.Println("sending")
				out<- state
			}
		case addWatchMsg:
			g.outs[event.id] = event.out
			fmt.Println("added watch")
			fmt.Println(event.id)
			fmt.Println(len(g.outs))
		case delWatchMsg:
			fmt.Println("deleted watch")
			fmt.Println(event.id)
			fmt.Println(len(g.outs))
			out, ok := g.outs[event.id]
			if ok {
				close(out)
				delete(g.outs, event.id)
			}
		default:
			fmt.Println("Unknown event type")
		}
	}

	fmt.Println("game completed")
	close(g.in)
	for _, out := range g.outs {
		close(out)
	}
}

func (g *Game) addGuess(guess string) bool {
	fmt.Println("Checking guess")

	for i, word := range g.words {
		if guess == word.text {
			g.words[i].guessed = true
			g.complete = g.isComplete()
			fmt.Println("Guess correct")
			return true
		}
	}

	fmt.Println("Guess incorrect")
	return false
}

func (g *Game) isComplete() bool {
	for _, word := range g.words {
		if !word.guessed {
			return false
		}
	}

	return true
}

func (g *Game) Watch(id int) <-chan *model.Game {
	out := make(chan *model.Game)

	g.in<- addWatchMsg{
		id: id,
		out: out,
	}

	return out
}

func (g *Game) Unwatch(id int) {
	g.in<- delWatchMsg{ id: id }
}

func query(db *sql.DB, letters string) ([]Word, error) {
	mask, primary := makeMasks(letters)

	rows, err := db.Query(
		"select word from word where bitmap & ? = 0 and bitmap & ? > 0;",
		mask,
		primary,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	words := make([]Word, 0)
	for rows.Next() {
		var text string;

		err := rows.Scan(&text)
		if err != nil {
			log.Fatal(err)
		}

		words = append(words, Word{
			text: text,
			guessed: false,
		})
	}

	return words, nil
}

func makeMasks(letters string) (uint32, uint32) {
	var mask uint32 = 0
	for i := 1; i < len(letters); i++ {
		letter := letters[i]
		mask = mask | 1 << (int(letter) - 97)
	}

	var primary uint32 = 1 << (int(letters[0]) - 97)
	mask = ^(mask | primary)

	return mask, primary
}

func randFrom(letters string) byte {
	return letters[rand.Intn(len(letters))]
}

func randVowel() byte {
	return randFrom("aeiou")
}

func randAlpha() byte {
	return randFrom("abcdefghijklmnoprtuvwy")
}

func randLetters(length int) string {
	if length < 1 {
		return ""
	}

	letters := make([]byte, 1)
	letters[0] = randVowel()

	for len(letters) < length {
		letter := randAlpha()
		contains := false
		for _, l := range letters {
			if l == letter {
				contains = true
			}
		}

		if !contains {
			letters = append(letters, letter)
		}
	}

	return string(shuffle(letters))
}

func shuffle(letters []byte) []byte {
	for i := len(letters) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		letters[i], letters[j] = letters[j], letters[i]
	}

	return letters
}
