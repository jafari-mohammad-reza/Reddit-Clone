
FROM golang:1.21.5 AS dev

ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app

#/home/helltion/Reddit-Clone/
COPY go.mod go.sum ./
RUN go mod download




COPY . .
RUN go install github.com/cespare/reflex@latest
# RUN go build -o /gateway-service
# Expose port 5051
EXPOSE 5050



CMD reflex -s -r '\.go' -R '^vendor/.' -R '^_.*' go run main.go
