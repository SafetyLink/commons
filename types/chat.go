package types

import "time"

type Chat struct {
	ChatID         int64     `json:"chat_id"`
	UnreadMessages int64     `json:"unread_messages"`
	LastMessageAt  time.Time `json:"last_message_at"`
	Viewed         bool      `json:"viewed"`
	ViewedAt       time.Time `json:"viewed_at"`
	Users          *ChatUser `json:"user"`
}

type ChatUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	AvatarID int64  `json:"avatar_id"`
}
