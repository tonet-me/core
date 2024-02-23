package oauth

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

type Adapter struct {
	Google Google
}

func New(google Google) Adapter {
	return Adapter{
		Google: google,
	}
}

func (a Adapter) ValidationAndGetInfoFromToken(ctx context.Context, oAuthType entity.OAuthType, token string) (*entity.OAuthUserInfo, error) {
	const op = richerror.OP("ValidationAndGetInfoFromToken")

	switch oAuthType {
	case entity.OAuthTypeGoogle:
		return a.Google.ValidationToken(ctx, token)
	default:
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindBadRequest),
			richerror.WithMessage(errmsg.ErrorMsgTypeOfOAuthInvalid),
		)
	}
}
