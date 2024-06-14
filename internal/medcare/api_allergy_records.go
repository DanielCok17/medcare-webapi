package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implAllergyRecordsAPI struct{}

// AllergyRecordsAPI interface definition
type AllergyRecordsAPI interface {
	addRoutes(routerGroup *gin.RouterGroup)
	CreateAllergyRecord(ctx *gin.Context)
	GetAllAllergyRecords(ctx *gin.Context)
	GetAllergyRecordById(ctx *gin.Context)
	DeleteAllergyRecord(ctx *gin.Context)
}

// NewAllergyRecordsAPI constructor
func NewAllergyRecordsAPI() AllergyRecordsAPI {
	return &implAllergyRecordsAPI{}
}

// addRoutes implementation
func (api *implAllergyRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/allergy_records", api.CreateAllergyRecord)
	routerGroup.GET("/allergy_records", api.GetAllAllergyRecords)
	routerGroup.GET("/allergy_records/:recordId", api.GetAllergyRecordById)
	routerGroup.DELETE("/allergy_records/:recordId", api.DeleteAllergyRecord)
}

// CreateAllergyRecord implementation
func (api *implAllergyRecordsAPI) CreateAllergyRecord(ctx *gin.Context) {
	value, exists := ctx.Get("allergy_db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AllergyRecord])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	allergyRecord := AllergyRecord{}
	err := ctx.BindJSON(&allergyRecord)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if allergyRecord.Id == "" {
		allergyRecord.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, allergyRecord.Id, &allergyRecord)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			allergyRecord,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Allergy record already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create allergy record in database",
				"error":   err.Error(),
			},
		)
	}
}

// GetAllAllergyRecords implementation
func (api *implAllergyRecordsAPI) GetAllAllergyRecords(ctx *gin.Context) {
	value, exists := ctx.Get("allergy_db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AllergyRecord])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	records, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to fetch allergy records from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, records)
}

// GetAllergyRecordById implementation
func (api *implAllergyRecordsAPI) GetAllergyRecordById(ctx *gin.Context) {
	value, exists := ctx.Get("allergy_db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AllergyRecord])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	recordId := ctx.Param("recordId")
	record, err := db.FindDocument(ctx, recordId)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Allergy record not found",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteAllergyRecord implementation
func (api *implAllergyRecordsAPI) DeleteAllergyRecord(ctx *gin.Context) {
	value, exists := ctx.Get("allergy_db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[AllergyRecord])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	recordId := ctx.Param("recordId")
	err := db.DeleteDocument(ctx, recordId)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Allergy record not found",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.Status(http.StatusNoContent)
}
