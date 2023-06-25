package dal

import (
	"Another-Nikki/dal/model"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Comment struct{}

var rx *rand.Rand	// 给 GetRandomBlog 使用，原来的写法是每次创建对象，总感觉这样浪费资源，现在这样只创建一次

func (*Comment) GetBlogList() ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	err := DB.Model(model.Blog{}).
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func (*Comment) GetLastSevenBlog(num int64) ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	var sum int64
	if err := DB.Model(model.Blog{}).Count(&sum).Error; err != nil {
		return nil, err
	}
	
	err := DB.Model(model.Blog{}).
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

func (*Comment) GetRandomBlog() ([]model.Blog, error) {
	res := make([]model.Blog, 0)
	var sum int64
	if err := DB.Model(model.Blog{}).Count(&sum).Error; err != nil {
		return nil, err
	}
	
	if rx == nil {
		rx = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	num := rx.Intn(int(sum))

	err := DB.Model(model.Blog{}).
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

func (*Comment) CreateBolg(content string) error {
	blog := &model.Blog{
		Content: content,
	}
	return DB.Model(model.Blog{}).
		Create(blog).Error
}

func (*Comment) DeleteBolgByID(id int64) error {
	return DB.Model(&model.Blog{}).
		Where("id = ?", id).
		Delete(nil).Error
}
