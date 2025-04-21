package main

import (
	"fmt"
	"go-rest-api/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handlers.SetupTodoRoutes(router)

	fmt.Println("Server está rodando na porta: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
