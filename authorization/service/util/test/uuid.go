package test

func (suite *PackageTestSuite) TestUUIDGenerate() {
	fstActualErr := suite.uuid.Generate()
	secActualErr := suite.uuid.Generate()
	suite.NotEqual(fstActualErr, secActualErr)
}
