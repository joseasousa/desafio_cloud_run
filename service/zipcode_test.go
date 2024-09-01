package service

import (
	"errors"
	"testing"

	"github.com/joseasousa/desafio_cloud_run/internal/domain"
	"github.com/joseasousa/desafio_cloud_run/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ZipCodeServiceSuite struct {
	suite.Suite
	mockRepo *mock.ZipCodeRepository
	service  ZipCodeService
}

func (suite *ZipCodeServiceSuite) TestGetLocationByZipCode() {
	suite.T().Run("Success", func(t *testing.T) {
		expectedLocation := &domain.Location{
			Cep:        "123",
			Logradouro: "Rua Exemplo",
			Bairro:     "Bairro exemplo",
			Localidade: "Cidade Exemplo",
			Uf:         "UF",
		}

		expectedZipCode := "123"
		suite.mockRepo.On("GetLocationByZipCode", expectedZipCode).Return(expectedLocation, nil)
		result, err := suite.service.GetLocationByZipCode(expectedZipCode)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedLocation, result)
		suite.mockRepo.AssertExpectations(t)
	})

	suite.T().Run("Error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		expectedZipCode := "90210"
		suite.mockRepo.On("GetLocationByZipCode", expectedZipCode).Return(nil, expectedError)

		result, err := suite.service.GetLocationByZipCode(expectedZipCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, expectedError.Error())
		suite.mockRepo.AssertExpectations(t)
	})
}
func (suite *ZipCodeServiceSuite) TearDownSuite() {
	suite.mockRepo = nil
	suite.service = nil
}
func TestZipCodeServiceSuite(t *testing.T) {
	suite.Run(t, new(ZipCodeServiceSuite))
}
func (suite *ZipCodeServiceSuite) SetupSuite() {
	suite.mockRepo = mock.NewZipCodeRepository(suite.T())
	suite.service = NewZipCodeService(suite.mockRepo)
}
