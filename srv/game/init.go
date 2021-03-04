package game

import (
	"log"
	"database/sql"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

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

func randCapsAlphaNum() byte {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	ALPHA := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "1234567890"

	return randFrom(alpha + ALPHA + num)
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

func randID(length int) string {
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = randCapsAlphaNum()
	}

	return string(id)
}

func RandID() string {
	return randID(10)
}
