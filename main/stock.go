package main

import (
	"strconv"
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

func parseLineToStock(line []string) Stock {
	price, _ := strconv.ParseFloat(line[3], 32)
	open, _ := strconv.ParseFloat(line[4], 32)
	high, _ := strconv.ParseFloat(line[5], 32)
	low, _ := strconv.ParseFloat(line[6], 32)
	data := Stock{line[2], float32(price), float32(open), float32(high), float32(low), line[7], line[8]}
	return data
}

func findStockByCode(stocks map[string][]Stock, code string) []Stock {
	for k, v := range stocks {
		if k == code {
			return v
		}
	}
	return nil
}

func filterStockByDate(stock []Stock, start, end string) []Stock {
	ret := make([]Stock, 0)
	startDate, endDate := parseStartEndDate(start, end)

	// Filter the stocks by finding those within the date range given
	for _, data := range stock {
		stockDate, _ := time.Parse("Jan 02, 2006", data.Date)
		if stockDate.After(startDate) && stockDate.Before(endDate) {
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
