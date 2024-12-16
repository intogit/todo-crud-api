package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ninja/Todo/services"
)

type Response struct {
	Msg  string
	Code int
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo services.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
	}
	err = services.InsertTodo(newTodo)
	if err != nil {
		errorRes := Response{"Error adding new Todo", 304}
		w.Header()
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	res := Response{"Successfully added new Todo ", 200}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

func HandleGetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := services.GetAllTodo()
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todos)
}

func HandleGetTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println(id)
	todo, err := services.GetTodoById(id)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todo)
}

func HandleHeathCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{"version 2 api - health check success", 200}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	// TODO 3:
	// check if this json.marshal is needed
	// res := Response {"health check success", 200 }
	// jsonRes, err := json.Marshal(res)
	// 	// if err != nil {
	// 	// 	log.Println(err)
	// 	// 	return
	// 	// }
	// 	// w.Header().Set("Content-Type", "application/json")
	// 	// w.Write(jsonRes)
	// 	// w.Write([]byte("welcome to golang ninja todo web-api;; \n, ------heath check success"))
}
