package util

import "fmt"

//go:generate mockery --name=Filters
type Filters interface {
	MakeIdFilters(id string) (filters []string)
	MakeAccessTokenFilter(access string) (filters []string)
	MakeClientIdString(clientID string) (filters string)
	MakeClientSecretString(clientSecret string) (filters string)
	MakeUserIDString(userID string) (filters []string)
}

type Filter struct{}

func NewFilters() (f *Filter) {
	return &Filter{}
}

func (f *Filter) MakeIdFilters(id string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", id),
	}
}

func (f *Filter) MakeAccessTokenFilter(access string) (filters []string) {
	return []string{
		fmt.Sprintf("accessToken:eq:%s", access),
	}
}

func (f *Filter) MakeClientIdString(clientID string) (filters string) {
	return fmt.Sprintf("clientId:eq:%s", clientID)
}

func (f *Filter) MakeClientSecretString(clientSecret string) (filters string) {
	return fmt.Sprintf("clientSecret:eq:%s", clientSecret)
}

func (f *Filter) MakeUserIDString(userID string) (filters []string) {
	return []string{
		fmt.Sprintf("userId:eq:%s", userID),
	}
}
