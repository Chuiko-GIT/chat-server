package chats

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

func (i *Implementation) Delete(ctx context.Context, req *chat_api.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.chatService.Delete(ctx, req.Id)
}
