package main

import (
	"gormAws/db"
	"gormAws/model"
)

func main() {
	test()
}

func test() ([]model.User, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	if err != nil {
		return []model.User{}, err
	}
	db2.AutoMigrate(&model.User{})
	account := &model.User{}
	account.Email = "test"
	var users []model.User
	db2.Where(account).Find(&users)
	return users, nil
}
