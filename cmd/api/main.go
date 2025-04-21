package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/internal/handlers"
	_ "go-rest-api/internal/models"
	_ "go-rest-api/internal/repository"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	handlers.SetupTodoRoutes(router)

	fmt.Println("Server está rodando na porta: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
