package oauth

import (
	"encoding/json"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func mapToGoogleClaim(data []byte) (*entity.OAuthUserInfo, error) {
	const op = richerror.OP("oauth.unmarshalFromOAuthDataToUserInfoEntity")

	googleClaim := new(GoogleClaim)
	if jErr := json.Unmarshal(data, googleClaim); jErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(jErr),
		)
	}

	return &entity.OAuthUserInfo{
		FirstName: googleClaim.GivenName,
		LastName:  googleClaim.FamilyName,
		Email:     googleClaim.Email,
	}, nil
}
