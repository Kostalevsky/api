package model

type ChangePrice struct {
	ChatId int `json:"chat_id"`
	Price  int `json:"price"`
}
