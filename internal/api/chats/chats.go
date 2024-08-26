package chats

import (
	"github.com/Chuiko-GIT/chat-server/internal/service"
	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

type Implementation struct {
	chat_api.UnimplementedChatApiServer
	chatService service.Chats
}

func NewImplementation(chatService service.Chats) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
