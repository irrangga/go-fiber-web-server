package user

import (
	"context"
	"main/internal/entity"
)

func (uc Usecase) DisplayAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := uc.repo.GetListUsers(ctx)
	if err != nil {
		return []entity.User{}, err
	}
	return users, nil
}

func (uc Usecase) DisplayUser(ctx context.Context, userid string) (entity.User, error) {
	user, err := uc.repo.GetUserByUserid(ctx, userid)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
