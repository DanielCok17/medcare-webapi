// package medcare

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type AllergyRecordsAPI interface {
// 	// internal registration of api routes
// 	addRoutes(routerGroup *gin.RouterGroup)

// 	// DeleteAllergyRecord - Delete an allergy record
// 	DeleteAllergyRecord(ctx *gin.Context)

// 	// GetAllAllergyRecords - Get all allergy records
// 	GetAllAllergyRecords(ctx *gin.Context)

// 	// GetAllergyRecordById - Get an allergy record by ID
// 	GetAllergyRecordById(ctx *gin.Context)
// }

// // partial implementation of AllergyRecordsAPI - all functions must be implemented in add on files
// type implAllergyRecordsAPI struct {
// }

// func newAllergyRecordsAPI() AllergyRecordsAPI {
// 	return &implAllergyRecordsAPI{}
// }

// func (this *implAllergyRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
// 	routerGroup.Handle(http.MethodDelete, "/allergy-records/:recordId", this.DeleteAllergyRecord)
// 	routerGroup.Handle(http.MethodGet, "/allergy-records", this.GetAllAllergyRecords)
// 	routerGroup.Handle(http.MethodGet, "/allergy-records/:recordId", this.GetAllergyRecordById)
// }

// // DeleteAllergyRecord - Delete an allergy record
// func (this *implAllergyRecordsAPI) DeleteAllergyRecord(ctx *gin.Context) {
// 	ctx.AbortWithStatus(http.StatusNotImplemented)
// }

// // GetAllAllergyRecords - Get all allergy records
// func (this *implAllergyRecordsAPI) GetAllAllergyRecords(ctx *gin.Context) {
// 	ctx.AbortWithStatus(http.StatusNotImplemented)
// }

// // GetAllergyRecordById - Get an allergy record by ID
// func (this *implAllergyRecordsAPI) GetAllergyRecordById(ctx *gin.Context) {
// 	ctx.AbortWithStatus(http.StatusNotImplemented)
// }

package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implAllergyRecordsAPI struct{}

// CreateAllergyRecord - Saves new allergy record
func (this *implAllergyRecordsAPI) CreateAllergyRecord(ctx *gin.Context) {
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

// GetAllAllergyRecords - Get all allergy records
func (this *implAllergyRecordsAPI) GetAllAllergyRecords(ctx *gin.Context) {
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

// GetAllergyRecordById - Get allergy record by ID
func (this *implAllergyRecordsAPI) GetAllergyRecordById(ctx *gin.Context) {
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

// DeleteAllergyRecord - Delete an allergy record
func (this *implAllergyRecordsAPI) DeleteAllergyRecord(ctx *gin.Context) {
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

// addRoutes - Adds routes for allergy records API
func (api *implAllergyRecordsAPI) addRoutes(group *gin.RouterGroup) {
	group.POST("/allergy_records", api.CreateAllergyRecord)
	group.GET("/allergy_records", api.GetAllAllergyRecords)
	group.GET("/allergy_records/:recordId", api.GetAllergyRecordById)
	group.DELETE("/allergy_records/:recordId", api.DeleteAllergyRecord)
}
