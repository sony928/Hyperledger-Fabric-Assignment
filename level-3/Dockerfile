FROM golang:1.22-alpine
WORKDIR /app

# Copy Go module files and download dependencies first
COPY api/go.mod api/go.sum ./
RUN go mod download

# Copy the rest of the API code
COPY api/ ./api/

EXPOSE 8080
CMD ["go", "run", "api/main.go"]

