FROM golang:1.24
EXPOSE 5000

ENV SRC_DIR=/go/src/github.com/Kaleidoscope-Backup/healthrecord-repository/
WORKDIR $SRC_DIR

# Copy source
COPY . .

# Initialize Go modules (only if not already present)
RUN go mod init github.com/Kaleidoscope-Backup/healthrecord-repository || true
RUN go mod tidy

# Install dependencies
RUN go install github.com/kevinburke/go-bindata/...@latest

# Generate schema code
RUN go generate ./schema

# Build binary
RUN go build -o health_record_repository ./cmd/health_record_repository

ENTRYPOINT ["./health_record_repository"]
