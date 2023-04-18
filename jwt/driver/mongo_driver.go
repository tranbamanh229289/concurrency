package mongo_driver

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Db *mongo.Database
}

var Mongo = &MongoDB{}

func Connect(user, password, dbname string) (*MongoDB){
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.9t2y8zq.mongodb.net/?retryWrites=true&w=majority", user, password)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Mongo.Db = client.Database(dbname)
	fmt.Println("connection success")
	return Mongo
}