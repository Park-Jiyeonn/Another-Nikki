package dal

import (
	"Another-Nikki/dal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func CreateLog(ctx context.Context, api string, status int, ip, resp string) error {
	log := &model.Log{
		Api: api,
		Status: status,
		IP: ip,
		Response: resp,
	}
	return DB.WithContext(ctx).Model(model.Log{}).
		Create(log).Error
}

func GetPageQue(ctx context.Context, page int) ([]model.Log, error) {
	res := make([]model.Log, 0)
	err := DB.WithContext(ctx).
		Model(model.Log{}).
		Limit(20).
		Offset((page - 1) * 20).
		Order("ID DESC").
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func GetLogCount(ctx context.Context) (int, error) {
	var sum int64
	if err := DB.WithContext(ctx).Model(model.Log{}).Count(&sum).Error; err != nil {
		return 0, err
	}
	return int(sum), nil
}