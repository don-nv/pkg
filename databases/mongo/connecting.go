package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.config.URI))
	if err != nil {
		return errors.WithMessage(err, "connecting to URI")
	}
	db.client = client

	err = db.CheckConnectionHealth(ctx)
	if err != nil {
		return errors.WithMessage(err, "checking connection health")
	}

	return nil
}
