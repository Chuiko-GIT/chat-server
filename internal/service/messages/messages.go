package messages

import (
	"github.com/Chuiko-GIT/chat-server/internal/repository"
	"github.com/Chuiko-GIT/chat-server/internal/service"
)

var _ service.Messages = &Serv{}

type Serv struct {
	messageRepo repository.Messages
}

func NewService(messageRepo repository.Messages) *Serv {
	return &Serv{
		messageRepo: messageRepo,
	}
}
