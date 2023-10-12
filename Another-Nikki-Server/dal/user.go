package dal

import (
	"Another-Nikki/dal/model"
	"errors"

	"gorm.io/gorm"
)

type User struct{}

func (*User) CreateUser(username, password string) (int, error) {
	user := &model.User{
		Username: username,
		Password: password,
	}
	err := DB.Model(model.User{}).Create(user).Error
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (User) GetUserById(userId int64) (*model.User, error) {
	var user model.User
	err := DB.Model(model.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, err
	}
	return &user, nil
}

func (User) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := DB.Model(model.User{}).Where("username = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}

func (User) UpdateUser(uid int64, userMap map[string]interface{}) (err error) {
	if err = DB.Model(model.User{}).Where("id = ?", uid).Updates(&userMap).Error; err != nil {
		return err
	}
	return nil
}
