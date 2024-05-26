package main

import (
	"github.com/joho/godotenv"
	"github.com/peterouob/devops_golang/database"
	"github.com/peterouob/devops_golang/router"
	"net/http"

	"log"
	"os"
)

func main() {
	go func() {
		if err := database.InitDynamoDB(); err != nil {
			log.Println("Error to Init DB", err.Error())
		}
	}()
	if err := godotenv.Load(); err != nil {
		log.Println("Error to load .env file:", err)
	}
	port := ":" + os.Getenv("PORT")
	//log.Println(gateway.ListenAndServe(port, router.HandleRouter()))
	log.Println(http.ListenAndServe(port, router.HandleRouter()))
}
