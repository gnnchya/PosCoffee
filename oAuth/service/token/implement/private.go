package implement

import "github.com/gnnchya/PosCoffee/oAuth/service/util"

func makeClientFilters(filter util.Filters, clientID string, clientSecret string) (filters []string) {
	return []string{
		filter.MakeClientIdString(clientID),
		filter.MakeClientSecretString(clientSecret),
	}
}

func makeConsumerFilters(filter util.Filters, clientID, clientSecret, redirectUri string) (filters []string) {
	return []string{
		filter.MakeClientIdString(clientID),
		filter.MakeClientSecretString(clientSecret),
		filter.MakeRedirectUriString(redirectUri),
	}
}