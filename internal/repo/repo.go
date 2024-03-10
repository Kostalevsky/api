package repo

import (
	"context"
	"project/internal/model"
)

type Repo interface {
	AddNewChat(context.Context, int, int, string, string, int) error
	GetChatsInfoByOwnerId(context.Context, int) ([]model.ChatInfo, error)
	DisableChat(context.Context, int) error
	ChangeDescription(context.Context, int, string) error
	ChangePrice(context.Context, int, int) error
	GetAllSlaves(context.Context, int) ([]int, error)

	NewSubscribe(context.Context, int, int) error
	GetAllSubsciptions(context.Context, int) ([]int, error)
	Pay(context.Context, int, int) error
	IsSubscribeExists(context.Context, int, int) (bool, error)
	IsPaid(context.Context, int, int) (bool, error)

	Close() error
}
