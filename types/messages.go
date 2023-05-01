package types

import "time"

type Message struct {
	MessageID           int64      `json:"message_id"`
	Author              Author     `json:"author"`
	Content             string     `json:"content"`
	Attachment          Attachment `json:"attachment"`
	Edited              bool       `json:"edited"`
	Pinned              bool       `json:"pinned"`
	EditedAt            time.Time  `json:"edited_at"`
	ReferencedMessageID int64      `json:"referenced_message_id"`
	IsViewed            bool       `json:"is_viewed"`
	CreatedAt           time.Time  `json:"created_at"`
}

type Author struct {
	AuthorID int64  `json:"author_id"`
	Username string `json:"username"`
}

type Attachment struct {
	AttachmentID int64  `json:"attachment_id"`
	FileName     string `json:"file_name,omitempty"`
	IsScreenshot bool   `json:"is_screenshot"`
	Type         string `json:"type"`
	TTL          TTL    `json:"ttl"`
}

type TTL struct {
	TTL       bool      `json:"ttl"`
	AfterView bool      `json:"after_view"`
	TimeLimit time.Time `json:"time_limit"`
}
