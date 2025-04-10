package nats

import (
	"errors"
	"github.com/nats-io/nats.go"
	"os"
)

func OpenNatsConnection() (*nats.Conn, error) {
	url, ok := os.LookupEnv("NATS_URL")
	if !ok {
		return nil, errors.New("NATS_URL environment variable not set")
	}

	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
