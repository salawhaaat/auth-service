package database

import (
	"github.com/salawhaaat/auth-service/internal/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository defines the interface for a database repository.
type Repository interface {
	Connect()
	GetRefreshToken(userId string) string
	StoreRefreshToken(userId string, refreshToken string) error
	Close() error
}

// NewMongoRepository creates a new MongoRepository.
func NewMongoRepository() Repository {
	var mongo mongo.Client
	return &mongodb.MongoRepository{DB: &mongo}
}
