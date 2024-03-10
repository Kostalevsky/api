package model

type ChangeDescription struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description"`
}
