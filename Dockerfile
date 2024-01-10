
FROM golang:1.21.5 AS dev

WORKDIR /app


COPY go.mod /app/
COPY go.sum /app/

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

EXPOSE 5050


CMD reflex -g "*.go" go run main.go --start-service
