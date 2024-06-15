package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AllergyRecordsAPI interface {
	// internal registration of API routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateAllergyRecord - Saves new allergy record
	CreateAllergyRecord(ctx *gin.Context)

	// GetAllAllergyRecords - Retrieves all allergy records
	GetAllAllergyRecords(ctx *gin.Context)

	// GetAllergyRecordById - Retrieves an allergy record by ID
	GetAllergyRecordById(ctx *gin.Context)

	// DeleteAllergyRecord - Deletes an allergy record
	DeleteAllergyRecord(ctx *gin.Context)
}

// Partial implementation of AllergyRecordsAPI - all functions must be implemented in add-on files
type implAllergyRecordsAPI struct{}

func NewAllergyRecordsAPI() AllergyRecordsAPI {
	return &implAllergyRecordsAPI{}
}

func (this *implAllergyRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/allergy_records", this.CreateAllergyRecord)
	routerGroup.Handle(http.MethodGet, "/allergy_records", this.GetAllAllergyRecords)
	routerGroup.Handle(http.MethodGet, "/allergy_records/:recordId", this.GetAllergyRecordById)
	routerGroup.Handle(http.MethodDelete, "/allergy_records/:recordId", this.DeleteAllergyRecord)
}
