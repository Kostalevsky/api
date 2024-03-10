package transport

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"project/internal/config"
	"project/internal/service"
)

type Transport interface {
	Run() error
	Close() error
}

type transport struct {
	router  *http.Server
	service service.Service
}

func NewTransport(ctx context.Context, cfg config.Config) (Transport, error) {
	s, err := service.NewService(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create new service. %w", err)
	}

	router := &http.Server{
		Addr: net.JoinHostPort(cfg.Transport.Host, cfg.Transport.Port),
	}

	tr := transport{
		router:  router,
		service: s,
	}

	tr.setupRoutes()

	return &tr, nil
}

func (t *transport) setupRoutes() {
	mx := http.NewServeMux()

	mx.HandleFunc("/add_new_chat", t.addNewChat)
	mx.HandleFunc("/get_chats", t.getChatsInfoByOwnerId)
	mx.HandleFunc("/disable_chat", t.disableChat)
	mx.HandleFunc("/change_description", t.changeDescription)
	mx.HandleFunc("/change_price", t.changePrice)
	mx.HandleFunc("/get_all_slaves", t.getAllSlaves)

	mx.HandleFunc("/new_subscribe", t.newSubscribe)
	mx.HandleFunc("/get_all_subscriptions", t.getAllSubsciptions)
	mx.HandleFunc("/pay", t.pay)
	mx.HandleFunc("/is_subscribe_exist", t.isSubscribeExists)
	mx.HandleFunc("/is_paid", t.isPaid)

	t.router.Handler = mx
}

func (t *transport) Run() error {
	if err := t.router.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to listen and serve. %w", err)
	}

	return nil
}

func (t *transport) Close() error {
	return nil
}
