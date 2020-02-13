package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yunuskilicdev/gormAws/v2/db"
	"github.com/yunuskilicdev/gormAws/v2/model"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func HandleRequest(ctx context.Context, request CreateUserRequest) (model.User, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return model.User{}, err
	}
	db2.AutoMigrate(&model.User{})
	account := &model.User{}
	account.Email = request.Email
	account.Name = request.Name
	db2.Create(account)
	return *account, nil
}
func main() {
	lambda.Start(HandleRequest)
}
