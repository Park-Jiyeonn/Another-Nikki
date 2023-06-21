package dal

import (
	"Another-Nikki/dal/model"
	"context"
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