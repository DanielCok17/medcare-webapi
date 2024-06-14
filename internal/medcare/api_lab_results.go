package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implLabResultsAPI struct{}

// LabResultsAPI interface definition
type LabResultsAPI interface {
	addRoutes(routerGroup *gin.RouterGroup)
	CreateLabResult(ctx *gin.Context)
	GetAllLabResults(ctx *gin.Context)
	GetLabResultById(ctx *gin.Context)
	UpdateLabResult(ctx *gin.Context)
	DeleteLabResult(ctx *gin.Context)
}

// NewLabResultsAPI constructor
func NewLabResultsAPI() LabResultsAPI {
	return &implLabResultsAPI{}
}

// addRoutes implementation
func (api *implLabResultsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/lab_results", api.CreateLabResult)
	routerGroup.GET("/lab_results", api.GetAllLabResults)
	routerGroup.GET("/lab_results/:recordId", api.GetLabResultById)
	routerGroup.PUT("/lab_results/:recordId", api.UpdateLabResult)
	routerGroup.DELETE("/lab_results/:recordId", api.DeleteLabResult)
}

// CreateLabResult implementation
func (api *implLabResultsAPI) CreateLabResult(ctx *gin.Context) {
	value, exists := ctx.Get("lab_result_db_service")
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

	db, ok := value.(db_service.DbService[LabResult])
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

	labResult := LabResult{}
	err := ctx.BindJSON(&labResult)
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

	if labResult.Id == "" {
		labResult.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, labResult.Id, &labResult)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			labResult,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Lab result already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create lab result in database",
				"error":   err.Error(),
			},
		)
	}
}

// GetAllLabResults implementation
func (api *implLabResultsAPI) GetAllLabResults(ctx *gin.Context) {
	value, exists := ctx.Get("lab_result_db_service")
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

	db, ok := value.(db_service.DbService[LabResult])
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
				"message": "Failed to fetch lab results from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, records)
}

// GetLabResultById implementation
func (api *implLabResultsAPI) GetLabResultById(ctx *gin.Context) {
	value, exists := ctx.Get("lab_result_db_service")
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

	db, ok := value.(db_service.DbService[LabResult])
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
				"message": "Lab result not found",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// UpdateLabResult implementation
func (api *implLabResultsAPI) UpdateLabResult(ctx *gin.Context) {
	value, exists := ctx.Get("lab_result_db_service")
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

	db, ok := value.(db_service.DbService[LabResult])
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
	labResult := LabResult{}
	err := ctx.BindJSON(&labResult)
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

	// Set the ID to the one from the URL parameter
	labResult.Id = recordId

	err = db.UpdateDocument(ctx, recordId, &labResult)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Lab result not found",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, labResult)
}

// DeleteLabResult implementation
func (api *implLabResultsAPI) DeleteLabResult(ctx *gin.Context) {
	value, exists := ctx.Get("lab_result_db_service")
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

	db, ok := value.(db_service.DbService[LabResult])
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
				"message": "Lab result not found",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.Status(http.StatusNoContent)
}
