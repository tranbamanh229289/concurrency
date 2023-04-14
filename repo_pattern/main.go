package main

import (
	"fmt"
	driver "goadvance/repo_pattern/driver"
	models "goadvance/repo_pattern/model"
	repoImpl "goadvance/repo_pattern/repository/impl"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "2292892000"
	dbname = "gorm"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect Success !")

	userRepo := repoImpl.NewUserRepoImpl(db.SQL)
	user1 := models.User{
		Id: 1,
		Name: "Manh Tran",
		Gender: "Male",
		Email: "manhtran@gmail.com",
	}

	user2 := models.User{
		Id: 2,
		Name: "Hue Le",
		Gender: "Female",
		Email: "huele@gmail.com",
	}

	userRepo.Insert(&user1)
	userRepo.Insert(&user2)
	
	users := userRepo.Select()

	for _, v := range users {
		fmt.Println(v)
	}

}