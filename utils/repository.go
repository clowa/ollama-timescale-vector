package utils

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InitiateDatabase() error {
	// Migrate the schema
	err := r.db.AutoMigrate(&Embedding{})
	if err != nil {
		return fmt.Errorf("cannot migrate schema: %w", err)
	}

	return nil
}

// StoreEmbeddingsInDB writes the embeddings to the database.
func (r *Repository) StoreEmbeddingsInDB(ctx context.Context, embedding Embedding) error {
	result := r.db.Create(&embedding)

	if result.Error != nil {
		return fmt.Errorf("cannot store embeddings in db currently: %w", result.Error)
	}

	return nil
}
