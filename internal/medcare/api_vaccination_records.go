package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VaccinationRecordsAPI interface {
	// internal registration of API routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateVaccinationRecord - Saves new vaccination record
	CreateVaccinationRecord(ctx *gin.Context)

	// GetAllVaccinationRecords - Retrieves all vaccination records
	GetAllVaccinationRecords(ctx *gin.Context)

	// GetVaccinationRecordById - Retrieves a vaccination record by ID
	GetVaccinationRecordById(ctx *gin.Context)

	// UpdateVaccinationRecord - Updates a vaccination record
	UpdateVaccinationRecord(ctx *gin.Context)

	// DeleteVaccinationRecord - Deletes a vaccination record
	DeleteVaccinationRecord(ctx *gin.Context)
}

// Partial implementation of VaccinationRecordsAPI - all functions must be implemented in add-on files
type implVaccinationRecordsAPI struct{}

func NewVaccinationRecordsAPI() VaccinationRecordsAPI {
	return &implVaccinationRecordsAPI{}
}

func (this *implVaccinationRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/vaccination_records", this.CreateVaccinationRecord)
	routerGroup.Handle(http.MethodGet, "/vaccination_records", this.GetAllVaccinationRecords)
	routerGroup.Handle(http.MethodGet, "/vaccination_records/:recordId", this.GetVaccinationRecordById)
	routerGroup.Handle(http.MethodPut, "/vaccination_records/:recordId", this.UpdateVaccinationRecord)
	routerGroup.Handle(http.MethodDelete, "/vaccination_records/:recordId", this.DeleteVaccinationRecord)
}
