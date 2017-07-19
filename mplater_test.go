package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) TestParseKeyValue() {
	result, err := parseKeyValues([]string{"key=value"})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result, map[string]string{"key": "value"})
}

func (suite *MainTestSuite) TestFailParseKeyValueNoValue() {
	_, err := parseKeyValues([]string{"key"})

	assert.Equal(suite.T(), err.Error(), "Invalid input value: key")
}

func (suite *MainTestSuite) TestPrimeTemplateEngine() {
	output, err := primeTemplateEngine("./test_data/simple_file.txt")

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
}

func TestRunSuite(t *testing.T) {
	testRunner := new(MainTestSuite)
	suite.Run(t, testRunner)
}
