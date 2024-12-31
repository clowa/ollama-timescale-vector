FROM python:3.8-slim

RUN pip install \
        ollama

RUN pip install \
        llama-index \
        llama-index-core \
        llama-index-llms-ollama \
        llama-index-embeddings-ollama
        
RUN pip install \
        openpyxl # for reading excel files

RUN mkdir /app
WORKDIR /app
COPY . .

CMD ["python", "app.py"]
