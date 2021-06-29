package userin

import (
	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/authen/domain"
)

type MemberGroup struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func (input *MemberGroup) ToDomain() (memberGroup *domain.MemberGroup) {
	if reflect2.IsNil(input) {
		return &domain.MemberGroup{}
	}

	return &domain.MemberGroup{
		Name:    input.Name,
		Members: input.Members,
	}
}
