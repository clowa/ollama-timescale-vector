# Overview

A containerized example app that uses LlamaIndex and Ollama to create a vector database.

This project combines several key components:

- Docker for containerization and deployment
- LlamaIndex for document indexing and retrieval
- Ollama for running local LLMs and embeddings
- TimescaleDB with [pgvector](https://github.com/pgvector/pgvector) and [pgai](https://github.com/timescale/pgai) extension as vector storage

## Prerequisites

- Docker and Docker Compose
- ~ 2GB disk space for LLM models
- Source documents to index mounted at `/data` in the app container

## Quick Start

1. Start the system:

```bash
docker compose up
```

This will:

- Start Ollama service. If the required modules haven't been downloaded they will be pulled on first execution. _(llama3.2:1b and nomic-embed-text)_
- Launch TimescaleDB for vector storage
- Start the main application that indexes your documents

### Devcontainer

If you are using VSCode, you can use the provided devcontainer to develop the application.

> [!TIP]
> The devcontainer is configured to share the host's ollama modells with the container, so the container doesn't grow huge in size. However, if you want to disable this behavior your can comment out the mount of the `~/.ollama` directory in the `devcontainer.json` file.

## Components

- `db` - Scripts to initiate the TimescaleDB with pgvector and pgai extensions
