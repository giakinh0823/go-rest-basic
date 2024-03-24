package storage

import (
	"context"
	"server/common"
	"server/modules/item/model"
)

func (sql *SqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := sql.db.Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
