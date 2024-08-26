package messages

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

func (i *Implementation) SendMessage(ctx context.Context, req *chat_api.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.messageService.SendMessage(ctx, req.From, req.Text)
}
