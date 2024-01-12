package dal

import (
	"Another-Nikki/dal/model"
	"context"
)

type Judge struct{}

func (*Judge) Create(ctx context.Context, uid int64, userName string, problemID int64, problemName, language, code string) error {
	judge := &model.Judge{
		UserID:      uid,
		UserName:    userName,
		ProblemID:   problemID,
		ProblemName: problemName,
		Language:    language,
		Code:        code,
	}
	err := DB.WithContext(ctx).Model(&model.Judge{}).Create(judge).Error
	return err
}
