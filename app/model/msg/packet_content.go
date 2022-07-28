package msg

type AtContent struct {
	Content string    `json:"Content"`
	UserExt []UserExt `json:"UserExt"`
	UserID  []int64   `json:"UserID"`
}

type UserExt struct {
	QQNick string `json:"QQNick"`
	QQUid  int64  `json:"QQUid"`
}
