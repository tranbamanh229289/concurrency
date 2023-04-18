package repository

import (
	models "goadvance/jwt/model"
)

type UserRepo interface {
	InsertUser(user *models.User)
}