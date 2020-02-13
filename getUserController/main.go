package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yunuskilicdev/gormAws/v2/db"
	"github.com/yunuskilicdev/gormAws/v2/model"
)

type GetUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func HandleRequest(ctx context.Context, request GetUserRequest) ([]model.User, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return []model.User{}, err
	}
	db2.AutoMigrate(&model.User{})
	account := &model.User{}
	if request.Email != "" {
		account.Email = request.Email
	}
	if request.Name != "" {
		account.Name = request.Name
	}
	var users []model.User
	db2.Where(account).Find(&users)
	return users, nil
}
func main() {
	lambda.Start(HandleRequest)
}
