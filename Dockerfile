
FROM golang:1.21.4 AS dev
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
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
FROM golang:1.21.4 AS prod
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app


COPY go.mod  .
COPY go.sum . 

RUN go mod download


COPY . .


EXPOSE ${PORT}


ENTRYPOINT ["make"]
CMD ["run"]