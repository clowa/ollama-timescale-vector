FROM ollama/ollama:0.5.4

COPY --chmod=0755 entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ["llama3.2:1b", "nomic-embed-text"]

# RUN ["ollama", "pull", "llama3.2:1b"]


# RUN ["ollama", "pull", "bge-large:335m-en-v1.5-fp16"]
# RUN ["ollama", "pull", "nomic-embed-text"]
