package repo

const addNewChatQuery = `
	insert into chat (chat_id, owner_id, name, description, price, is_active)
	values
	($1, $2, $3, $4, $5, true)
`

const getChatsInfoByOwnerIdQuery = `
	select chat_id, name, description, price from chat where owner_id = $1
`

const disableChatQuery = `
	update chat set is_active = false where chat_id = $1
`

const changeDescriptionQuery = `
	update chat set description = $1 where chat_id = $2
`

const changePriceQuery = `
	update chat set price = $1 where chat_id = $2
`

const getAllSlavesQuery = `
	select user_id from users where chat_id = $1
`
