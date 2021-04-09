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
	Date     string  `json:"date"`
	Price    float32 `json:"price"`
	Open     float32 `json:"open"`
	High     float32 `json:"high"`
	Low      float32 `json:"low"`
	Volume   string  `json:"volume"`
	Change   string  `json:"change"`
	Currency string  `json:"currency"`
}

var stocks = make(map[string][]Stock, 0)

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

		data := Stock{
			Date:     entry[2],
			Price:    float32(price),
			Open:     float32(open),
			High:     float32(high),
			Low:      float32(low),
			Volume:   entry[7],
			Change:   entry[8],
			Currency: "PHP",
		}

		if stock, ok := stocks[entry[1]]; ok {
			stocks[entry[1]] = append(stock, data)
			continue
		}

		stocks[entry[1]] = []Stock{data}
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
	code := strings.ToUpper(params["code"])
	count, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 8)
	if err != nil || count <= 0 {
		count = int64(len(stocks[code]))
	}

	// Return the stock with the matching code
	for k, v := range stocks {
		if k == code {
			_ = json.NewEncoder(w).Encode(
				struct {
					Code string  `json:"code"`
					Data []Stock `json:"data"`
				}{k, v[:count]},
			)
			break
		}
	}
}
