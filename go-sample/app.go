package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("../quotes.txt")
	if err != nil {
		panic(err)
	}

	quotes := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	numQuotes := len(quotes)

	http.HandleFunc("/franklin-says", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n", quotes[rand.Intn(numQuotes)])
	})

	log.Fatal(http.ListenAndServe(":8001", nil))
}
