package model

func GetMessage(userId int64) []Message {
	sql := "select * from messages where status=0 and user_id=?"
	var messages []Message
	Conn.Raw(sql, userId).Scan(&messages)
	return messages
}
