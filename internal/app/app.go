package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/SaveliiYam/ProtoAPI_Auth/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int, storagepath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище
	//TODO: init auth service
	grpcApp := grpcApp.New(log, grpcPort)
	return &App{GRPCSrv: grpcApp}
}
