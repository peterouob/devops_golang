package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/peterouob/devops_golang/database"
	"github.com/peterouob/devops_golang/model"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(database.TableName),
	}
	result, err := database.DB.Scan(input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":  err,
			"code": -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result.Items,
		"code":   0,
	})
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	input := &dynamodb.GetItemInput{
		TableName: aws.String(database.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(id)},
		},
	}
	result, err := database.DB.GetItem(input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":  err,
			"code": -1,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result.Item,
		"code":   0,
	})
}

func Create(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error to bind JSON ", err)
		return
	}
	user.ID = uuid.NewString()
	item := &dynamodb.PutItemInput{
		TableName: aws.String(database.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"ID":       {S: aws.String(user.ID)},
			"username": {S: aws.String(user.Username)},
			"password": {S: aws.String(user.Password)},
		},
	}
	if _, err := database.DB.PutItem(item); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":  err.Error(),
			"code": -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 0,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err :": err.Error(),
			"code":  -1,
		})
		return
	}
	update := map[string]*dynamodb.AttributeValueUpdate{
		"username": {
			Action: aws.String("PUT"),
			Value:  &dynamodb.AttributeValue{S: aws.String(user.Username)},
		},
		"password": {
			Action: aws.String("PUT"),
			Value:  &dynamodb.AttributeValue{S: aws.String(user.Password)},
		},
	}
	updateItem := &dynamodb.UpdateItemInput{
		TableName: aws.String(database.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(id)},
		},
		AttributeUpdates: update,
	}
	if _, err := database.DB.UpdateItem(updateItem); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err :": err.Error(),
			"code":  -1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 0,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(database.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(id)},
		},
	}
	if _, err := database.DB.DeleteItem(deleteInput); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err :": err.Error(),
			"code":  -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 0,
	})
}
