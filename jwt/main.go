package main

import (
	driver "goadvance/jwt/driver"
	models "goadvance/jwt/model"
	repo "goadvance/jwt/repository/impl"
)

const (
	USER     = "root"
	PASSWORD = "2292892000"
	DBNAME = "jwt"
)

func main() {
	mongo := driver.Connect(USER, PASSWORD, DBNAME)
	userRepoImpl := repo.NewUserRepoImpl(mongo.Db)
	user := models.User{
		Email: "manh@gmail.com",
		Password: "123456",
		Name: "Manh Tran",
	}
	err := userRepoImpl.InsertUser(&user)
	if err != nil {
		panic(err)
	}
}