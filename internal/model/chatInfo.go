package model

type ChatInfo struct {
	ChatId      int    `json:"chat_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
