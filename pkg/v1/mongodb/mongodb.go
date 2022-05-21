package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}
