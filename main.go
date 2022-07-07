package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHAndler(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Hola mundo2"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHAndler)
	http.ListenAndServe(":3000", r)
}
