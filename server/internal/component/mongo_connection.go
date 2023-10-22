package component

import (
	"context"
	"github.com/pramudyas69/Go-Grpc/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client  *mongo.Client
	session mongo.Session
}

func NewMongoClient(cnf *config.Config) (*MongoClient, error) {
	clientOpt := options.Client().ApplyURI(cnf.Mongo.Uri)

	client, err := mongo.Connect(context.Background(), clientOpt)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client:  client,
		session: nil,
	}, nil
}

func (m *MongoClient) BeginTransaction() error {
	if m.session != nil {
		return nil
	}

	session, err := m.Client.StartSession()
	if err != nil {
		return err
	}

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	m.session = session
	return nil
}

func (m *MongoClient) CommitTransaction() error {
	if m.session == nil {
		return nil
	}

	err := m.session.CommitTransaction(context.Background())
	if err != nil {
		return err
	}

	m.session = nil
	return nil
}

func (m *MongoClient) RollbackTransaction() error {
	if m.session == nil {
		return nil
	}

	err := m.session.AbortTransaction(context.Background())
	if err != nil {
		return err
	}

	m.session = nil
	return nil
}

func (m *MongoClient) Disconnect() {
	if m.session != nil {
		m.session.EndSession(context.Background())
	}

	_ = m.Client.Disconnect(context.Background())
}
