package service

import "context"

type (
	Chats interface {
		Create(ctx context.Context, usernames []string) (int64, error)
		Delete(ctx context.Context, id int64) error
	}

	Messages interface {
		SendMessage(ctx context.Context, messageFrom, messageText string) error
	}
)
