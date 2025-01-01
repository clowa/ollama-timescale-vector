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

## Components

- `app` - Python application that handles document processing and indexing
