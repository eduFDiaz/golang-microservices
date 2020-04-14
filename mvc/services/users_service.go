package services

import (
	"github.com/golang-microservices/mvc/domain"
	"github.com/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
