package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/modern-go/reflect2"
)

type MetaDataStruct struct {
	Prefix            LanguageStruct `bson:"prefix" json:"prefix"`
	Name              LanguageStruct `bson:"name" json:"name"`
	Lastname          LanguageStruct `bson:"lastname" json:"lastname"`
	Email             string         `bson:"email" json:"email"`
	MobileNumber      string         `bson:"mobile_number" json:"mobile_number"`
	LineID            string         `bson:"line_id" json:"line_id"`
	Facebook          string         `bson:"facebook" json:"facebook"`
	BankAccount       string         `bson:"bank_account" json:"bank_account"`
	Gender            string         `bson:"gender" json:"gender"`
	BirthDate         int64          `bson:"birth_date" json:"birth_date"`
	RequiredGender    bool
	RequiredBirthDate bool
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
