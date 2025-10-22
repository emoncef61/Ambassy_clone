package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repositories struct {
	UserRepository *UserRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
