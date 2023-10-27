package dal

import (
	"Another-Nikki/dal/model"
)

type Problem struct {
}

func (*Problem) Create(name, content string) error {
	problem := &model.Problem{
		Name:    name,
		Content: content,
	}
	return DB.Model(model.Problem{}).Create(problem).Error
}

func (*Problem) GetInfoById(id int) (*model.Problem, error) {
	var problem model.Problem
	err := DB.Model(model.Problem{}).
		Where("id = ?", id).
		First(&problem).
		Error
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (*Problem) GetList(page int) ([]model.Problem, error) {
	res := make([]model.Problem, 0)
	err := DB.Model(model.Problem{}).
		Limit(20).
		Offset((page - 1) * 20).
		Order("ID DESC").
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*Problem) UpdateProblem(id int, name, content string) error {
	err := DB.Model(model.Problem{}).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{
			"name":    name,
			"content": content,
		}).
		Error
	return err
}
