package app

import (
	"fmt"
	"net/http"
)

func StartServer() error {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

	return nil
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
