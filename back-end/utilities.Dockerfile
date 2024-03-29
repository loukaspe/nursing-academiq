FROM golang:1.19.3-alpine3.16

# Install git

RUN apk update && apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/joho/godotenv/cmd/godotenv@v1.4.0
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

# Run tests
CMD CGO_ENABLED=0 go test -v  ./...