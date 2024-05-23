package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamoDBConnectionAndCreateTable(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("localhost:8000"),
	})
	assert.Nil(t, err)
	DB := dynamodb.New(sess)

	// 测试连接
	output, err := DB.ListTables(&dynamodb.ListTablesInput{})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	t.Log("Connected to DynamoDB, tables:", output.TableNames)

	// 创建表
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("MyTestTable"),
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
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err = DB.CreateTable(input)
	assert.Nil(t, err)
	t.Log("Created the table: MyTestTable")

	// 清理，删除表
	_, err = DB.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String("MyTestTable"),
	})
	assert.Nil(t, err)
	t.Log("Deleted the table: MyTestTable")
}
