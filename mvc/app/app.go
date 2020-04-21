package app

import (
	"net/http"

	"github.com/golang-microservices/mvc/controllers"
)

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if error := http.ListenAndServe("localhost:8080", nil); error != nil {
		panic(error)
	}
}
