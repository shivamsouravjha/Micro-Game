FROM golang:1.17.1-alpine3.14

WORKDIR /self/content

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8002

RUN go build

EXPOSE $PORT

CMD ["./ContentService"]