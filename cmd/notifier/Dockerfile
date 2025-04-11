FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download &&  \
    go mod verify

COPY . .

RUN go build -o ./notifier ./cmd/notifier

FROM alpine:3.18

WORKDIR /app

COPY --from=0 /app/notifier ./

CMD [ "./notifier" ]