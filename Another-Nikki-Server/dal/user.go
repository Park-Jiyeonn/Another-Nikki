package dal

import (
	"Another-Nikki/dal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, username, password string) (int, error) {
	user := &model.User{
		Username: username,
		Password: password,
	}
	err := DB.WithContext(ctx).Model(model.User{}).Create(user).Error
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func GetUserById(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Model(model.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, err
	}
	return &user, nil
}

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Model(model.User{}).Where("username = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, uid int64, userMap *map[string]interface{}) (err error) {
	if err = DB.WithContext(ctx).Model(model.User{}).Where("id = ?", uid).Updates(&userMap).Error; err != nil {
		return err
	}
	return nil
}
