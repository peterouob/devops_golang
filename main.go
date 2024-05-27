package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/peterouob/devops_golang/database"
	"github.com/peterouob/devops_golang/services"
	"log"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", services.GetAll)
	r.GET("/:id", services.GetByID)
	r.POST("/create", services.Create)
	r.PUT("/update/:id", services.Update)
	r.DELETE("/delete/:id", services.Delete)

	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	go func() {
		if err := database.InitDynamoDB(); err != nil {
			log.Println("Error to Init DB", err.Error())
		}
	}()
	if err := godotenv.Load(); err != nil {
		log.Println("Error to load .env file:", err)
	}
	//port := ":" + os.Getenv("PORT")
	lambda.Start(HandleRequest)
	//log.Println(http.ListenAndServe(port, router.HandleRouter()))
}
