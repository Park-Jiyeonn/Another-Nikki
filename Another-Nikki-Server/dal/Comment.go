package dal

import (
	"Another-Nikki/dal/model"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Comment struct{}

var rx *rand.Rand // 给 GetRandomComment 使用，原来的写法是每次创建对象，总感觉这样浪费资源，现在这样只创建一次

func getChildren(root uint) []model.Comment {
	res := make([]model.Comment, 0)
	DB.Model(model.Comment{}).
		Where("root_id = ?", root). // 从来没写过这种不等于
		Find(&res)
	return res
}

func (*Comment) GetCommentList() ([]model.Comment, error) {
	res := make([]model.Comment, 0)
	err := DB.Model(model.Comment{}).
		Where("root_id = 0"). // 首先查出顶级评论，其父亲 id 为 0
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	for i, _ := range res {
		res[i].Children = getChildren(res[i].ID)
	}
	return res, nil
}

func (*Comment) GetLastSevenComment(num int64) ([]model.Comment, error) {
	res := make([]model.Comment, 0)
	var sum int64
	if err := DB.Model(model.Comment{}).
		Where("root_id = 0").
		Count(&sum).Error; err != nil {
		return nil, err
	}

	err := DB.Model(model.Comment{}).
		Where("root_id = 0").
		Limit(int(num)).
		Offset(int(sum - num)).
		Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	for i, _ := range res {
		res[i].Children = getChildren(res[i].ID)
	}
	return res, nil
}

func (*Comment) GetRandomComment() ([]model.Comment, error) {
	res := make([]model.Comment, 0)
	var sum int64
	if err := DB.Model(model.Comment{}).Count(&sum).Error; err != nil {
		return nil, err
	}

	if rx == nil {
		rx = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	num := rx.Intn(int(sum))

	err := DB.Model(model.Comment{}).
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

func (*Comment) CreateComment(content, AuthorName, ParentName, AuthorAvatar string, AuthorID, rootID, ParentID int) error {
	comment := &model.Comment{
		Content:      content,
		AuthorID:     AuthorID,
		AuthorName:   AuthorName,
		AuthorAvatar: AuthorAvatar,
		RootID:       rootID,
		ParentID:     ParentID,
		ParentName:   ParentName,
	}
	return DB.Model(model.Comment{}).
		Create(comment).Error
}

func (*Comment) DeleteCommentByID(id int64) error {
	return DB.Model(&model.Comment{}).
		Where("id = ?", id).
		Delete(nil).Error
}
