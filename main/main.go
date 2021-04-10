package main

import (
	"encoding/csv"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var stocks = make(map[string][]Stock, 0)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/stocks/{code}", getHandler)
	log.Println("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func init() {
	path := os.Getenv("R_PATH")
	csvPath := filepath.Join(path, "main/stocks.csv")
	lines, err := readCsvFile(csvPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, entry := range lines {
		data := parseLineToStock(entry)
		if stock, ok := stocks[entry[1]]; ok {
			stocks[entry[1]] = append(stock, data)
			continue
		}
		stocks[entry[1]] = []Stock{data}
	}
}

func readCsvFile(csvPath string) ([][]string, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
