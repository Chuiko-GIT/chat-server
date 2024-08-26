package messages

import (
	"github.com/Chuiko-GIT/chat-server/internal/service"
	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

type Implementation struct {
	chat_api.UnimplementedChatApiServer
	messageService service.Messages
}

func NewImplementation(messageService service.Messages) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
