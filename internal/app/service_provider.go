package app

import (
	"context"
	"log"

	uImpl "github.com/Chuiko-GIT/chat-server/internal/api/chats"
	"github.com/Chuiko-GIT/chat-server/internal/api/messages"
	mImpl "github.com/Chuiko-GIT/chat-server/internal/api/messages"
	db "github.com/Chuiko-GIT/chat-server/internal/client"
	"github.com/Chuiko-GIT/chat-server/internal/client/db/pg"
	"github.com/Chuiko-GIT/chat-server/internal/closer"
	"github.com/Chuiko-GIT/chat-server/internal/config"
	"github.com/Chuiko-GIT/chat-server/internal/config/env"
	"github.com/Chuiko-GIT/chat-server/internal/repository"
	uRepo "github.com/Chuiko-GIT/chat-server/internal/repository/chats"
	mRepo "github.com/Chuiko-GIT/chat-server/internal/repository/messages"
	"github.com/Chuiko-GIT/chat-server/internal/service"
	uServise "github.com/Chuiko-GIT/chat-server/internal/service/chats"
	mServise "github.com/Chuiko-GIT/chat-server/internal/service/messages"

	"github.com/Chuiko-GIT/chat-server/internal/api/chats"
)

type ServiceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	dbClient   db.Client

	chatRepository    repository.Chats
	messageRepository repository.Messages

	chatService    service.Chats
	messageService service.Messages

	chatImpl    *chats.Implementation
	messageImpl *messages.Implementation
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *ServiceProvider) ChatRepository(ctx context.Context) repository.Chats {
	if s.chatRepository == nil {
		s.chatRepository = uRepo.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *ServiceProvider) ChatService(ctx context.Context) service.Chats {
	if s.chatService == nil {
		s.chatService = uServise.NewService(s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *ServiceProvider) ChatImpl(ctx context.Context) *chats.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = uImpl.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *ServiceProvider) MessageRepository(ctx context.Context) repository.Messages {
	if s.messageRepository == nil {
		s.messageRepository = mRepo.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *ServiceProvider) MessageService(ctx context.Context) service.Messages {
	if s.messageService == nil {
		s.messageService = mServise.NewService(s.MessageRepository(ctx))
	}

	return s.messageService
}

func (s *ServiceProvider) MessageImpl(ctx context.Context) *messages.Implementation {
	if s.messageImpl == nil {
		s.messageImpl = mImpl.NewImplementation(s.MessageService(ctx))
	}

	return s.messageImpl
}
