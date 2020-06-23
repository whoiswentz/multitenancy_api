package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseMongo interface {
	connect(context.Context)
	ping(context.Context)
}

// MongoDB represents a mongo
type MongoDB struct{
	Client *mongo.Client
}

// New function construct a mongodb struct
func New(uri string) *MongoDB {
	opts := options.Client().ApplyURI(uri)

	c, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatalln(err)
	}

	return &MongoDB{Client: c}
}

// Connect receives a mongo uri and returns a mongo client
func (m *MongoDB) Connect() context.CancelFunc {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second)

	m.connect(ctx)
	m.ping(ctx)

	return cancel
}

func (m *MongoDB) connect(ctx context.Context) {
	err := m.Client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func (m *MongoDB) ping(ctx context.Context) {
	rp, _ := readpref.New(readpref.SecondaryMode)
	err := m.Client.Ping(ctx, rp)
	if err != nil {
		log.Fatalln(err)
	}
}