package implement

import "github.com/gnnchya/PosCoffee/email/service/util"

func makeClientFilters(filter util.Filters, clientID string, clientSecret string) (filters []string) {
	return []string{
		filter.MakeClientIdString(clientID),
		filter.MakeClientSecretString(clientSecret),
	}
}

func makeConsumerFilters(filter util.Filters, clientID string, clientSecret string) (filters []string) {
	return []string{
		filter.MakeClientIdString(clientID),
		filter.MakeClientSecretString(clientSecret),
	}
}