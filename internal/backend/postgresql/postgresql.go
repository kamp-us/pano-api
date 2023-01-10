package postgresql

// TODO: ask umut, gorm association, cascade?

import (
	"context"
	"time"

	"github.com/gosimple/slug"
	"github.com/kamp-us/pano-api/internal/backend"
	"github.com/kamp-us/pano-api/internal/models"
	"gorm.io/gorm"
)

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) backend.Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}

func (b *PostgreSQLBackend) GetBatchPosts(ctx context.Context, ids []string) ([]*models.Post, error) {
	var posts []*models.Post
	result := b.DB.Find(&posts, ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (b *PostgreSQLBackend) GetPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	result := b.DB.Find(&posts)
	if result == nil {
		return nil, result.Error
	}

	return posts, nil
}

func (b *PostgreSQLBackend) CreatePost(ctx context.Context, title string, url string, content string, userId string) (*models.Post, error) {
	slug := slug.MakeLang(title, "tr")

	post := models.Post{
		Title:   title,
		Url:     url,
		Content: content,
		Slug:    slug,
		UserID:  userId,
	}

	result := b.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (b *PostgreSQLBackend) UpdatePost(ctx context.Context, id string, title *string, url *string, content *string) error {
	post := models.Post{}
	result := b.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	slug := slug.MakeLang(*title, "tr")

	updates := models.Post{
		Title:   *title,
		Url:     *url,
		Content: *content,
		Slug:    slug,
	}

	result = b.DB.Model(&post).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *PostgreSQLBackend) DeletePost(ctx context.Context, id string) error {
	post := models.Post{}
	result := b.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// TODO: we should seperate at least on folder level backend/comments,  backend/posts
// TODO: why do we return array of pointers?
// Comment
func (b *PostgreSQLBackend) GetBatchComments(ctx context.Context, ids []string) ([]*models.Comment, error) {
	// TODO: implement
	var comments []*models.Comment
	result := b.DB.Find(comments, ids)

	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (b *PostgreSQLBackend) GetComments(ctx context.Context) ([]*models.Comment, error) {
	// TODO: implement
	var comments []*models.Comment
	result := b.DB.Find(comments)

	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (b *PostgreSQLBackend) CreateComment(ctx context.Context, content string, postId string, userId string, parentId *string, deletedAt *time.Time) (*models.Comment, error) {
	comment := models.Comment{
		Content: content,
		PostID:  postId,
		UserID:  userId,
	}

	// TODO: is this the way?
	if parentId != nil {
		comment.ParentID = *parentId
	}

	if deletedAt != nil {
		comment.DeletedAt = *deletedAt
	}

	result := b.DB.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func (b *PostgreSQLBackend) UpdateComment(ctx context.Context, id string, content string) error {
	comment := models.Comment{}

	result := b.DB.First(&comment, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	updates := models.Comment{
		Content: content,
	}

	result = b.DB.Model(&comment).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *PostgreSQLBackend) DeleteComment(ctx context.Context, id string) error {
	comment := models.Comment{}

	result := b.DB.First(&comment, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
