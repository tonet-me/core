package userservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
)

type Repository interface {
	IsUserExistByEmail(ctx context.Context, email string) (bool, entity.User, error)
	CreateNewUser(ctx context.Context, user entity.User) (entity.User, error)
	DeActiveUser(ctx context.Context, userID string) (bool, error)
	ActiveUser(ctx context.Context, userID string) (bool, error)
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	UpdateUser(ctx context.Context, userID string, user entity.User) (bool, error)
}

type AuthGenerator interface {
	CreateAccessToken(userAuthData entity.Authenticable) (string, error)
	CreateRefreshToken(userAuthData entity.Authenticable) (string, error)
}

type OAuthService interface {
	ValidationAndGetInfoFromToken(token string) (entity.OAuthUserInfo, error)
}

type Service struct {
	repo          Repository
	authGenerator AuthGenerator
	oAuthSvc      OAuthService
}
