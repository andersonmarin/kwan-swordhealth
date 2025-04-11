FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download &&  \
    go mod verify

COPY . .

RUN go build -o ./api ./cmd/api

FROM alpine:3.18

WORKDIR /app

COPY --from=0 /app/api ./

ENV PORT="80"

CMD [ "./api" ]

EXPOSE 80