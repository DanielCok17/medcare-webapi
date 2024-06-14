package medcare

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// DbServiceMock is a mock implementation of the DbService interface
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

func (this *DbServiceMock[DocType]) Disconnect(ctx context.Context) error {
	args := this.Called(ctx)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) FindAllDocuments(ctx context.Context) ([]*DocType, error) {
	args := this.Called(ctx)
	return args.Get(0).([]*DocType), args.Error(1)
}

// LabResultsSuite is a test suite for LabResultsAPI
type LabResultsSuite struct {
	suite.Suite
	dbServiceMock *DbServiceMock[LabResult]
}

func TestLabResultsSuite(t *testing.T) {
	suite.Run(t, new(LabResultsSuite))
}

func (suite *LabResultsSuite) SetupTest() {
	suite.dbServiceMock = &DbServiceMock[LabResult]{}

	// Compile time Assert that the mock is of type db_service.DbService[LabResult]
	var _ db_service.DbService[LabResult] = suite.dbServiceMock

	suite.dbServiceMock.
		On("FindDocument", mock.Anything, mock.Anything).
		Return(
			&LabResult{
				Id:     "labresult1",
				Result: "Normal",
			},
			nil,
		)
}

func (suite *LabResultsSuite) Test_UpdateLabResult_DbServiceUpdateCalled() {
	// ARRANGE
	suite.dbServiceMock.
		On("UpdateDocument", mock.Anything, "labresult1", mock.Anything).
		Return(nil)

	json := `{
        "patientId": "patient1",
        "testType": "Blood",
        "result": "Updated Result"
    }`

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Set("lab_result_db_service", suite.dbServiceMock)
	ctx.Params = []gin.Param{
		{Key: "recordId", Value: "labresult1"},
	}
	ctx.Request = httptest.NewRequest("PUT", "/api/lab_results/labresult1", strings.NewReader(json))

	sut := implLabResultsAPI{}

	// ACT
	sut.UpdateLabResult(ctx)

	// ASSERT
	suite.dbServiceMock.AssertCalled(suite.T(), "UpdateDocument", mock.Anything, "labresult1", mock.Anything)
	suite.Equal(http.StatusOK, recorder.Code)
}

func (suite *LabResultsSuite) Test_UpdateLabResult_InvalidRequest() {
	// ARRANGE
	json := `invalid json`

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Set("lab_result_db_service", suite.dbServiceMock)
	ctx.Params = []gin.Param{
		{Key: "recordId", Value: "labresult1"},
	}
	ctx.Request = httptest.NewRequest("PUT", "/api/lab_results/labresult1", strings.NewReader(json))

	sut := implLabResultsAPI{}

	// ACT
	sut.UpdateLabResult(ctx)

	// ASSERT
	suite.Equal(http.StatusBadRequest, recorder.Code)
}
