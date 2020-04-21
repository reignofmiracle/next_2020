package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Haaaaello World"))
	})

	http.ListenAndServe("localhost:8088", nil)
}
