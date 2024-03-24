package biz

import (
	"context"
	"server/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type DeleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *DeleteItemBiz {
	return &DeleteItemBiz{storage: store}
}

func (biz *DeleteItemBiz) DeleteItemBy(ctx context.Context, id int) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data == nil || *data.Status == model.DELETED {
		return model.ErrItemNotFound
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
