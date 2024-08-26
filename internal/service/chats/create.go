package chats

import "context"

func (s Serv) Create(ctx context.Context, usernames []string) (int64, error) {
	return s.chatRepo.Create(ctx, usernames)
}
