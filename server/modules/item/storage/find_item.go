package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/modules/item/model"
)

func (sql *SqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	err := sql.db.Where(cond).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
