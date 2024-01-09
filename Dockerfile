
FROM golang:1.21.5 AS dev

WORKDIR /app


COPY go.mod /app/
COPY go.sum /app/

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

EXPOSE 5050


ENTRYPOINT ["make"]
CMD ["dev"]
FROM golang:1.21.5 AS prod


WORKDIR /app


COPY go.mod go.sum .

RUN go mod download


COPY . .


EXPOSE 5050


ENTRYPOINT ["make"]
CMD ["run"]