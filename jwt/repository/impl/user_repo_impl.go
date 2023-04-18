package impl

import (
	"context"
	models "goadvance/jwt/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepoImpl(db *mongo.Database) *UserRepoImpl{
	return &UserRepoImpl{
		Db: db,
	}
}

func(userRepoImpl *UserRepoImpl) InsertUser(user *models.User) error{
	userBytes, _ := bson.Marshal(user)
	_, err := userRepoImpl.Db.Collection("users").InsertOne(context.Background(), userBytes)
	if err != nil {
		return err
	}
	return nil 
}
