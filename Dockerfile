
FROM golang:1.21.5 AS dev

RUN export PATH=$PATH:/bin:/usr/bin

ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH=/home/helltion/go:/home/helltion/sdk/go1.21.5
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app


COPY go.mod go.sum /home/helltion/Reddit-Clone/
RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest
# RUN go build -o /gateway-service
# Expose port 5051
EXPOSE 5050



CMD reflex -s -r '\.go' -R '^vendor/.' -R '^_.*' go run main.go
