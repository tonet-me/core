package userservice

import (
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) GenerateTokens(authenticate entity.Authenticable) (*userparam.Tokens, error) {
	const op = richerror.OP("userservice.GenerateTokens")

	accessToken, caErr := s.authGenerator.CreateAccessToken(authenticate)
	if caErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(caErr))
	}

	refreshToken, crEre := s.authGenerator.CreateRefreshToken(authenticate)
	if crEre != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(crEre))
	}

	return &userparam.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
