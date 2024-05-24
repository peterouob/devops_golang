package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestDynamoDBConnectionAndCreateTable(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file :", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})
	assert.Nil(t, err)
	DB := dynamodb.New(sess)

	// 测试连接
	output, err := DB.ListTables(&dynamodb.ListTablesInput{})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	t.Log("Connected to DynamoDB, tables:", output.TableNames)

	tableId := uuid.NewString()
	// 创建表
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableId),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
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

	_, err = DB.CreateTable(input)
	assert.Nil(t, err)
	t.Log("Created the table " + tableId)

	// 清理，删除表
	_, err = DB.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableId),
	})
	assert.Nil(t, err)
	t.Log("Deleted the table" + tableId)
}
