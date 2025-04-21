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
	fmt.Println("\n\n·················································\n:    _    ____ ___           ____  _   _  ____  :\n:   / \\  |  _ |_ _|         |  _ \\| | | |/ ___| :\n:  / _ \\ | |_) | |   _____  | |_) | | | | |     :\n: / ___ \\|  __/| |  |_____| |  __/| |_| | |___  :\n:/_/  _\\_|_|__|___|     ____|_|____\\___/_\\____| :\n:    |  _ \\| ____\\ \\   / / _ \\|  _ \\/ ___|      :\n:    | | | |  _|  \\ \\ / | | | | |_) \\___ \\      :\n:    | |_| | |___  \\ V /| |_| |  __/ ___) |     :\n:    |____/|_____|  \\_/  \\___/|_|   |____/      :\n·················································\n\n")

	fmt.Println("Server está rodando na porta: 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
