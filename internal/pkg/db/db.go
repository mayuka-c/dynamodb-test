package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/config"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/models"
)

const (
	personalSortKey = "personal_username"
)

type DBHandlers interface {
	CreateItem(ctx context.Context, content models.Content) (*models.Content, error)
}

// Database database module structure
type Database struct {
	dynamo    dynamodbiface.DynamoDBAPI
	tableName string
}

// NewDatabase creates a new database object with aws session
func NewDatabase(dbConfig config.DBConfiguration) *Database {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(dbConfig.Region),
		Credentials: credentials.NewStaticCredentials(dbConfig.AccessKey, dbConfig.SecretKey, ""),
		Endpoint:    aws.String(dbConfig.EndpointURL),
	})

	if err != nil {
		panic(err)
	}

	return &Database{
		tableName: dbConfig.TableName,
		dynamo:    dynamodb.New(sess),
	}
}

// CreateItem - create item in db
func (d Database) CreateItem(ctx context.Context, content models.Content) (*models.Content, error) {

	fmt.Println("Creating Item in db")

	content.SortKey = personalSortKey
	contentItem, _ := dynamodbattribute.MarshalMap(content)

	input := &dynamodb.PutItemInput{
		Item:      contentItem,
		TableName: aws.String(d.tableName),
	}

	_, err := d.dynamo.PutItem(input)
	if err != nil {
		return &content, err
	}

	fmt.Println("Successfully created an item in db")

	return &content, nil
}
