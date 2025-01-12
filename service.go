package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/clowa/ollama-timescale-vector/utils"
)

type Service struct {
	repository *utils.Repository
	ollama     *utils.OllamaClient
	logger     utils.Logger
	// options    ServiceOption
}

// type ServiceOption struct {
// 	EmbeddingModel string
// 	LLMModel       string
// }

func NewService(repository *utils.Repository, ollama *utils.OllamaClient, logger utils.Logger) *Service {
	return &Service{
		repository: repository,
		ollama:     ollama,
		logger:     logger,
		// options:    options,
	}
}

// func (s *Service) InitiateDatabase() error {

// }

func (s *Service) CreateFileEmbedding(path string) error {
	// Index the file
	s.logger.LogInfo("Indexing file: ", path)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	// Read the file in chunks of max prompt length
	ctx := context.Background()
	r := bufio.NewReader(f)

	const maxPromptSize = 8192 * 4
	buf := make([]byte, maxPromptSize)
	for {
		n, err := r.Read(buf)

		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading file: %w", err)
		}
		if n == 0 || err == io.EOF {
			break
		}

		buf := buf[:n]

		// Read the next chunk until a newline
		nextUntillNewline, err := r.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		// Create an embedding for the chunk if the md5 hash doesn't exist in the database
		s.logger.LogInfo("Processing chunk of size: ", len(buf))
		embed := utils.NewEmbedding(buf)
		if exists, err := embed.Exists(ctx, s.repository); err != nil {
			return fmt.Errorf("cannot check existence of embedding: %w", err)
		} else if exists {
			s.logger.LogInfo("Embedding already exists in the database")
			continue
		}

		err = embed.PerformTextEmbedding(ctx, s.ollama)
		if err != nil {
			return fmt.Errorf("cannot embed text: %w", err)
		}

		// Store the embedding in the database
		err = s.repository.StoreEmbeddingsInDB(ctx, *embed)
		if err != nil {
			return fmt.Errorf("cannot store embeddings in db: %w", err)
		}
	}
	return nil
}
