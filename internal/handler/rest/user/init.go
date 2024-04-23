package user

import (
	"context"
	"main/internal/entity"
)

type UcInterface interface {
	DisplayAllUsers(ctx context.Context) ([]entity.User, error)
	DisplayUser(ctx context.Context, id string) (entity.User, error)
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
