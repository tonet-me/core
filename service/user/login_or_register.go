package userservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
)

func (s Service) LoginOrRegister(ctx context.Context, req userparam.LoginOrRegisterRequest) (*userparam.LoginOrRegisterResponse, error) {

	userInfoFromToken, vErr := s.oAuthSvc.ValidationAndGetInfoFromToken(req.Token)
	if vErr != nil {
		return nil, vErr
	}

	isExisted, takenUser, gErr := s.repo.IsUserExistByEmail(ctx, userInfoFromToken.Email)
	if gErr != nil {
		return nil, gErr
	}
	var user entity.User
	if isExisted {
		user = takenUser
	} else {
		newUser, cErr := s.repo.CreateNewUser(ctx, entity.User{
			FirstName:       userInfoFromToken.FirstName,
			LastName:        userInfoFromToken.LastName,
			Email:           userInfoFromToken.Email,
			ProfilePhotoURL: userInfoFromToken.ProfilePhotoURL,
			Status:          entity.UserStatusActive,
		})
		if cErr != nil {
			return nil, cErr
		}
		user = newUser
	}

	authenticate := entity.Authenticable{
		ID: user.ID,
	}
	accessToken, caErr := s.authGenerator.CreateAccessToken(authenticate)
	if caErr != nil {
		return nil, caErr
	}

	refreshToken, crEre := s.authGenerator.CreateRefreshToken(authenticate)
	if crEre != nil {
		return nil, crEre
	}

	return &userparam.LoginOrRegisterResponse{
		User: user,
		Tokens: userparam.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

}
