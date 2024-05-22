package main

import (
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
	r.Run(":80")
}
