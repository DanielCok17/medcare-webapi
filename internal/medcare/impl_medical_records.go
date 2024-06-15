package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateMedicalRecord - Saves new medical record
func (this *implMedicalRecordsAPI) CreateMedicalRecord(ctx *gin.Context) {
	value, exists := ctx.Get("medical_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[MedicalRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	medicalRecord := MedicalRecord{}
	err := ctx.BindJSON(&medicalRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if medicalRecord.Id == "" {
		medicalRecord.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, medicalRecord.Id, &medicalRecord)
	if err != nil {
		switch err {
		case db_service.ErrConflict:
			ctx.JSON(http.StatusConflict, gin.H{
				"status":  "Conflict",
				"message": "Medical record already exists",
				"error":   err.Error(),
			})
		default:
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create medical record in database",
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusCreated, medicalRecord)
}

// GetAllMedicalRecords - Retrieves all medical records
func (this *implMedicalRecordsAPI) GetAllMedicalRecords(ctx *gin.Context) {
	value, exists := ctx.Get("medical_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[MedicalRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	records, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "Bad Gateway",
			"message": "Failed to fetch medical records from database",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, records)
}

// GetMedicalRecordById - Retrieves a medical record by ID
func (this *implMedicalRecordsAPI) GetMedicalRecordById(ctx *gin.Context) {
	value, exists := ctx.Get("medical_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[MedicalRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	recordId := ctx.Param("recordId")
	record, err := db.FindDocument(ctx, recordId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Medical record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// UpdateMedicalRecord - Updates a medical record
func (this *implMedicalRecordsAPI) UpdateMedicalRecord(ctx *gin.Context) {
	value, exists := ctx.Get("medical_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[MedicalRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	recordId := ctx.Param("recordId")
	medicalRecord := MedicalRecord{}
	err := ctx.BindJSON(&medicalRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Set the ID to the one from the URL parameter
	medicalRecord.Id = recordId

	err = db.UpdateDocument(ctx, recordId, &medicalRecord)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Medical record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, medicalRecord)
}

// DeleteMedicalRecord - Deletes a medical record
func (this *implMedicalRecordsAPI) DeleteMedicalRecord(ctx *gin.Context) {
	value, exists := ctx.Get("medical_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[MedicalRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	recordId := ctx.Param("recordId")
	err := db.DeleteDocument(ctx, recordId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Medical record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
