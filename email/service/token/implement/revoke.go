package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/email/domain"
	"strings"
)

func (impl *implementation) RevokeToken(ctx context.Context, input *string) (err error) {
	header := strings.ReplaceAll(*input, "Bearer ", "")
	filters := impl.filter.MakeAccessTokenFilter(header)
	token := &domain.TokenStruct{}
	err = impl.tokenRepository.Read(ctx, filters, token)
	if err != nil {
		return err
	}

	err = impl.tokenRepository.Delete(ctx, filters)
	if err != nil {
		return err
	}

	return err
}
