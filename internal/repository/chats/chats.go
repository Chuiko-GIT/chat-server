package chats

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	db "github.com/Chuiko-GIT/chat-server/internal/client"
	"github.com/Chuiko-GIT/chat-server/internal/repository"
)

const (
	tableChats = "chats"

	chatColumnID        = "id"
	chatColumnUsernames = "usernames"
	chatColumnCreatedAt = "created_at"
	chatColumnUpdatedAt = "updated_at"
)

var _ repository.Chats = &Repo{}

type Repo struct {
	db db.Client
}

func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

func (r Repo) Create(ctx context.Context, usernames []string) (int64, error) {
	builder := sq.Insert(tableChats).
		PlaceholderFormat(sq.Dollar).
		Columns(chatColumnUsernames).
		Values(usernames).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chats.repository.Create",
		QueryRaw: query,
	}

	var chatID int64
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID); err != nil {
		return 0, err
	}

	return chatID, nil
}

func (r Repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableChats).
		PlaceholderFormat(sq.Dollar).
		Where("id = $1", id)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chats.repository.Delete",
		QueryRaw: query,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}
