package repositories

import (
	"auth-service/internal/api/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ur *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.Metadata = models.Metadata{
		CreatedAt:     time.Now(),
		LastUpdated:   time.Now(),
		AccountStatus: "active",
		Active:        true,
	}
	coll := ur.db.Collection("users")
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
