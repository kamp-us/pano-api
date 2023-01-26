package backend

import (
	"context"

	"github.com/kamp-us/pano-api/internal/models"
)

type Backender interface {
	// Post
	GetBatchPosts(ctx context.Context, ids []string) ([]*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
	CreatePost(ctx context.Context, args models.CreatePostArgs) (*models.Post, error)
	UpdatePost(ctx context.Context, args models.UpdatePostArgs) error
	DeletePost(ctx context.Context, id string) error
}
