package service

import (
	"context"
	"github.com/pramudyas69/Go-Grpc/server/domain"
	"github.com/pramudyas69/Go-Grpc/server/dto"
	"github.com/pramudyas69/Go-Grpc/server/internal/config"
	"github.com/pramudyas69/Go-Grpc/server/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
	cnf      *config.Config
}

func NewUser(useRepo domain.UserRepository, cnf *config.Config) domain.UserService {
	return &userService{
		userRepo: useRepo,
		cnf:      cnf,
	}
}

func (u userService) Register(ctx context.Context, req dto.RegisterReq) (dto.EmptyRes, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.EmptyRes{}, err
	}

	user := domain.User{
		FullName: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}

	err = u.userRepo.Insert(ctx, &user)

	return dto.EmptyRes{}, nil
}

func (u userService) Login(ctx context.Context, req dto.LoginReq) (dto.LoginRes, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.LoginRes{}, err
	}

	if user.Username == "" {
		return dto.LoginRes{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.LoginRes{}, err
	}

	token, err := util.GenerateToken(u.cnf, user.ID.Hex(), 15)
	if err != nil {
		return dto.LoginRes{}, err
	}

	return dto.LoginRes{
		Token: token,
	}, nil
}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.EmptyRes, error) {
	userID := ctx.Value("x-user")
	_, err := u.userRepo.FindByID(ctx, userID.(string))
	if err != nil {
		return dto.EmptyRes{}, err
	}

	return dto.EmptyRes{}, nil
}
