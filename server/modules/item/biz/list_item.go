package biz

import (
	"context"
	"server/common"
	"server/modules/item/model"
)

type ListItemStorage interface {
	FindItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}

type ListItemBiz struct {
	storage ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *ListItemBiz {
	return &ListItemBiz{storage: store}
}

func (biz *ListItemBiz) ListItemBy(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.storage.FindItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
