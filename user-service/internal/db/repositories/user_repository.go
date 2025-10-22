package repositories

import (
	"auth-service/internal/api/models"
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepositoryInterface interface {
	Create(*models.User) error
}

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) error {
	coll := ur.db.Collection("users")
	_, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
