package userservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"strings"
	"time"
)

func (s Service) LoginOrRegister(ctx context.Context, req userparam.LoginOrRegisterRequest) (*userparam.LoginOrRegisterResponse, error) {
	const op = richerror.OP("userservice.LoginOrRegister")
	userInfoFromToken, vErr := s.oAuthSvc.ValidationAndGetInfoFromToken(ctx, req.ProviderID, req.Token)
	if vErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(vErr))
	}

	isExisted, takenUser, gErr := s.repo.IsUserExistByEmail(ctx, userInfoFromToken.Email)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	var user entity.User
	var isNewUser bool
	if isExisted {
		user = takenUser
	} else {
		newUser, cErr := s.repo.CreateNewUser(ctx, entity.User{
			FirstName:       userInfoFromToken.FirstName,
			LastName:        userInfoFromToken.LastName,
			Email:           strings.ToLower(userInfoFromToken.Email),
			ProfilePhotoURL: userInfoFromToken.ProfilePhotoURL,
			Status:          entity.UserStatusActive,
			EmailVerified:   true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		})
		if cErr != nil {
			return nil, richerror.New(richerror.WithOp(op),
				richerror.WithInnerError(cErr))
		}
		user = newUser
		isNewUser = true
	}

	authenticate := entity.Authenticable{
		ID: user.ID,
	}

	tokens, tErr := s.GenerateTokens(authenticate)
	if tErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(tErr))
	}

	return &userparam.LoginOrRegisterResponse{
		User:    user,
		Tokens:  *tokens,
		NewUser: isNewUser,
	}, nil

}
