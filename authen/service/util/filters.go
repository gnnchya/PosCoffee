package util

import (
	"fmt"
)

type Filters interface {
	MakeIdFilters(id string) (filters []string)
	MakeIdFilterString(id string) (filters string)
	MakeFilterEmailString(email string) (filters string)
	MakeIdFiltersNotEqualString(id string) (filters string)
	MakeFilterMobileNumberString(email string) (filters string)
	MakeFilterUserName(username string) (filters []string)
	MakeUIDFilters(uid string) (filters []string)
	MakePasswordFilter(password string) (filter []string)
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

func (f *Filter) MakeUIDFilters(uid string) (filters []string) {
	return []string{
		fmt.Sprintf("uid:eq:%s", uid),
	}
}


func (f *Filter) MakeIdFilterString(id string) (filters string) {
	return fmt.Sprintf("_id:eq:%s", id)
}

func (f *Filter) MakeFilterEmailString(email string) (filters string) {
	return fmt.Sprintf("meta_data.email:eq:%s", email)
}

func (f *Filter) MakeIdFiltersNotEqualString(id string) (filters string) {
	return fmt.Sprintf("_id:ne:%s", id)
}

func (f *Filter) MakeFilterMobileNumberString(mobileNumber string) (filters string) {
	return fmt.Sprintf("meta_data.mobile_number:eq:%s", mobileNumber)
}

func (f *Filter) MakeFilterUserName(username string) (filters []string) {
	return []string{
		fmt.Sprintf("username:eq:%s", username),
	}
}

func (f *Filter) MakePasswordFilter(password string) (filter []string){
	return []string{
		fmt.Sprintf("password:eq:%s", password),
	}
}
