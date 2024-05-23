package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/peterouob/devops_golang/database"
	"github.com/peterouob/devops_golang/router"
	"log"
)

func main() {
	r := gin.Default()
	go func() {
		if err := database.InitDynamoDB(); err != nil {
			log.Println("Error to Init DB", err.Error())
		}
	}()
	router.HandleRouter(r)
	ginLambda := ginadapter.New(r)
	lambda.Start(ginLambda.Proxy)
	r.Run(":80")
}
