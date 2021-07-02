package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
	config Config
}

func NewDB(config Config) *DB {
	return &DB{
		config: config,
	}
}

type GetCollection func(database string, collection string) *mongo.Collection

func (db *DB) GetCollection(database string, collection string) *mongo.Collection {
	return db.client.Database(database).Collection(collection)
}

func (db *DB) CheckConnectionHealth(ctx context.Context) error {
	err := db.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.WithMessage(err, "pinging database with primary preference")
	}
	return nil
}
