package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashcott/go-project/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Printf("Starting server at port 3000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
