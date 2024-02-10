package userservice

import (
	"tonet-core/entity"
	"tonet-core/param/user"
)

func (s Service) LoginOrRegister(req userparam.LoginOrRegisterRequest) (*userparam.LoginOrRegisterResponse, error) {

	userInfoFromToken, vErr := s.oAuthSvc.ValidationAndGetInfoFromToken(req.Token)
	if vErr != nil {
		return nil, vErr
	}

	isExisted, takenUser, gErr := s.repo.IsUserExistByEmail(userInfoFromToken.Email)
	if gErr != nil {
		return nil, gErr
	}
	var user entity.User
	if isExisted {
		user = takenUser
	} else {
		newUser, cErr := s.repo.CreateNewUser(entity.User{
			FirstName:       userInfoFromToken.FirstName,
			LastName:        userInfoFromToken.LastName,
			Email:           userInfoFromToken.Email,
			ProfilePhotoURL: userInfoFromToken.ProfilePhotoURL,
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
