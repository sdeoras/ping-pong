# Use base golang image from Docker Hub
FROM golang:1.12

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the applications and copy to /usr/local/bin.
RUN go build -o /usr/local/bin/ping-pong-client -v github.com/sdeoras/ping-pong/client/cmd/src/ping-pong-client
RUN go build -o /usr/local/bin/ping-pong-server -v github.com/sdeoras/ping-pong/server/cmd/src/ping-pong-server
