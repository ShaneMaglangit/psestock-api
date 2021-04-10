package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getStock")

	// Get the params from the request
	params := mux.Vars(r)
	code := strings.ToUpper(params["code"])

	// Time params
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	// Find the stock with the matching code
	stock := findStockByCode(stocks, code)

	// Filter by time
	stock = filterStockByDate(stock, start, end)

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
