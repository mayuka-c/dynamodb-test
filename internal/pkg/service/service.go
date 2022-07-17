package service

import (
	"context"
	"fmt"

	"github.com/mayuka-c/dynamodb-test/internal/pkg/config"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/db"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/models"
	"github.com/mayuka-c/dynamodb-test/internal/pkg/utils"
)

type Service struct {
	db db.DBHandlers
}

func NewService(dbConfig config.DBConfiguration) Service {
	return Service{
		db: db.NewDatabase(dbConfig),
	}
}

func (s *Service) CreateItem(ctx context.Context, content models.Content) {
	fmt.Println("Calling db handlers from service layer")

	out, err := s.db.CreateItem(ctx, content)
	if err != nil {
		panic(err)
	}

	outputContent := models.Content{}
	utils.DeepCopy(out, &outputContent)

	fmt.Println("Output Content from Service layer", outputContent)
}
