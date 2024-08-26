package converter

import (
	"github.com/Chuiko-GIT/chat-server/internal/model"
	dbModel "github.com/Chuiko-GIT/chat-server/internal/repository/messages/model"
)

func ToMessageFromRepo(message dbModel.Message) model.Message {
	return model.Message{
		ID:          message.ID,
		MessageFrom: message.MessageFrom,
		MessageText: message.MessageText,
		CreatedAt:   message.CreatedAt,
		UpdatedAt:   message.UpdatedAt,
	}
}
