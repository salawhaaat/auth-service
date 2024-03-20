package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	DB *mongo.Client
}

func (m *MongoRepository) Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	m.DB = client
}

func (m *MongoRepository) GetRefreshToken(userId string) string {
	collection := m.DB.Database("test").Collection("refreshTokens")
	filter := bson.M{"userId": userId}
	var result struct {
		RefreshToken string `bson:"refreshToken"`
	}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return "" // if there is no refresh token, return an empty string
	}
	return result.RefreshToken
}

func (m *MongoRepository) StoreRefreshToken(userId string, refresh string) error {
	collection := m.DB.Database("test").Collection("refreshTokens")
	filter := bson.M{"userId": userId}
	update := bson.M{"$set": bson.M{"refreshToken": refresh}}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}

func (m *MongoRepository) Close() error {
	err := m.DB.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
