package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/clowa/ollama-timescale-vector/utils"
)

const (
	embeddingModel           = "nomic-embed-text:v1.5"
	embeddingMaxPromptLength = 8192
	llmModel                 = "llama3.2:1b"

	POSTGRES_HOST     = "db"
	POSTGRES_USER     = "dbuser"
	POSTGRES_PASSWORD = "U-5afgda3"
	POSTGRES_DB       = "llama-vector"
)

func main() {
	// Parse command line arguments
	dataDir := flag.String("directory", "", "Directory containing the data to index")
	flag.Parse()

	// Prepare the logger
	logger := utils.Logger{}
	logger.LoggerInit()

	// Prepare the repository
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", POSTGRES_HOST, 5432, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogError(err.Error())
		return
	}
	repository := utils.NewRepository(db)
	repository.InitiateDatabase()

	// Prepare the OLLAMA service
	ollamaClient := utils.NewOllamaClient()
	utils.Must(ollamaClient.Ping())
	ollamaClient.SetLLM(embeddingModel)

	service := NewService(repository, ollamaClient, logger)

	////////////////////////////////////////
	indexer := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return service.CreateFileEmbedding(path)
	}
	////////////////////////////////////////

	err = filepath.Walk(*dataDir, indexer)
	if err != nil {
		panic(err)
	}
}
