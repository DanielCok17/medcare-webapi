// // package medcare

// // import (
// // 	"net/http"

// // 	"github.com/gin-gonic/gin"
// // )

// // type LabResultsAPI interface {
// // 	// internal registration of api routes
// // 	addRoutes(routerGroup *gin.RouterGroup)

// // 	// GetAllLabResults - Get all lab results
// // 	GetAllLabResults(ctx *gin.Context)

// // 	// GetLabResultById - Get lab result by ID
// // 	GetLabResultById(ctx *gin.Context)

// // 	// UpdateLabResult - Update a lab result
// // 	UpdateLabResult(ctx *gin.Context)
// // }

// // // partial implementation of LabResultsAPI - all functions must be implemented in add on files
// // type implLabResultsAPI struct {
// // }

// // func newLabResultsAPI() LabResultsAPI {
// // 	return &implLabResultsAPI{}
// // }

// // func (this *implLabResultsAPI) addRoutes(routerGroup *gin.RouterGroup) {
// // 	routerGroup.Handle(http.MethodGet, "/lab-results", this.GetAllLabResults)
// // 	routerGroup.Handle(http.MethodGet, "/lab-results/:resultId", this.GetLabResultById)
// // 	routerGroup.Handle(http.MethodPut, "/lab-results/:resultId", this.UpdateLabResult)
// // }

// // // GetAllLabResults - Get all lab results
// // func (this *implLabResultsAPI) GetAllLabResults(ctx *gin.Context) {
// // 	ctx.AbortWithStatus(http.StatusNotImplemented)
// // }

// // // GetLabResultById - Get lab result by ID
// // func (this *implLabResultsAPI) GetLabResultById(ctx *gin.Context) {
// // 	ctx.AbortWithStatus(http.StatusNotImplemented)
// // }

// // // UpdateLabResult - Update a lab result
// // func (this *implLabResultsAPI) UpdateLabResult(ctx *gin.Context) {
// // 	ctx.AbortWithStatus(http.StatusNotImplemented)
// // }

// package medcare

// import (
// 	"net/http"

// 	"github.com/DanielCok17/medcare-webapi/internal/db_service"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type implLabResultsAPI struct{}

// // CreateLabResult - Saves new lab result
// func (this *implLabResultsAPI) CreateLabResult(ctx *gin.Context) {
// 	value, exists := ctx.Get("db_service")
// 	if !exists {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db not found",
// 				"error":   "db not found",
// 			})
// 		return
// 	}

// 	db, ok := value.(db_service.DbService[LabResult])
// 	if !ok {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db context is not of required type",
// 				"error":   "cannot cast db context to db_service.DbService",
// 			})
// 		return
// 	}

// 	labResult := LabResult{}
// 	err := ctx.BindJSON(&labResult)
// 	if err != nil {
// 		ctx.JSON(
// 			http.StatusBadRequest,
// 			gin.H{
// 				"status":  "Bad Request",
// 				"message": "Invalid request body",
// 				"error":   err.Error(),
// 			})
// 		return
// 	}

// 	if labResult.Id == "" {
// 		labResult.Id = uuid.New().String()
// 	}

// 	err = db.CreateDocument(ctx, labResult.Id, &labResult)

// 	switch err {
// 	case nil:
// 		ctx.JSON(
// 			http.StatusCreated,
// 			labResult,
// 		)
// 	case db_service.ErrConflict:
// 		ctx.JSON(
// 			http.StatusConflict,
// 			gin.H{
// 				"status":  "Conflict",
// 				"message": "Lab result already exists",
// 				"error":   err.Error(),
// 			},
// 		)
// 	default:
// 		ctx.JSON(
// 			http.StatusBadGateway,
// 			gin.H{
// 				"status":  "Bad Gateway",
// 				"message": "Failed to create lab result in database",
// 				"error":   err.Error(),
// 			},
// 		)
// 	}
// }

