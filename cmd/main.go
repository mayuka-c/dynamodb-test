package main

import (
	"context"

	"github.com/mayuka-c/dynamodb-test/internal/pkg/config"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/models"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/service"
)

func main() {
	dbConfig, err := config.GetDBConfig()
	if err != nil {
		panic(err)
	}
	service := service.NewService(dbConfig)

	content := models.Content{
		UserName:     "Mango",
		PhoneNumber:  "1234434242",
		Address:      "Mandya",
		Organization: "Avast",
	}

	service.CreateItem(context.Background(), content)
}
