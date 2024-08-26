package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/chat-server/internal/config"
	"github.com/Chuiko-GIT/chat-server/internal/config/env"
	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	chat_api.UnimplementedChatApiServer
	pool *pgxpool.Pool
}

func (s server) Create(ctx context.Context, req *chat_api.CreateRequest) (*chat_api.CreateResponse, error) {
	if len(req.Usernames) == 0 {
		return &chat_api.CreateResponse{}, errors.New("failed usernames is nil")
	}

	builder := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("usernames").
		Values(req.Usernames).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return &chat_api.CreateResponse{}, errors.New("failed to build query")
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return &chat_api.CreateResponse{}, errors.New("failed to create chat")
	}

	return &chat_api.CreateResponse{Id: chatID}, nil
}

func (s server) Delete(ctx context.Context, req *chat_api.DeleteRequest) (*emptypb.Empty, error) {
	builder := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where("id = $1", req.Id)

	query, args, err := builder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, errors.New("failed to build query")
	}

	if _, err = s.pool.Exec(ctx, query, args...); err != nil {
		return &emptypb.Empty{}, errors.New("failed to delete chat")
	}

	return &emptypb.Empty{}, nil
}

func (s server) SendMessage(ctx context.Context, req *chat_api.SendMessageRequest) (*emptypb.Empty, error) {
	builder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("message_from", "message_text").
		Values(req.From, req.Text).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, errors.New("failed to build query")
	}

	if _, err = s.pool.Query(ctx, query, args...); err != nil {
		return &emptypb.Empty{}, errors.New("failed to send message")
	}
	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	if err := config.Load(configPath); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_api.RegisterChatApiServer(s, server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
