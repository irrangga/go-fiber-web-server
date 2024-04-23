package user

import (
	"context"
	"main/internal/entity"
)

func (r Repository) GetListUsers(ctx context.Context) ([]entity.User, error) {
	var users []User
	var usersEntity []entity.User

	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return []entity.User{}, err
	}
	for _, user := range users {
		usersEntity = append(usersEntity, entity.User{
			Userid: user.Userid,
			Name:   user.Name,
		})
	}
	return usersEntity, nil
}

func (r Repository) GetUserByUserid(ctx context.Context, userid string) (entity.User, error) {
	var user User

	err := r.db.WithContext(ctx).First(&user, userid).Error
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		Userid: user.Userid,
		Name:   user.Name,
	}, nil
}
