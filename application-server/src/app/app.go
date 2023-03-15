package app

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Executing /%s!", r.URL.Path[1:])
}
