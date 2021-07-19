package test

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/service/util/mocks"
	"github.com/gnnchya/PosCoffee/authorize/service/validator"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	repo      *mocks.Repository
	validator validator.Validator
	vRepo     *validator.ValidatorRepository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = &mocks.Repository{}
	suite.vRepo = &validator.ValidatorRepository{}
	suite.validator = validator.New(suite.vRepo)
}

type SimpleTestStruct struct {
	Title string `validate:"required"`
	Body  string `validate:"max=5"`
}

func (suite *PackageTestSuite) TestValidatorValidateValid() {
	given := &SimpleTestStruct{
		Title: "test",
		Body:  "AAA",
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestValidatorValidateInvalid() {
	given := &SimpleTestStruct{
		Title: "",
		Body:  "AAAAAAA",
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}
