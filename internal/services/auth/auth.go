package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/SaveliiYam/ProtoAPI_Auth/internal/domain/models"
)

type Auth struct {
	log          *slog.Logger
	userProvider UserProvider
	userSaver    UserSaver
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appId int) (models.App, error)
}

// New returns a new instance of the Auth service
func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		log:          log,
		tokenTTL:     tokenTTL,
	}
}
