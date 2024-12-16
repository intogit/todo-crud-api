package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ninja/Todo/db"
	"github.com/ninja/Todo/handlers"
	"github.com/ninja/Todo/services"
)

type Application struct {
	// TODO 1:
	// There should not be need of any seperate file serviceModels etc.. to import Todo Struct
	// might be a kind of bring modularity to code.. microservices type..
	// TodoModel services.Todo

	Models services.TodoModel
}

func main() {
	fmt.Println("Todo main file is executing")
	fmt.Println("DB use - connection initiated")
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}
	todoContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer func() {
		if err := mongoClient.Disconnect(todoContext); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)
	log.Println("services running on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
	// TODO 2:
	// CHECK IF this can work without configuring the cors using chi in handlers/routers.go
	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe(":8080", handler)
}
