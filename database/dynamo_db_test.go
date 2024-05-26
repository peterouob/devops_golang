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
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// 验证环境变量是否正确加载
	log.Println("AWS_REGION:", os.Getenv("AWS_REGION"))
	log.Println("AWS_ACCESS_KEY_ID:", os.Getenv("AWS_ACCESS_KEY_ID"))
	log.Println("AWS_SECRET_ACCESS_KEY:", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	// 初始化 AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	DB := dynamodb.New(sess)

	// 列出表
	output, err := DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("Failed to list tables: %v", err)
	}
	assert.NotNil(t, output)
	t.Log("Connected to DynamoDB, tables:", output.TableNames)

	// 创建表
	tableId := uuid.NewString()
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

	// 检查创建表的返回错误
	_, err = DB.CreateTable(input)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	t.Log("Created the table " + tableId)

	// 删除表
	//_, err = DB.DeleteTable(&dynamodb.DeleteTableInput{
	//	TableName: aws.String(tableId),
	//})
	//if err != nil {
	//	log.Fatalf("Failed to delete table: %v", err)
	//}
	//t.Log("Deleted the table " + tableId)
}
