package user

import (
	"context"
	"main/internal/entity"
)

type RepoInterface interface {
	GetListUsers(ctx context.Context) ([]entity.User, error)
	GetUserByUserid(ctx context.Context, id string) (entity.User, error)
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
