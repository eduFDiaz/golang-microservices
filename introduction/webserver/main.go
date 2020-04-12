package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello yaniel!"))
	})

	if error := http.ListenAndServe("localhost:8080", nil); error != nil {
		panic(error)
	}
}
