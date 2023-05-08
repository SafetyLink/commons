package types

import "time"

type Message struct {
	MessageID           int64       `json:"message_id"`
	ChatID              int64       `json:"chat_id"`
	AuthorID            int64       `json:"author_id"`
	Content             string      `json:"content"`
	Attachment          *Attachment `json:"attachment"`
	TTL                 *TTL        `json:"ttl"`
	IsViewed            bool        `json:"is_viewed"`
	IsScreenshot        bool        `json:"is_screenshot"`
	Edited              bool        `json:"edited"`
	EditedAt            time.Time   `json:"edited_at"`
	Pinned              bool        `json:"pinned"`
	PinnedAt            time.Time   `json:"pinned_at"`
	ReferencedMessageID int64       `json:"referenced_message_id"`

	CreatedAt time.Time `json:"created_at"`
}

type Attachment struct {
	AttachmentID int64          `json:"attachment_id"`
	FileName     string         `json:"file_name,omitempty"`
	IsScreenshot bool           `json:"is_screenshot"`
	Type         string         `json:"type"`
	TTL          *AttachmentTTL `json:"ttl"`
}

type AttachmentTTL struct {
	TTL       bool      `json:"ttl"`
	AfterView bool      `json:"after_view"`
	TimeLimit time.Time `json:"time_limit"`
}
type TTL struct {
	TTL       bool      `json:"ttl"`
	AfterView bool      `json:"after_view"`
	TimeLimit time.Time `json:"time_limit"`
}
