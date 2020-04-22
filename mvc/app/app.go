package app

import (
	"net/http"

	"github.com/golang-microservices/mvc/controllers"
)

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if error := http.ListenAndServe("localhost:9191", nil); error != nil {
		panic(error)
	}
}
