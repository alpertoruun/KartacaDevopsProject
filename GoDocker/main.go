package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"math/rand"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Go!")
}

func randomCountry(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo1:27017,stajdb")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	countriesCollection := client.Database("stajdb").Collection("ulkeler")
	cursor, err := countriesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var countries []bson.M
	if err = cursor.All(context.Background(), &countries); err != nil {
		log.Fatal(err)
	}

	randomIndex := rand.Intn(len(countries))
	randomCountry := countries[randomIndex]

	fmt.Fprintf(w, "Random country: %s", randomCountry)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/staj", randomCountry)

	log.Fatal(http.ListenAndServe(":5555", nil))
}