package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-microservices/mvc/services"
	"github.com/golang-microservices/mvc/utils"
)

// GetUser by id from the services
func GetUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	userID, error := strconv.ParseInt(req.URL.Query().Get("userid"), 10, 64)
	if error != nil {
		log.Printf("error = %v", error)

		apiErr := &utils.ApplicationError{
			Message:    "user id needs to be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, err := services.UserService.GetUser(userID)
	if err != nil {
		log.Printf("error = %v", error)

		jsonValue, _ := json.Marshal(err)

		resp.WriteHeader(err.StatusCode)
		resp.Write(jsonValue)
		return
	}

	log.Printf("user = %v", user)

	jsonValue, _ := json.Marshal(user)

	resp.Write(jsonValue)
}
