package mongo

import (
	"context"

	"github.com/pkg/errors"
)

func (db *DB) Disconnect(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
