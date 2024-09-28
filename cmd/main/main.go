package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinayak3010/Book-Line/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
