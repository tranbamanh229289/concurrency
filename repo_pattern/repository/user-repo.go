package repository

import (
	models "goadvance/repo_pattern/model"
)

type UserRepo interface {
	Select()([]*models.User)
	Insert(user *models.User)
}