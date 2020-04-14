package domain

import (
	"fmt"
	"net/http"

	"github.com/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {
			ID:        1,
			FirstName: "Eduardo",
			LastName:  "Fernandez",
			Email:     "fenandez9000@gmail.com",
		},
	}
)

// GetUser from data layer
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	user := users[userID]

	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User with id = %v not found", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}

	return user, nil
}
