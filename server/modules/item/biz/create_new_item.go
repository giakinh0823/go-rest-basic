package biz

import (
	"context"
	"server/common"
	"server/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type CreateItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *CreateItemBiz {
	return &CreateItemBiz{storage: store}
}

func (biz *CreateItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreate) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(model.ENTITY_NAME, err)
	}

	return nil
}
