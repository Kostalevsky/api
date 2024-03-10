package repo

import (
	"context"
	"fmt"
	"net"
	"project/internal/config"
	"project/internal/model"
	"sync"
	"time"

	"project/internal/logger"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool     *pgxpool.Pool
	poolOnce sync.Once
)

func dbSetup(ctx context.Context, cfg config.Repo) (*pgxpool.Pool, error) {
	var err error

	poolOnce.Do(func() {
		dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.User, cfg.Pass, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Database, cfg.SSLMode)

		var poolConfig *pgxpool.Config

		poolConfig, err = pgxpool.ParseConfig(dsn)
		if err != nil {
			logger.GetLogger().Debug().Err(err).Msg("failed to parse db config")

			return
		}

		poolConfig.ConnConfig.Tracer = otelpgx.NewTracer()

		pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			logger.GetLogger().Debug().Err(err).Msg("failed to connect to db")

			return
		}
	})

	if err != nil {
		return nil, fmt.Errorf("failed to setup db. %w", err)
	}

	return pool, nil
}

type pg struct {
	pool *pgxpool.Pool
}

func NewPgRepo(ctx context.Context, cfg config.Repo) (Repo, error) {
	pool, err := dbSetup(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup db. %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping. %w", err)
	}

	result := &pg{
		pool: pool,
	}

	return result, nil
}

func (p *pg) Close() error {
	p.pool.Close()

	return nil
}

func (p *pg) AddNewChat(ctx context.Context, chat_id int, owner_id int, name string, description string, price int) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction. %w", err)
	}

	if _, err := tx.Exec(ctx, addNewChatQuery, chat_id, owner_id, name, description, price); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit query. %w", err)
	}

	return nil
}

func (p *pg) GetChatsInfoByOwnerId(ctx context.Context, owner_id int) ([]model.ChatInfo, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction. %w", err)
	}

	rows, err := tx.Query(ctx, getChatsInfoByOwnerIdQuery, owner_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query. %w", err)
	}

	info := make([]model.ChatInfo, 0)

	for rows.Next() {
		var (
			chat_id     int
			name        string
			description string
			price       int
		)

		if err := rows.Scan(&chat_id, &name, &description, &price); err != nil {
			return nil, fmt.Errorf("failed to scan rows. %w", err)
		}

		info = append(info, model.ChatInfo{
			ChatId:      chat_id,
			Name:        name,
			Description: description,
			Price:       price,
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction. %w", err)
	}

	return info, nil
}

func (p *pg) DisableChat(ctx context.Context, chat_id int) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction. %w", err)
	}

	if _, err := tx.Exec(ctx, disableChatQuery, chat_id); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction. %w", err)
	}

	return nil
}

func (p *pg) ChangeDescription(ctx context.Context, chat_id int, description string) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction. %w", err)
	}

	if _, err := tx.Exec(ctx, changeDescriptionQuery, description, chat_id); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction. %w", err)
	}

	return nil
}

func (p *pg) ChangePrice(ctx context.Context, chat_id int, price int) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction. %w", err)
	}

	if _, err := tx.Exec(ctx, changePriceQuery, price, chat_id); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction. %w", err)
	}

	return nil
}

func (p *pg) GetAllSlaves(ctx context.Context, chat_id int) ([]int, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction. %w", err)
	}

	rows, err := tx.Query(ctx, getAllSlavesQuery, chat_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query. %w", err)
	}

	slaves := make([]int, 0)

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan rows. %w", err)
		}

		slaves = append(slaves, id)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction. %w", err)
	}

	return slaves, nil
}

func (p *pg) NewSubscribe(ctx context.Context, chat_id int, user_id int) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction %w", err)
	}

	if _, err := tx.Exec(ctx, newSubscribeQuery, chat_id, user_id, time.Now().AddDate(0, 1, 0)); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction. %w", err)
	}

	return nil
}

func (p *pg) GetAllSubsciptions(ctx context.Context, user_id int) ([]int, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction. %w", err)
	}

	rows, err := tx.Query(ctx, getAllSubscriptionsQuery, user_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query. %w", err)
	}

	subs := make([]int, 0)

	for rows.Next() {
		var sub int

		if err := rows.Scan(&sub); err != nil {
			return nil, fmt.Errorf("failed to scan rows. %w", err)
		}

		subs = append(subs, sub)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction. %w", err)
	}

	return subs, nil
}

func (p *pg) Pay(ctx context.Context, chat_id int, user_id int) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin new transaction. %w", err)
	}

	if _, err := tx.Exec(ctx, payQuery, chat_id, user_id); err != nil {
		return fmt.Errorf("failed to execute query. %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction. %w", err)
	}

	return nil
}

func (p *pg) IsSubscribeExists(ctx context.Context, chat_id int, user_id int) (bool, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction. %w", err)
	}

	rows, err := tx.Query(ctx, isSubscribeExistsQuery, chat_id, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to execute transaction. %w", err)
	}

	res := false

	for rows.Next() {
		var user int

		if err := rows.Scan(&user); err != nil {
			return false, fmt.Errorf("failed to scan rows. %w", err)
		}

		res = user == user_id
	}

	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("failed to commit transaction. %w", err)
	}

	return res, nil
}

func (p *pg) IsPaid(ctx context.Context, chat_id int, user_id int) (bool, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction. %w", err)
	}

	rows, err := tx.Query(ctx, isPaidQuery, chat_id, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to execute query. %w", err)
	}

	var res bool

	for rows.Next() {
		r := false

		if err := rows.Scan(&r); err != nil {
			return false, fmt.Errorf("failed to scan rows. %w", err)
		}

		res = r
	}

	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("failed to commit transaction. %w", err)
	}

	return res, nil
}
