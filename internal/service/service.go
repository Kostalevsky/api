package service

import (
	"context"
	"fmt"
	"project/internal/config"
	"project/internal/model"
	"project/internal/repo"
)

type Service interface {
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

type service struct {
	repo repo.Repo
}

func NewService(ctx context.Context, cfg config.Config) (Service, error) {
	r, err := repo.NewPgRepo(ctx, cfg.Repo)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pg repo")
	}

	return &service{
		repo: r,
	}, nil
}

func (s *service) AddNewChat(ctx context.Context, chat_id int, owner_id int, name string, desciption string, price int) error {
	if err := s.repo.AddNewChat(ctx, chat_id, owner_id, name, desciption, price); err != nil {
		return fmt.Errorf("failed to add new chat into repo. %w", err)
	}

	return nil
}

func (s *service) GetChatsInfoByOwnerId(ctx context.Context, owner_id int) ([]model.ChatInfo, error) {
	chats, err := s.repo.GetChatsInfoByOwnerId(ctx, owner_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get chats by owner id in repo. %w", err)
	}

	return chats, nil
}

func (s *service) DisableChat(ctx context.Context, chat_id int) error {
	if err := s.repo.DisableChat(ctx, chat_id); err != nil {
		return fmt.Errorf("failed to disable chat in repo. %w", err)
	}

	return nil
}

func (s *service) ChangeDescription(ctx context.Context, chat_id int, description string) error {
	if err := s.repo.ChangeDescription(ctx, chat_id, description); err != nil {
		return fmt.Errorf("failed to change description. %w", err)
	}

	return nil
}

func (s *service) ChangePrice(ctx context.Context, chat_id int, price int) error {
	if err := s.repo.ChangePrice(ctx, chat_id, price); err != nil {
		return fmt.Errorf("failed to change price. %w", err)
	}

	return nil
}

func (s *service) GetAllSlaves(ctx context.Context, chat_id int) ([]int, error) {
	slaves, err := s.repo.GetAllSlaves(ctx, chat_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get all slaves. %w", err)
	}

	return slaves, nil
}

func (s *service) NewSubscribe(ctx context.Context, chat_id int, user_id int) error {
	if err := s.repo.NewSubscribe(ctx, chat_id, user_id); err != nil {
		return fmt.Errorf("failed to make new subcribe in repo. %w", err)
	}

	return nil
}

func (s *service) GetAllSubsciptions(ctx context.Context, user_id int) ([]int, error) {
	subs, err := s.repo.GetAllSubsciptions(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get all subscriptions in repo. %w", err)
	}

	return subs, nil
}

func (s *service) Pay(ctx context.Context, chat_id int, user_id int) error {
	if err := s.repo.Pay(ctx, chat_id, user_id); err != nil {
		return fmt.Errorf("failed to pay in repo. %w", err)
	}

	return nil
}

func (s *service) IsSubscribeExists(ctx context.Context, chat_id int, users_id int) (bool, error) {
	ok, err := s.repo.IsSubscribeExists(ctx, chat_id, users_id)
	if err != nil {
		return false, fmt.Errorf("failed to check sub exist in repo. %w", err)
	}

	return ok, nil
}

func (s *service) IsPaid(ctx context.Context, chat_id int, user_id int) (bool, error) {
	ok, err := s.repo.IsPaid(ctx, chat_id, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to check paid status in repo. %w", err)
	}

	return ok, nil
}

func (s *service) Close() error {
	if err := s.repo.Close(); err != nil {
		return fmt.Errorf("failed to close repo. %w", err)
	}

	return nil
}
