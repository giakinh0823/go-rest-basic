package biz

import (
	"context"
	"server/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

type UpdateItemBiz struct {
	storage UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *UpdateItemBiz {
	return &UpdateItemBiz{storage: store}
}

func (biz *UpdateItemBiz) UpdateItemBy(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data == nil || *data.Status == model.DELETED {
		return model.ErrItemNotFound
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
