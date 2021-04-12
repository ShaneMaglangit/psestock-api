package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
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
	router.Handle("/favicon.ico", http.NotFoundHandler())
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/{code}", stockHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	var err error

	atlasUri := os.Getenv("ATLAS_URI")
	if atlasUri == "" {
		log.Fatal("Cannot find ATLAS_URI environmental variable")
		return
	}

	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(atlasUri))
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

	//err = mongoClient.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	log.Fatal(err)
	//}
}
