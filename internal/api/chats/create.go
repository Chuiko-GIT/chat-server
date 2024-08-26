package chats

import (
	"context"
	"errors"

	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

func (i *Implementation) Create(ctx context.Context, req *chat_api.CreateRequest) (*chat_api.CreateResponse, error) {
	if len(req.Usernames) == 0 {
		return &chat_api.CreateResponse{}, errors.New("failed usernames is nil")
	}

	id, err := i.chatService.Create(ctx, req.Usernames)
	if err != nil {
		return &chat_api.CreateResponse{}, errors.New("failed to create chat")
	}

	return &chat_api.CreateResponse{Id: id}, nil
}
