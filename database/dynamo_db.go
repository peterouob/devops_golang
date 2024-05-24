package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *dynamodb.DynamoDB

const TableName = "user"

func InitDynamoDB() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file :", err)
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})
	if err != nil {
		return err
	}
	DB = dynamodb.New(sess)

	//test connection
	if _, err := DB.ListTables(&dynamodb.ListTablesInput{}); err != nil {
		return err
	}
	log.Println("Connect success")

	if err := createTable(); err != nil {
		return err
	}
	return nil
}

func createTable() error {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(TableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			// 只需要定義主鍵即可
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
	if _, err := DB.CreateTable(input); err != nil {
		if err.Error() == dynamodb.ErrCodeTableAlreadyExistsException {
			return nil
		}
		return err
	}
	log.Println("Success create table")
	return nil
}
