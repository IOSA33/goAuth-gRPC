FROM golang:1.24-alpine

# Downloading SQLite dependencies
RUN apk add --no-cache gcc musl-dev

# This is folder where all our application in container
WORKDIR /app

# Copying dependencies from go.mod
COPY go.mod go.sum ./
RUN go mod download

# Copying whole code
COPY . .

# Making up our application
RUN go build -o migrator ./cmd/migrator

# Creating folder for out database
RUN mkdir -p /app/storage

ENTRYPOINT ["./migrator"]