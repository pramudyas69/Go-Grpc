package repository

import (
	"context"
	"github.com/pramudyas69/Go-Grpc/server/domain"
	"github.com/pramudyas69/Go-Grpc/server/internal/component"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	mongoClient *component.MongoClient
}

func NewUser(mongoClient *component.MongoClient) domain.UserRepository {
	return &userRepository{
		mongoClient: mongoClient,
	}
}

func (u userRepository) FindByID(ctx context.Context, id string) (user domain.User, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	err = u.mongoClient.Client.Database("grpc").Collection("users").FindOne(ctx, bson.M{
		"_id": objectID,
	}).Decode(&user)

	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = u.mongoClient.Client.Database("grpc").Collection("users").FindOne(ctx, bson.M{
		"username": username,
	}).Decode(&user)

	return
}

func (u userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := u.mongoClient.BeginTransaction()
	if err != nil {
		return err
	}
	_, err = u.mongoClient.Client.Database("grpc").Collection("users").
		InsertOne(ctx, user)
	defer u.mongoClient.RollbackTransaction()

	err = u.mongoClient.CommitTransaction()
	if err != nil {
		return err
	}

	return err
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
