package impl

import (
	"database/sql"
	models "goadvance/repo_pattern/model"
	repo "goadvance/repo_pattern/repository"
	"log"
)

type UserRepoImpl struct {
	SQL *sql.DB
}

func NewUserRepoImpl(sql *sql.DB) repo.UserRepo {
	return &UserRepoImpl{
		SQL: sql,
	}
}

func (userRepoImpl *UserRepoImpl) Select() ([]*models.User){
	users := make([]*models.User, 0)
	stsQuery := "SELECT * FROM users"
	rows, err := userRepoImpl.SQL.Query(stsQuery)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		panic(err)
	}
	return users
}

func (userRepoImpl *UserRepoImpl) Insert(user *models.User) {
	stsCommand := `INSERT INTO users (id, name, gender, email) VALUES ($1, $2, $3, $4)`
	_, err := userRepoImpl.SQL.Exec(stsCommand, user.Id, user.Name, user.Gender, user.Email)
	if err != nil {
		log.Fatal(err)
	}
}