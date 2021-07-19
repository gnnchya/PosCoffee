package test

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/suite"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/implement"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/util/mocks"
	validatorMocks "github.com/gnnchya/PosCoffee/authorize/service/validator/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *validatorMocks.Validator
	repo      *mocks.Repository
	uuid      *mocks.UUID
	service   permissions.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repo = &mocks.Repository{}
	suite.uuid = &mocks.UUID{}
	suite.service = implement.New(suite.validator, suite.repo, suite.uuid)
}

func (suite *PackageTestSuite) makeTestPermission() (permission *domain.Permissions) {
	now := carbon.Now().Unix()
	return &domain.Permissions{
		ID:        "test",
		Method:    "test",
		Endpoint:  "test",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *permissionsin.CreateInput) {
	return &permissionsin.CreateInput{
		Method:   "test",
		Endpoint: "test",
	}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *permissionsin.ReadInput) {
	return &permissionsin.ReadInput{}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *permissionsin.UpdateInput) {
	return &permissionsin.UpdateInput{}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *permissionsin.DeleteInput) {
	return &permissionsin.DeleteInput{}
}

func (suite *PackageTestSuite) makePermissionIDFilters(permissionID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", permissionID),
		"deletedAt:isNull",
	}
}
