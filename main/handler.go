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
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getStock")

	// Get the params from the request
	params := mux.Vars(r)
	code := strings.ToUpper(params["code"])

	// Time params
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	// Count params
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 8)

	// Find the stock with the matching code
	stocks := queryStockFromDb(code, start, end, limit)

	// Write data to response
	_ = json.NewEncoder(w).Encode(stocks)
}

func queryStockFromDb(code, start, end string, limit int64) []Stock {
	stocks := make([]Stock, 0)
	startDate, endDate := parseStartEndDate(start, end)

	db := mongoClient.Database("psestocks")
	collection := db.Collection("prices")

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	cursor, err := collection.Find(ctx, bson.M{"code": code, "date": bson.M{"$gte": startDate, "$lte": endDate}}, findOptions)
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
		stocks = append(stocks, stock)
	}

	return stocks
}
