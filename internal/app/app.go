package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/SaveliiYam/ProtoAPI_Auth/internal/app/grpc"
	"github.com/SaveliiYam/ProtoAPI_Auth/internal/services/auth"
	"github.com/SaveliiYam/ProtoAPI_Auth/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int, storagepath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagepath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcApp.New(log, authService, grpcPort)
	return &App{GRPCSrv: grpcApp}
}
