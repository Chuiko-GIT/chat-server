package messages

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	db "github.com/Chuiko-GIT/chat-server/internal/client"
	"github.com/Chuiko-GIT/chat-server/internal/repository"
)

const (
	tableMessages = "messages"

	messageColumnID          = "id"
	messageColumnMessageFrom = "message_from"
	messageColumnMessageText = "message_text"
	messageColumnCreatedAt   = "created_at"
	messageColumnUpdatedAt   = "updated_at"
)

var _ repository.Messages = &Repo{}

type Repo struct {
	db db.Client
}

func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

func (r Repo) SendMessage(ctx context.Context, messageFrom, messageText string) error {
	builder := sq.Insert(tableMessages).
		PlaceholderFormat(sq.Dollar).
		Columns(messageColumnMessageFrom, messageColumnMessageText).
		Values(messageFrom, messageText).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "messages.repository.SendMessage",
		QueryRaw: query,
	}

	if _, err = r.db.DB().QueryContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}
