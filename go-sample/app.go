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

	http.HandleFunc("/franklin-says", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", quotes[rand.Intn(len(quotes))])
	})

	log.Fatal(http.ListenAndServe(":8001", nil))
}
