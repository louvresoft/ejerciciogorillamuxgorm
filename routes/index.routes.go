package routes

import "net/http"

func HomeHAndler(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Hola mundo2"))
}
