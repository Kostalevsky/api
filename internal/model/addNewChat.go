package model

type AddNewChat struct {
	ChatId      int    `json:"chat_id"`
	OwnerId     int    `json:"owner_id"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// {
//     "chat_id": 1312312312,
//     "owner_id": 12413413,
//     "price": 1312,
//     "name": "Порно 18+",
//     "description": "Новое порно каждый день"
// }