// // GetAllLabResults - Get all lab results
// func (this *implLabResultsAPI) GetAllLabResults(ctx *gin.Context) {
// 	value, exists := ctx.Get("db_service")
// 	if !exists {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db not found",
// 				"error":   "db not found",
// 			})
// 		return
// 	}

// 	db, ok := value.(db_service.DbService[LabResult])
// 	if !ok {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db context is not of required type",
// 				"error":   "cannot cast db context to db_service.DbService",
// 			})
// 		return
// 	}

// 	records, err := db.FindAllDocuments(ctx)
// 	if err != nil {
// 		ctx.JSON(
// 			http.StatusBadGateway,
// 			gin.H{
// 				"status":  "Bad Gateway",
// 				"message": "Failed to fetch lab results from database",
// 				"error":   err.Error(),
// 			},
// 		)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, records)
// }

// // GetLabResultById - Get lab result by ID
// func (this *implLabResultsAPI) GetLabResultById(ctx *gin.Context) {
// 	value, exists := ctx.Get("db_service")
// 	if !exists {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db not found",
// 				"error":   "db not found",
// 			})
// 		return
// 	}

// 	db, ok := value.(db_service.DbService[LabResult])
// 	if !ok {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db context is not of required type",
// 				"error":   "cannot cast db context to db_service.DbService",
// 			})
// 		return
// 	}

// 	recordId := ctx.Param("recordId")
// 	record, err := db.FindDocument(ctx, recordId)
// 	if err != nil {
// 		ctx.JSON(
// 			http.StatusNotFound,
// 			gin.H{
// 				"status":  "Not Found",
// 				"message": "Lab result not found",
// 				"error":   err.Error(),
// 			},
// 		)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, record)
// }

// // DeleteLabResult - Delete a lab result
// func (this *implLabResultsAPI) DeleteLabResult(ctx *gin.Context) {
// 	value, exists := ctx.Get("db_service")
// 	if !exists {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db not found",
// 				"error":   "db not found",
// 			})
// 		return
// 	}

// 	db, ok := value.(db_service.DbService[LabResult])
// 	if !ok {
// 		ctx.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  "Internal Server Error",
// 				"message": "db context is not of required type",
// 				"error":   "cannot cast db context to db_service.DbService",
// 			})
// 		return
// 	}

// 	recordId := ctx.Param("recordId")
// 	err := db.DeleteDocument(ctx, recordId)
// 	if err != nil {
// 		ctx.JSON(
// 			http.StatusNotFound,
// 			gin.H{
// 				"status":  "Not Found",
// 				"message": "Lab result not found",
// 				"error":   err.Error(),
// 			},
// 		)
// 		return
// 	}

// 	ctx.Status(http.StatusNoContent)
// }

// // addRoutes - Adds routes for lab results API
// func (api *implLabResultsAPI) addRoutes(group *gin.RouterGroup) {
// 	group.POST("/lab_results", api.CreateLabResult)
// 	group.GET("/lab_results", api.GetAllLabResults)
// 	group.GET("/lab_results/:resultId", api.GetLabResultById)
// 	group.DELETE("/lab_results/:resultId", api.DeleteLabResult)
// }

package medcare

import (
	"net/http"

	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implLabResultsAPI struct{}

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

// GetAllLabResults - Get all lab results
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

// GetLabResultById - Get lab result by ID
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

// DeleteLabResult - Delete a lab result
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

// / UpdateLabResult - Update a lab result
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

// addRoutes - Adds routes for lab results API
func (api *implLabResultsAPI) addRoutes(group *gin.RouterGroup) {
	group.POST("/lab_results", api.CreateLabResult)
	group.GET("/lab_results", api.GetAllLabResults)
	group.GET("/lab_results/:recordId", api.GetLabResultById)
	group.PUT("/lab_results/:recordId", api.UpdateLabResult)
	group.DELETE("/lab_results/:recordId", api.DeleteLabResult)
}
