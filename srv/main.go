package main

import (
	"os"
	"io"
	"log"
	"fmt"
	"errors"
	"net/http"
	"bufio"
	"strings"
	"encoding/hex"
	"encoding/binary"
	"io/ioutil"
)

func main() {
	db, err := ReadDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", HomeHandler(db))
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func HomeHandler(db []Entry) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		EnableCors(&w)

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		body := string(bodyBytes)
		if len(body) == 0 {
			return
		}
		fmt.Println(body)

		words := Lookup(db, body)

		for i := 0; i < len(words); i++ {
			io.WriteString(w, words[i])
			io.WriteString(w, "\n")
		}
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Entry struct {
	Word string
	Bitmap uint32
}

func ReadBitmap(bitmap string) (uint32, error) {
	bytes, err := hex.DecodeString(bitmap)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(bytes), nil
}

func ReadDB() ([]Entry, error) {
	root := os.Getenv("ROOT")
	if root == "" {
		return nil, errors.New("Could not find ROOT env variable")
	}

	file, err := os.Open(root + "/words/data.csv")
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	entries := []Entry{}
	for true {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = line[0:len(line)-1]
		split := strings.Split(line, ",")

		word := split[0]
		if len(word) <= 3 {
			continue
		}

		bitmap, err := ReadBitmap(split[1])
		if err != nil {
			return nil, err
		}

		entries = append(entries, Entry {
			Word: split[0],
			Bitmap: bitmap,
		})
	}

	return entries, nil
}

func Lookup(entries []Entry, letters string) []string {
	var mask uint32 = 0
	for i := 1; i < len(letters); i++ {
		letter := letters[i]
		mask = mask | 1 << (int(letter) - 97)
	}

	var primary uint32 = 1 << (int(letters[0]) - 97)
	mask = ^(mask | primary)

	res := []string{}
	for i := 1; i < len(entries); i++ {
		entry := entries[i]
		if entry.Bitmap & primary > 0 && entry.Bitmap & mask == 0 {
			res = append(res, entry.Word)
		}
	}

	return res
}
