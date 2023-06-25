package dal

import (
	"Another-Nikki/dal/model"

	"gorm.io/gorm"
)

type Article struct{}

func (*Article) Create(title, content, description string) error {
	article := &model.Article{
		Title: title,
		Content: content,
		Description: description,
	}
	return DB.Model(model.Article{}).Create(article).Error
}

func (*Article) GetInfoById(id int) (*model.Article, error) {
	var article model.Article
	err := DB.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func (*Article) GetList(page int) ([]model.Article, error) {
	res := make([]model.Article, 0)
	err := DB.Model(model.Article{}).
		Limit(20).
		Offset((page - 1) * 20).
		Order("ID DESC").
		Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}