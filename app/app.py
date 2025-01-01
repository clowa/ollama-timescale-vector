import os

from llama_index.core import Settings, VectorStoreIndex, SimpleDirectoryReader, StorageContext
from llama_index.llms.ollama import  Ollama
from llama_index.embeddings.ollama import OllamaEmbedding
# from llama_index.vector_stores import TimescaleVectorStore

from ollama import Client
from llama_index.vector_stores.postgres import PGVectorStore
from sqlalchemy import make_url


postgres_service_url = os.environ["POSTGRES_SERVICE_URL"]
postgres = make_url(postgres_service_url)
ollama_url = os.environ["OLLAMA_SERVICE_URL"]

data_directory = "/data"
ollama_model = "llama3.2:1b"
embed_model = "nomic-embed-text:v1.5"

# Ensure that the models are available
ollama_client = Client(ollama_url)
ollama_client.pull(ollama_model)
ollama_client.pull(embed_model)

# Configure the LLM used by lama-index
Settings.llm = Ollama(
    model=ollama_model,
    base_url=ollama_url,
    request_timeout=120,
    # temperature=0.1,
)

# Configure the embedding model used by lama-index
Settings.embed_model = OllamaEmbedding(
    # model_name="bge-large:335m-en-v1.5-fp16",
    model_name=embed_model,
    base_url=ollama_url,
    ollama_additional_kwargs={"mirostat": 0}
)

# Create embedding index
loader = SimpleDirectoryReader(
    input_dir=data_directory,
    required_exts=[".go", ".md"],
    recursive=True,
)

documents = loader.load_data()

# Create a TimescaleVectorStore instance
vector_store = PGVectorStore.from_params(
    host=postgres.host,
    port=postgres.port,
    user=postgres.username,
    password=postgres.password,
    database=postgres.database,
    table_name="documents",
    embed_dim=768,
    hnsw_kwargs={
        "hnsw_m": 16,
        "hnsw_ef_construction": 64,
        "hnsw_ef_search": 40,
        "hnsw_dist_method": "vector_cosine_ops",
    },
)

# Create a new VectorStoreIndex using the TimescaleVectorStore
storage_context = StorageContext.from_defaults(vector_store=vector_store)
index = VectorStoreIndex.from_documents(
    documents, storage_context=storage_context
)
