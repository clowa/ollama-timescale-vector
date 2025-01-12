package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tmc/langchaingo/llms/ollama"
)

type OllamaClient struct {
	// initiated bool
	endpoint string
	LLM      *ollama.LLM
}

func NewOllamaClient() *OllamaClient {
	c := OllamaClient{}
	c.Init()

	return &c
}

// Init loads the connection details for the ollama service. It returns an error if the OLLAMA service URL is not set otherwise nil.
func (o *OllamaClient) Init() error {
	url := os.Getenv("OLLAMA_SERVICE_URL")

	// Return an error if the OLLAMA service URL is not set
	if url == "" {
		return fmt.Errorf("Environment variable OLLAMA_SERVICE_URL is not set")
	}

	o.endpoint = url
	return nil
}

func (o *OllamaClient) Ping() (bool, error) {
	// Fail if the OllamaClient hasn't been initiated
	// if !o.initiated {
	// 	return false, fmt.Errorf("OllamaClient hasn't been initiated. Call Init() first")
	// }

	// Check if the OLLAMA service is reachable
	resp, err := http.Get(o.endpoint)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Ollama service is not reachable. Ollama service returned StatusCode: %s", resp.Status)
	}

	return true, nil
}

// GetEndpoint returns the endpoint of the Ollama service.
func (o *OllamaClient) GetEndpoint() string {
	return o.endpoint
}

func (o *OllamaClient) SetLLM(model string) error {
	llm, err := ollama.New(
		ollama.WithServerURL(o.endpoint),
		ollama.WithModel(model),
	)

	if err != nil {
		return fmt.Errorf("Failed to create LLM: %w", err)
	}

	o.LLM = llm
	return nil
}
