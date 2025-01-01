FROM python:3.13-slim-bookworm

ENV DEBIAN_FRONTEND=noninteractive

RUN pip install \
        ollama

RUN pip install \
        llama-index \
        llama-index-core \
        llama-index-llms-ollama \
        llama-index-embeddings-ollama

RUN apt-get update \
        && apt-get -yqq install \
                libpq-dev \
                build-essential \
        && pip install \
                llama-index-vector-stores-postgres \
                sqlalchemy \
        && rm -rf /var/lib/apt/lists/*
        
RUN pip install \
        openpyxl # for reading excel files

RUN mkdir /app
WORKDIR /app
COPY . .

CMD ["python", "app.py"]
