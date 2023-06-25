package dal

import (
	"Another-Nikki/dal/model"
	"errors"

	"gorm.io/gorm"
)

type Log struct{}

func (*Log)CreateLog(api string, status int, ip, resp string) error {
	log := &model.Log{
		Api: api,
		Status: status,
		IP: ip,
		Response: resp,
	}
	return DB.Model(model.Log{}).
		Create(log).Error
}

func (*Log)GetPageQue(page int) ([]model.Log, error) {
	res := make([]model.Log, 0)
	err := DB.Model(model.Log{}).
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

func (*Log)GetLogCount() (int, error) {
	var sum int64
	if err := DB.Model(model.Log{}).Count(&sum).Error; err != nil {
		return 0, err
	}
	return int(sum), nil
}