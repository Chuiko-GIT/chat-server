package chats

import (
	"github.com/Chuiko-GIT/chat-server/internal/repository"
	"github.com/Chuiko-GIT/chat-server/internal/service"
)

var _ service.Chats = &Serv{}

type Serv struct {
	chatRepo repository.Chats
}

func NewService(chatRepo repository.Chats) *Serv {
	return &Serv{
		chatRepo: chatRepo,
	}
}
