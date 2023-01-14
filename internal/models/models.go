package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string
	Url     string
	Content string
	Slug    string

	UserID string
}

type Upvote struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	EntityID   string
	EntityType     string

	UserID string
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&Post{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Upvote{}); err != nil {
		return err
	}
	return nil
}
