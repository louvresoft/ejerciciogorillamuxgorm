package main

import (
	"github.com/gorilla/mux"
	"github.com/louvresoft/go-gorm-gorilla-mux/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHAndler)
	http.ListenAndServe(":3000", r)
}
