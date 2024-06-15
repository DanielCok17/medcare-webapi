package medcare

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type VaccinationSuite struct {
	suite.Suite
	dbServiceMock *DbServiceMock[VaccinationRecord]
}

func TestVaccinationSuite(t *testing.T) {
	suite.Run(t, new(VaccinationSuite))
}

type DbServiceMock[DocType interface{}] struct {
	mock.Mock
}

func (this *DbServiceMock[DocType]) CreateDocument(ctx context.Context, id string, document *DocType) error {
	args := this.Called(ctx, id, document)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) FindDocument(ctx context.Context, id string) (*DocType, error) {
	args := this.Called(ctx, id)
	return args.Get(0).(*DocType), args.Error(1)
}

func (this *DbServiceMock[DocType]) UpdateDocument(ctx context.Context, id string, document *DocType) error {
	args := this.Called(ctx, id, document)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) DeleteDocument(ctx context.Context, id string) error {
	args := this.Called(ctx, id)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) FindAllDocuments(ctx context.Context) ([]*DocType, error) {
	args := this.Called(ctx)
	return args.Get(0).([]*DocType), args.Error(1)
}

func (this *DbServiceMock[DocType]) Disconnect(ctx context.Context) error {
	args := this.Called(ctx)
	return args.Error(0)
}

func (suite *VaccinationSuite) SetupTest() {
	suite.dbServiceMock = &DbServiceMock[VaccinationRecord]{}

	// Compile time Assert that the mock is of type db_service.DbService[VaccinationRecord]
	var _ db_service.DbService[VaccinationRecord] = suite.dbServiceMock

	suite.dbServiceMock.
		On("FindDocument", mock.Anything, mock.Anything).
		Return(
			&VaccinationRecord{
				Id:        "test-vaccination",
				PatientId: "test-patient",
				Vaccine:   "test-vaccine",
				Date:      "2024-01-01T00:00:00Z",
			},
			nil,
		)
}

func (suite *VaccinationSuite) Test_UpdateVaccinationRecord_DbServiceUpdateCalled() {
	// ARRANGE
	suite.dbServiceMock.
		On("UpdateDocument", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	json := `{
		"id": "test-vaccination",
		"patientId": "test-patient",
		"vaccine": "test-vaccine",
		"date": "2024-01-01T00:00:00Z"
	}`

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Set("vaccination_record_db_service", suite.dbServiceMock)
	ctx.Params = []gin.Param{
		{Key: "recordId", Value: "test-vaccination"},
	}
	ctx.Request = httptest.NewRequest("PUT", "/vaccination_records/test-vaccination", strings.NewReader(json))

	sut := implVaccinationRecordsAPI{}

	// ACT
	sut.UpdateVaccinationRecord(ctx)

	// ASSERT
	suite.dbServiceMock.AssertCalled(suite.T(), "UpdateDocument", mock.Anything, "test-vaccination", mock.Anything)
}
