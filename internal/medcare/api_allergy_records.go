package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AllergyRecordsAPI interface {
	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// DeleteAllergyRecord - Delete an allergy record
	DeleteAllergyRecord(ctx *gin.Context)

	// GetAllAllergyRecords - Get all allergy records
	GetAllAllergyRecords(ctx *gin.Context)

	// GetAllergyRecordById - Get an allergy record by ID
	GetAllergyRecordById(ctx *gin.Context)
}

// partial implementation of AllergyRecordsAPI - all functions must be implemented in add on files
type implAllergyRecordsAPI struct {
}

func newAllergyRecordsAPI() AllergyRecordsAPI {
	return &implAllergyRecordsAPI{}
}

func (this *implAllergyRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodDelete, "/allergy-records/:recordId", this.DeleteAllergyRecord)
	routerGroup.Handle(http.MethodGet, "/allergy-records", this.GetAllAllergyRecords)
	routerGroup.Handle(http.MethodGet, "/allergy-records/:recordId", this.GetAllergyRecordById)
}

// DeleteAllergyRecord - Delete an allergy record
func (this *implAllergyRecordsAPI) DeleteAllergyRecord(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetAllAllergyRecords - Get all allergy records
func (this *implAllergyRecordsAPI) GetAllAllergyRecords(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetAllergyRecordById - Get an allergy record by ID
func (this *implAllergyRecordsAPI) GetAllergyRecordById(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
