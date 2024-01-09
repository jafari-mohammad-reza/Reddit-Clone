
FROM golang:1.21.5 AS dev
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /app


COPY go.mod /app/
COPY go.sum /app/

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

#EXPOSE 80/tcp


CMD reflex -c "go run ./ main.go" --start-service
