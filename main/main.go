package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type Stock struct {
	Code  string
	Name  string
	Price float32
}

var stocks []Stock

func main() {
	handle()
}

func init() {
	stocks = []Stock{
		{"A", "AlphabetA", 1.25},
		{"B", "AlphabetB", 1.5},
		{"C", "AlphabetC", 2.0},
	}
}

func handle() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/stocks", getAllStocks)
	router.HandleFunc("/stocks/{code}", getStock)
	log.Println("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")
}

func getAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getAllStocks")

	// Encode stocks and write as part of response
	_ = json.NewEncoder(w).Encode(stocks)
}

func getStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getStock")

	// Get the params from the request
	params := mux.Vars(r)
	code := params["code"]

	// Return the stock with the matching code
	for _, stock := range stocks {
		if stock.Code == strings.ToUpper(code) {
			_ = json.NewEncoder(w).Encode(stock)
		}
	}
}
