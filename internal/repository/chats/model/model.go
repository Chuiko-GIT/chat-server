package model

import (
	"database/sql"
	"time"
)

type (
	Chat struct {
		ID        int64        `db:"id"`
		Usernames []string     `db:"usernames"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
)
