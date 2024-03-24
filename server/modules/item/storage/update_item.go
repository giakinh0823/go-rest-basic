package storage

import (
	"context"
	"server/modules/item/model"
)

func (sql *SqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
	if err := sql.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
