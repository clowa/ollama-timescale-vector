package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
)

type Embedding struct {
	ID        string `gorm:"type:uuid;primaryKey;not null"`
	Md5       string `gorm:"uniqueIndex"`
	Text      string
	CreatedAt time.Time       `gorm:"<-:false"`
	Embedding pgvector.Vector `gorm:"type:vector(768);index:embedding_idx,type:hnsw"` // Nomic Embed v1.5 embeddings have 64 to 768 dimensions
}

func NewEmbedding(data []byte) *Embedding {
	embed := &Embedding{
		ID:        uuid.NewString(),
		Text:      string(data),
		CreatedAt: time.Now(),
	}

	md5Sum := md5.Sum(data)
	embed.Md5 = hex.EncodeToString(md5Sum[:])

	return embed
}

func (e *Embedding) Exists(ctx context.Context, r *Repository) (bool, error) {
	var count int64
	err := r.db.Model(&Embedding{}).Where("md5 = ?", e.Md5).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("cannot check existence of embedding: %w", err)
	}

	return count > 0, nil
}

// PerformTextEmbedding creates an embedding for the text using the ollama service.
func (e *Embedding) PerformTextEmbedding(ctx context.Context, ollamaClient *OllamaClient) error {
	results, err := ollamaClient.LLM.CreateEmbedding(ctx, []string{e.Text})
	if err != nil {
		return fmt.Errorf("cannot embed text: %w", err)
	}

	e.CreatedAt = time.Now()
	e.Embedding = pgvector.NewVector(results[0])
	return nil
}

func (e *Embedding) String() string {
	return fmt.Sprintf("UUID: %s, Md5: %s, Text: %s, CreatedAt: %s, Embedding: %v", e.ID, e.Md5, e.Text, e.CreatedAt, e.Embedding)
}
