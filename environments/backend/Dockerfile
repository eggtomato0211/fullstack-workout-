FROM golang:1.24-alpine

WORKDIR /app

# Install tools
RUN apk add --no-cache git curl gcc musl-dev

# Install Air (hot reload)
ENV GOTOOLCHAIN=auto
RUN go install github.com/air-verse/air@latest

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Default command: run with Air hot reload
CMD ["sh", "-c", "if [ -f go.mod ]; then air; else echo 'No Go project found. Create one with: go mod init <module-name>' && tail -f /dev/null; fi"]
