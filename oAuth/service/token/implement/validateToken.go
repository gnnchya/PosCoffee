package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/out"
	"strings"
	"time"
)

func (impl *implementation) ValidateToken(ctx context.Context, input *string) (view *out.ValidateTokenView, err error) {
	header := strings.ReplaceAll(*input, "Bearer ", "")
	token := &domain.TokenStruct{}
	filters := impl.filter.MakeAccessTokenFilter(header)
	err = impl.tokenRepository.Read(ctx, filters, token)
	if err != nil {
		return nil, err
	}
	fmt.Println("validate token", token)
	refresh := time.Unix(token.CreatedAt, 0)
	refreshExpiresIn := time.Hour * 24 * 3
	accessCreateAt := time.Unix(token.CreatedAt, 0)
	accessExpiresIn := time.Minute * 2
	timeNow := time.Now()
	if token.AccessToken != header {
		return nil, fmt.Errorf("error : Unauthorized: Invalid Access Token") //Unauthorized(errors.ErrInvalidAccessToken)
	} else if token.RefreshToken != "" && token.RefreshExpire != 0 &&
		refresh.Add(refreshExpiresIn).Before(timeNow) {
		return nil, fmt.Errorf("error : Unauthorized: Expired Refresh Token") //util.Unauthorized(errors.ErrExpiredRefreshToken)
	} else if token.CreatedAt != 0 && accessCreateAt.Add(accessExpiresIn).Before(timeNow) {
		return nil, fmt.Errorf("error : Unauthorized: Expired Access Token") //util.Unauthorized(errors.ErrExpiredAccessToken)
	}

	token.CreatedAt = int64(accessCreateAt.Add(accessExpiresIn).Sub(timeNow).Seconds())
	fmt.Println("output token", out.ValidateTokenToView(token))
	return out.ValidateTokenToView(token), err
}
