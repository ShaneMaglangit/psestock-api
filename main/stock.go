package main

import (
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

func filterStockByDate(stocks []Stock, start, end string) []Stock {
	ret := make([]Stock, 0)
	startDate, endDate := parseStartEndDate(start, end)

	// Filter the stocks by finding those within the date range given
	for _, data := range stocks {
		if data.Date.After(startDate) && data.Date.Before(endDate) {
			ret = append(ret, data)
		}
	}

	return ret
}

func parseStartEndDate(start string, end string) (time.Time, time.Time) {
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

	return startDate, endDate
}
