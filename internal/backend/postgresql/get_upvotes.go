package postgresql

import (
	"context"
	"github.com/kamp-us/pano-api/internal/models"
)

func (b *PostgreSQLBackend) GetUpvotes(ctx context.Context, entityId string, entityType string) ([]*models.Upvote, error) {
	var upvotes []*models.Upvote
	result := b.DB.Find(&upvotes)
	if result == nil {
		return nil, result.Error
	}

	return upvotes, nil
}