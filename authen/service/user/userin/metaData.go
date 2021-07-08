package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/modern-go/reflect2"
)

type MetaDataStruct struct {
	Prefix            LanguageStruct `bson:"prefix" json:"prefix" validate:"required"`
	Name              LanguageStruct `bson:"name" json:"name" validate:"required"`
	Lastname          LanguageStruct `bson:"lastname" json:"lastname" validate:"required"`
	Email             string         `bson:"email" json:"email" validate:"required"`
	MobileNumber      string         `bson:"mobile_number" json:"mobile_number" validate:"required"`
	LineID            string         `bson:"line_id" json:"line_id"`
	Facebook          string         `bson:"facebook" json:"facebook"`
	BankAccount       string         `bson:"bank_account" json:"bank_account" validate:"required" `
	Gender            string         `bson:"gender" json:"gender" validate:"required"`
	BirthDate         int64          `bson:"birth_date" json:"birth_date" validate:"required"`
}

func (input *MetaDataStruct) ToDomain() (metadata *domain.MetaDataStruct) {
	if reflect2.IsNil(input) {
		return &domain.MetaDataStruct{}
	}

	return &domain.MetaDataStruct{
		Prefix: input.Prefix.ToDomain(),
		Name: input.Name.ToDomain(),
		Lastname: input.Lastname.ToDomain(),
		Email: input.Email,
		MobileNumber: input.MobileNumber,
		LineID: input.LineID,
		Facebook: input.Facebook,
		BankAccount: input.BankAccount,
		Gender: input.Gender,
		BirthDate: input.BirthDate,
	}
}
