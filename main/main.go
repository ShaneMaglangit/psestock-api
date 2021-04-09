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
	"time"
)

type Stock struct {
	Date   string  `json:"date"`
	Price  float32 `json:"price"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Volume string  `json:"volume"`
	Change string  `json:"change"`
}

var stocks = make(map[string][]Stock, 0)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/stocks/{code}", getStock)
	log.Println("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
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
			Date:   entry[2],
			Price:  float32(price),
			Open:   float32(open),
			High:   float32(high),
			Low:    float32(low),
			Volume: entry[7],
			Change: entry[8],
		}

		if stock, ok := stocks[entry[1]]; ok {
			stocks[entry[1]] = append(stock, data)
			continue
		}

		stocks[entry[1]] = []Stock{data}
	}
}

func getStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getStock")

	// Get the params from the request
	params := mux.Vars(r)
	code := strings.ToUpper(params["code"])

	// Time params
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	// Find the stock with the matching code
	stock := getStockByCode(code)

	// Filter by time
	stock = filterStockByTime(stock, start, end)

	// Count params
	count, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 8)
	if err != nil || count <= 0 {
		count = int64(len(stock))
	}

	// Write data to response
	_ = json.NewEncoder(w).Encode(
		struct {
			Code     string  `json:"code"`
			Currency string  `json:"currency"`
			Data     []Stock `json:"data"`
		}{code, "PHP", stock[:count]},
	)
}

func getStockByCode(code string) []Stock {
	for k, v := range stocks {
		if k == code {
			return v
		}
	}
	return nil
}

func filterStockByTime(stock []Stock, start, end string) []Stock {
	ret := make([]Stock, 0)

	// Set start and end date to default
	startDate, _ := time.Parse("2006-01-02", "1927-08-08")
	endDate := time.Now()

	// Parse the dates from the parameters
	if start != "" {
		startDate, _ = time.Parse("2006-01-02", start)
	}
	if end != "" {
		endDate, _ = time.Parse("2006-01-02", end)
	}

	// Filter the stocks by finding those within the date range given
	for _, data := range stock {
		stockDate, _ := time.Parse("Jan 02, 2006", data.Date)
		if stockDate.After(startDate) && stockDate.Before(endDate) {
			ret = append(ret, data)
		}
	}

	return ret
}
