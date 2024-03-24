package biz

import (
	"context"
	"server/common"
	"server/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type GetItemBiz struct {
	storage GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *GetItemBiz {
	return &GetItemBiz{storage: store}
}

func (biz *GetItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrorCannotGetEntity(model.ENTITY_NAME, err)
	}

	return data, nil
}
