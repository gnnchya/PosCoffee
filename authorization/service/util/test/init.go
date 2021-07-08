package test

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/service/validator"

	"github.com/stretchr/testify/suite"

	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	uuid      util.UUID
	validator validator.Validator
	vRepo     *validator.ValidatorRepository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	var err error
	suite.uuid, err = util.NewUUID()
	suite.NoError(err)
	suite.vRepo = &validator.ValidatorRepository{}
	suite.validator = validator.New(suite.vRepo)
}
