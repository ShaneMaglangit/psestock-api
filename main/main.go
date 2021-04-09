package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Stock struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Date     string  `json:"date"`
	Price    float32 `json:"price"`
	Open     float32 `json:"open"`
	High     float32 `json:"high"`
	Low      float32 `json:"low"`
	Volume   string  `json:"volume"`
	Change   string  `json:"change"`
	Currency string  `json:"currency"`
}

var stocks = make([]Stock, 0)

func main() {
	handle()
}

func init() {
	file, err := os.Open("main/stocks.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, entry := range data {
		price, _ := strconv.ParseFloat(entry[3], 32)
		open, _ := strconv.ParseFloat(entry[4], 32)
		high, _ := strconv.ParseFloat(entry[5], 32)
		low, _ := strconv.ParseFloat(entry[6], 32)

		stock := Stock{
			Name:     entry[0],
			Code:     entry[1],
			Date:     entry[2],
			Price:    float32(price),
			Open:     float32(open),
			High:     float32(high),
			Low:      float32(low),
			Volume:   entry[7],
			Change:   entry[8],
			Currency: "PHP",
		}

		stocks = append(stocks, stock)
	}
}

func handle() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/stocks/{code}", getStock)
	log.Println("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
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
