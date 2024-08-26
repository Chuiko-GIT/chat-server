package chats

import "context"

func (s Serv) Delete(ctx context.Context, id int64) error {
	return s.chatRepo.Delete(ctx, id)
}
