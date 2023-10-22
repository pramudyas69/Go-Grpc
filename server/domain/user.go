package domain

import (
	"context"
	"github.com/pramudyas69/Go-Grpc/server/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FullName string             `bson:"full_name"`
	Email    string             `bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type UserService interface {
	Register(ctx context.Context, req dto.RegisterReq) (dto.EmptyRes, error)
	Login(ctx context.Context, req dto.LoginReq) (dto.LoginRes, error)
	ValidateToken(ctx context.Context, token string) (dto.EmptyRes, error)
}
