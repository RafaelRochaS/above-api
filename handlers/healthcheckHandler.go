package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/RafaelRochaS/above-api/above_api"
	"github.com/RafaelRochaS/above-api/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HealthStatus struct {
	AccountService string
}

func HandleHealthcheck(rg *gin.RouterGroup) {
	healthcheck := rg.Group("/health")

	healthcheck.GET("/", func(ctx *gin.Context) {
		var generalStatus = &HealthStatus{}
		status, err := healthcheckAccountService()
		if err != nil {
			generalStatus.AccountService = "unavailable"
		} else {
			generalStatus.AccountService = status
		}

		ctx.JSON(http.StatusOK, generalStatus)
	})
}

func healthcheckAccountService() (string, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", utils.ACCOUNT_SERVICE_GRPC_HOST, utils.ACCOUNT_SERVICE_GRPC_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("HealthcheckHandler :: Error creating gRPC connection: %v", err)
		return "", err
	}

	defer conn.Close()

	client := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := client.Healthcheck(ctx, &pb.HealthcheckRequest{})

	if err != nil {
		log.Printf("HealthcheckHandler :: Error creating calling account service: %v", err)
		return "", err
	}

	return response.GetStatus(), nil
}
