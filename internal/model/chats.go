package model

import (
	"database/sql"
	"time"
)

type (
	Chat struct {
		ID        int64
		Usernames []string
		CreatedAt time.Time
		UpdatedAt sql.NullTime
	}
)
