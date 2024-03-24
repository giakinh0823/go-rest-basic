package storage

import (
	"context"
	"server/modules/item/model"
)

func (sql *SqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	err := sql.db.Where(cond).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
