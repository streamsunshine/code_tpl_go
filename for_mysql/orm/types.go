package orm

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID                int            `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index:idx_chat_user_deleted,priority:2"`
	UserID            string         `json:"user_id" gorm:"index:idx_chat_user_deleted,priority:1"`
	Name              string         `json:"name"`
	Count             int            `json:"count"` // Record 条数
	MaxLengthExceeded bool           `json:"max_length_exceeded"`
	Type              int            `json:"type"`
	Status            int            `json:"status"`
}

func (Chat) TableName() string {
	return "chat"
}

type Record struct {
	ID           int            `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index:idx_record_chat_deleted,priority:2"`
	Duration     float64        `json:"duration"` // 处理时间，单位 s
	ChatID       int            `json:"chat_id" gorm:"index:idx_record_chat_deleted,priority:1"`
	Request      string         `json:"request"`
	Response     string         `json:"response"`
	Context      string         `json:"context"`
	SeqLen       int            `json:"seq_len"`
	GroupID      string         `json:"group_id"`
	SessionID    string         `json:"session_id"`
	ModelID      int            `json:"model_id"`
	Feedback     string         `json:"feedback"`
	ComplainType string         `json:"complain_type"`
	ContextLen   int            `json:"context_len"`
	Status       int            `json:"status"` // 0 正常， 1 被打断
}

func (Record) TableName() string {
	return "record"
}
