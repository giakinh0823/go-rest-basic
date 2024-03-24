package storage

import (
	"context"
	"server/common"
	"server/modules/item/model"
)

func (sql *SqlStore) FindItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error) {
	paging.Process()

	var result []model.TodoItem

	db := sql.db.Where("status <> ?", "DELETED")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Table(model.TodoItem{}.TableName()).
		Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
