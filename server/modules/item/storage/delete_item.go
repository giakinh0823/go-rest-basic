package storage

import (
	"context"
	"server/modules/item/model"
)

func (sql *SqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deleted := model.DELETED

	if err := sql.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deleted.String(),
	}).Error; err != nil {
		return err
	}

	return nil
}
