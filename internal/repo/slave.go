package repo

const getAllSubscriptionsQuery = `
	select chat_id from users where user_id = $1
`

const isSubscribeExistsQuery = `
	select user_id from users where chat_id = $1 and user_id = $2
`

const newSubscribeQuery = `
	insert into users (chat_id, user_id, is_active, expired_date)
	values
	($1, $2, false, $3)
`

const payQuery = `
	update users set is_active = true where chat_id = $1 and user_id = $2
`

const isPaidQuery = `
	select is_active from users where chat_id = $1 and user_id = $2
`
