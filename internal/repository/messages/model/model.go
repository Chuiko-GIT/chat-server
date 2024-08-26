package model

import (
	"database/sql"
	"time"
)

type (
	Message struct {
		ID          int64        `db:"id"`
		MessageFrom string       `db:"message_from"`
		MessageText string       `db:"message_text"`
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   sql.NullTime `db:"updated_at"`
	}
)
