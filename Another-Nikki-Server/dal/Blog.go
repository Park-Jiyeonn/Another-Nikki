package dal

import (
	"Another-Nikki/dal/model"
	"context"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func GetBlogList(ctx context.Context) ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	err := DB.WithContext(ctx).
		Model(model.Blog{}).
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func GetLastSevenBlog(ctx context.Context, num int64) ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	var sum int64
	if err := DB.WithContext(ctx).Model(model.Blog{}).Count(&sum).Error; err != nil {
		return nil, err
	}
	
	err := DB.WithContext(ctx).
		Model(model.Blog{}).
		Limit(int(num)).
		Offset(int(sum - num)).
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func GetRandomBlog(ctx context.Context) ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	var sum int64
	if err := DB.WithContext(ctx).Model(model.Blog{}).Count(&sum).Error; err != nil {
		return nil, err
	}
	
	r := rand.New(rand.NewSource(time.Now().Unix()))
	num := r.Intn(int(sum))

	err := DB.WithContext(ctx).
		Model(model.Blog{}).
		Limit(1).
		Offset(num).
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func CreateBolg(ctx context.Context, content string) error {
	blog := &model.Blog{
		Content: content,
	}
	return DB.WithContext(ctx).Model(model.Blog{}).
		Create(blog).Error
}

func DeleteBolgByID(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).Model(&model.Blog{}).
		Where("id = ?", id).
		Delete(nil).Error
}
