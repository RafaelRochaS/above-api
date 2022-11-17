package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RafaelRochaS/above-api/models"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

func HandleAccounts(rg *gin.RouterGroup) {
	accounts := rg.Group("/accounts")

	accounts.POST("/", handleAccountCreation)
}

func handleAccountCreation(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto := &models.UserDto{
		User:      user,
		Timestamp: time.Now().In(time.UTC),
		TrxId:     shortuuid.New(),
	}

	printUser(*userDto)

	ctx.JSON(http.StatusCreated, "[WIP] account created")
}

func printUser(user models.UserDto) {
	log.Println("Received user: ")
	log.Println("Name: ", fmt.Sprintf("%s %s", user.User.FirstName, user.User.LastName))
	log.Println("Email: ", user.User.Email)
	log.Println("Address: ", user.User.Address)
	log.Println("Age: ", user.User.Age)
	log.Println("Time: ", user.Timestamp)
	log.Println("TrxID: ", user.TrxId)
}
