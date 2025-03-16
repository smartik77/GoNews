package main

import (
	"cmd/pkg/api"
	"cmd/pkg/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var serv server

	// Подключение к БД постгре
	postgreStr := "postgres://testdbuser:testpass@127.0.0.1:5432/news"
	db1, err := pgxpool.New(context.Background(), postgreStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer db1.Close()
	err = db1.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL")

	// Подключение к монгодб
	mongoStr := options.Client().ApplyURI("mongodb://0.0.0.0:27017/")
	db2, err := mongo.Connect(context.Background(), mongoStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db2.Disconnect(context.Background())

	err = db2.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	serv.db = db1 // капец
	serv.api = api.New(serv.db)

	http.ListenAndServe(":8080", serv.api.Router())

	posts, err := serv.db.GetAllPosts()
	if err != nil {
		fmt.Println("Error getting posts: ", err)
	}
	fmt.Println(posts)

}
