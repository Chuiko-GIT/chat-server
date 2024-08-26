package model

import (
	"database/sql"
	"time"
)

type (
	Message struct {
		ID          int64
		MessageFrom string
		MessageText string
		CreatedAt   time.Time
		UpdatedAt   sql.NullTime
	}
)
