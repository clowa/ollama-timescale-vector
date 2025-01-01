FROM python:3.13-slim-bookworm

ENV DEBIAN_FRONTEND=noninteractive

## Required for some pip packages to build
RUN apt-get update \
        && apt-get -yqq install \
                libpq-dev \
                build-essential \
        && rm -rf /var/lib/apt/lists/*
        
RUN mkdir /app
WORKDIR /app

COPY . .
RUN pip install -r requirements.txt

CMD ["python", "app.py"]
