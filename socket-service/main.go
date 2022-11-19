package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielyea/go-sockets/socket-service/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	h := handlers.NewHandler()
	router.HandleFunc("/socket", h.WScable)
	router.HandleFunc("/", h.Home)
	fmt.Printf("Running socket on port :8000 \n")
	http.ListenAndServe(":8000", router)
}
