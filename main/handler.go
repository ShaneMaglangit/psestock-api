package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Stock struct {
	Code   string    `json:"code" bson:"code"`
	Date   time.Time `json:"date" bson:"date"`
	Price  float32   `json:"price" bson:"price"`
	Open   float32   `json:"open" bson:"open"`
	High   float32   `json:"high" bson:"high"`
	Low    float32   `json:"low" bson:"low"`
	Volume string    `json:"volume" bson:"volume"`
	Change string    `json:"change" bson:"change"`
}

type RequestParams struct {
	code      string
	startDate time.Time
	endDate   time.Time
	asc       bool
	limit     int64
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")
}

func stockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getStock")

	// Get the params from the request
	vars := mux.Vars(r)

	// Extract the params
	params := RequestParams{
		code:      strings.ToUpper(vars["code"]),
		startDate: parseDate(time.Date(1927, time.August, 8, 0, 0, 0, 0, time.UTC), r.URL.Query().Get("start")),
		endDate:   parseDate(time.Now(), r.URL.Query().Get("end")),
		asc:       strings.ToUpper(r.URL.Query().Get("asc")) == "TRUE",
		limit:     getLimit(r.URL.Query().Get("limit")),
	}
	_ = params

	// Find the stock with the matching code
	stocks := queryStockFromDb(params)

	// Write data to
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	_ = json.NewEncoder(w).Encode(stocks)
}

func queryStockFromDb(params RequestParams) []Stock {
	stocks := make([]Stock, 0)

	db := mongoClient.Database("psestocks")
	collection := db.Collection("prices")

	findOptions := options.Find()
	findOptions.SetLimit(params.limit)

	cursor, err := collection.Find(ctx, bson.M{"code": params.code, "date": bson.M{"$gte": params.startDate, "$lte": params.endDate}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var stock Stock
		if err = cursor.Decode(&stock); err != nil {
			log.Fatal(err)
			return nil
		}

		if params.asc {
			stocks = append([]Stock{stock}, stocks...)
			continue
		}
		stocks = append(stocks, stock)
	}

	return stocks
}

func parseDate(def time.Time, timeString string) time.Time {
	if timeString == "" {
		return def
	}
	date, _ := time.Parse("2006-01-02", timeString)
	return date
}

func getLimit(s string) int64 {
	switch s {
	case "":
		return 1
	case "all":
		return 0
	default:
		limit, _ := strconv.ParseInt(s, 10, 8)
		return limit
	}
}
