package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

var mongoClient *mongo.Client
var ctx context.Context

func main() {
	defer func() {
		log.Println("Server closing...")
		mongoClient.Disconnect(ctx)
		ctx.Done()
	}()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{code}", getHandler)
	log.Println("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func init() {
	var err error

	atlasUri := flag.String("atlasUri", "", "driver URI from the Atlas Dashboard")
	flag.Parse()

	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(*atlasUri))
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx, _ = context.WithCancel(context.Background())
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
