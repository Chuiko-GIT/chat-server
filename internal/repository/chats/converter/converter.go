package converter

import (
	"github.com/Chuiko-GIT/chat-server/internal/model"
	dbModel "github.com/Chuiko-GIT/chat-server/internal/repository/chats/model"
)

func ToChatFromRepo(message dbModel.Chat) model.Chat {
	return model.Chat{
		ID:        message.ID,
		Usernames: message.Usernames,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
