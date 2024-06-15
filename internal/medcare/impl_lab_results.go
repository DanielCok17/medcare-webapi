package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateLabResult - Saves new lab result
func (this *implLabResultsAPI) CreateLabResult(ctx *gin.Context) {
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

// GetAllLabResults - Retrieves all lab results
func (this *implLabResultsAPI) GetAllLabResults(ctx *gin.Context) {
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

// GetLabResultById - Retrieves a lab result by ID
func (this *implLabResultsAPI) GetLabResultById(ctx *gin.Context) {
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

// UpdateLabResult - Updates a lab result
func (this *implLabResultsAPI) UpdateLabResult(ctx *gin.Context) {
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

// DeleteLabResult - Deletes a lab result
func (this *implLabResultsAPI) DeleteLabResult(ctx *gin.Context) {
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
