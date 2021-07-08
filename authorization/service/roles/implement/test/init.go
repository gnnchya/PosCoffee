package test

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/suite"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/implement"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util/mocks"
	validatorMocks "github.com/gnnchya/PosCoffee/authorize/service/validator/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *validatorMocks.Validator
	repo      *mocks.Repository
	uuid      *mocks.UUID
	service   roles.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repo = &mocks.Repository{}
	suite.uuid = &mocks.UUID{}
	suite.service = implement.New(suite.validator, suite.repo, suite.uuid)
}

func (suite *PackageTestSuite) makeRoleIDFilters(roleID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", roleID),
		"deletedAt:isNull",
	}
}

func (suite *PackageTestSuite) makeTestPageOption() (opt *domain.PageOption) {
	return &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: nil,
	}
}

func (suite *PackageTestSuite) makeTestDomainRole() (roleDomain *domain.Roles) {
	now := carbon.Now()
	permTest := "test"
	return &domain.Roles{
		Name: &domain.Lang{
			Th: "test",
			En: "test",
		},
		Permissions: []*string{
			&permTest,
		},
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *rolesin.CreateInput) {
	permTest := "test"
	return &rolesin.CreateInput{
		Name: &rolesin.Lang{
			Th: "test",
			En: "test",
		},
		Permission: []*string{
			&permTest,
		},
	}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *rolesin.ReadInput) {
	return &rolesin.ReadInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *rolesin.UpdateInput) {
	permTest := "test"
	return &rolesin.UpdateInput{
		Name: &rolesin.Lang{
			Th: "test",
			En: "test",
		},
		Permission: []*string{
			&permTest,
		},
	}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *rolesin.DeleteInput) {
	return &rolesin.DeleteInput{
		ID: "test",
	}
}