// +build integration

package test

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/repository/mongodb"
)

type SimpleTestStruct struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title"`
	List   []string           `bson:"list"`
	Text   *string            `bson:"text"`
	Number *int               `bson:"number"`
}

type PackageTestSuite struct {
	suite.Suite
	ctx  context.Context
	repo *mongodb.Repository
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	err := godotenv.Load()
	suite.NoError(err)
	conf := config.Get()

	suite.repo, err = mongodb.New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBCompanyTableName)
	suite.NoError(err)
}

func (suite *PackageTestSuite) SetupTest() {
	err := godotenv.Load()
	suite.NoError(err)
	conf := config.Get()

	suite.repo, err = mongodb.New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBCompanyTableName)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TearDownTest() {
	_, _ = suite.repo.Coll.DeleteMany(suite.ctx, bson.M{})
}

func (suite *PackageTestSuite) TearDownSuite() {
	_ = suite.repo.DB.Drop(suite.ctx)
}

func (suite *PackageTestSuite) makeTestStruct(title string, list ...string) (test *SimpleTestStruct) {
	return &SimpleTestStruct{
		Title: title,
		List:  list,
	}
}
