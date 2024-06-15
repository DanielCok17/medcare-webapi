package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MedicalRecordsAPI interface {
	// internal registration of API routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateMedicalRecord - Saves new medical record
	CreateMedicalRecord(ctx *gin.Context)

	// GetAllMedicalRecords - Retrieves all medical records
	GetAllMedicalRecords(ctx *gin.Context)

	// GetMedicalRecordById - Retrieves a medical record by ID
	GetMedicalRecordById(ctx *gin.Context)

	// UpdateMedicalRecord - Updates a medical record
	UpdateMedicalRecord(ctx *gin.Context)

	// DeleteMedicalRecord - Deletes a medical record
	DeleteMedicalRecord(ctx *gin.Context)
}

// Partial implementation of MedicalRecordsAPI - all functions must be implemented in add-on files
type implMedicalRecordsAPI struct{}

func NewMedicalRecordsAPI() MedicalRecordsAPI {
	return &implMedicalRecordsAPI{}
}

func (this *implMedicalRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/medical_records", this.CreateMedicalRecord)
	routerGroup.Handle(http.MethodGet, "/medical_records", this.GetAllMedicalRecords)
	routerGroup.Handle(http.MethodGet, "/medical_records/:recordId", this.GetMedicalRecordById)
	routerGroup.Handle(http.MethodPut, "/medical_records/:recordId", this.UpdateMedicalRecord)
	routerGroup.Handle(http.MethodDelete, "/medical_records/:recordId", this.DeleteMedicalRecord)
}
