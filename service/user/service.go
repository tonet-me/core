package userservice

import (
	"tonet-core/entity"
)

type Repository interface {
	IsUserExistByEmail(email string) (bool, entity.User, error)
	CreateNewUser(user entity.User) (entity.User, error)
	DeActiveUser(userID string) (bool, error)
	ActiveUser(userID string) (bool, error)
	GetUserByID(userID string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
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
