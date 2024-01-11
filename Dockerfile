
FROM golang:1.21.5 AS dev
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH e/helltion/sdk/go1.21.5
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app


COPY go.mod /app/
COPY go.sum /app/

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

EXPOSE 5050


ENTRYPOINT ["make"]
CMD ["dev"]
