package messages

import "context"

func (s Serv) SendMessage(ctx context.Context, messageFrom, messageText string) error {
	return s.messageRepo.SendMessage(ctx, messageFrom, messageText)
}
