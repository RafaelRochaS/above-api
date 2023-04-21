package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RafaelRochaS/above-api/models"
	"github.com/RafaelRochaS/above-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/RafaelRochaS/above-api/above_api"
)

func HandleAccounts(rg *gin.RouterGroup) {
	accounts := rg.Group("/accounts")

	accounts.POST("/", handleAccountCreation)
}

func handleAccountCreation(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("AccountHandler :: Error binding input: %v", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto := &models.UserDto{
		User:      user,
		Timestamp: time.Now().In(time.UTC),
		TrxId:     shortuuid.New(),
	}

	printUser(*userDto)
	err := callgRPCAccountCreation(*userDto)
	if err != nil {
		log.Printf("AccountHandler :: Error calling gRPC service: %v", err)
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error\n%v", err.Error()))
		return
	}

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

func callgRPCAccountCreation(user models.UserDto) error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", utils.ACCOUNT_SERVICE_GRPC_HOST, utils.ACCOUNT_SERVICE_GRPC_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("AccountHandler :: Error creating gRPC connection: %v", err)
		return err
	}

	defer conn.Close()

	client := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := client.CreateUser(ctx,
		&pb.UserCreateRequest{
			FirstName: user.User.FirstName,
			LastName:  user.User.LastName,
			Email:     user.User.Email,
			Address:   user.User.Address,
			Age:       int32(user.User.Age),
		})

	if err != nil {
		log.Printf("AccountHandler :: Error calling gRPC service: %v", err)
		return err
	}

	log.Printf("AccountHandler :: Server response: %v", response.GetStatus())

	if response.GetStatus() != "success" {
		return errors.New("account creation service returned failure")
	}

	return nil
}
