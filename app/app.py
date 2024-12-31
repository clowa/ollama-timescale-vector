from llama_index.core import Settings, VectorStoreIndex, SimpleDirectoryReader
from llama_index.llms.ollama import  Ollama
from llama_index.embeddings.ollama import OllamaEmbedding
# from llama_index.vector_stores import TimescaleVectorStore

from ollama import Client


ollama_url = "http://ollama:11434"

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
index = VectorStoreIndex.from_documents(documents)
