package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implVaccinationRecordsAPI struct{}

// VaccinationRecordsAPI interface definition
type VaccinationRecordsAPI interface {
	addRoutes(routerGroup *gin.RouterGroup)
	CreateVaccinationRecord(ctx *gin.Context)
	GetAllVaccinationRecords(ctx *gin.Context)
	GetVaccinationRecordById(ctx *gin.Context)
	UpdateVaccinationRecord(ctx *gin.Context)
	DeleteVaccinationRecord(ctx *gin.Context)
}

// NewVaccinationRecordsAPI constructor
func NewVaccinationRecordsAPI() VaccinationRecordsAPI {
	return &implVaccinationRecordsAPI{}
}

// addRoutes implementation
func (api *implVaccinationRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/vaccination_records", api.CreateVaccinationRecord)
	routerGroup.GET("/vaccination_records", api.GetAllVaccinationRecords)
	routerGroup.GET("/vaccination_records/:recordId", api.GetVaccinationRecordById)
	routerGroup.PUT("/vaccination_records/:recordId", api.UpdateVaccinationRecord)
	routerGroup.DELETE("/vaccination_records/:recordId", api.DeleteVaccinationRecord)
}

// CreateVaccinationRecord implementation
func (api *implVaccinationRecordsAPI) CreateVaccinationRecord(ctx *gin.Context) {
	value, exists := ctx.Get("vaccination_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[VaccinationRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	vaccinationRecord := VaccinationRecord{}
	err := ctx.BindJSON(&vaccinationRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if vaccinationRecord.Id == "" {
		vaccinationRecord.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, vaccinationRecord.Id, &vaccinationRecord)
	if err != nil {
		switch err {
		case db_service.ErrConflict:
			ctx.JSON(http.StatusConflict, gin.H{
				"status":  "Conflict",
				"message": "Vaccination record already exists",
				"error":   err.Error(),
			})
		default:
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create vaccination record in database",
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusCreated, vaccinationRecord)
}

// GetAllVaccinationRecords implementation
func (api *implVaccinationRecordsAPI) GetAllVaccinationRecords(ctx *gin.Context) {
	value, exists := ctx.Get("vaccination_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[VaccinationRecord])
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
			"message": "Failed to fetch vaccination records from database",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, records)
}

// GetVaccinationRecordById implementation
func (api *implVaccinationRecordsAPI) GetVaccinationRecordById(ctx *gin.Context) {
	value, exists := ctx.Get("vaccination_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[VaccinationRecord])
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
			"message": "Vaccination record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// UpdateVaccinationRecord implementation
func (api *implVaccinationRecordsAPI) UpdateVaccinationRecord(ctx *gin.Context) {
	value, exists := ctx.Get("vaccination_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[VaccinationRecord])
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db context is not of required type",
			"error":   "cannot cast db context to db_service.DbService",
		})
		return
	}

	recordId := ctx.Param("recordId")
	vaccinationRecord := VaccinationRecord{}
	err := ctx.BindJSON(&vaccinationRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Set the ID to the one from the URL parameter
	vaccinationRecord.Id = recordId

	err = db.UpdateDocument(ctx, recordId, &vaccinationRecord)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Vaccination record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, vaccinationRecord)
}

// DeleteVaccinationRecord implementation
func (api *implVaccinationRecordsAPI) DeleteVaccinationRecord(ctx *gin.Context) {
	value, exists := ctx.Get("vaccination_record_db_service")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "db not found",
			"error":   "db not found",
		})
		return
	}

	db, ok := value.(db_service.DbService[VaccinationRecord])
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
			"message": "Vaccination record not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
