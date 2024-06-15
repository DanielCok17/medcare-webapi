package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LabResultsAPI interface {
	// internal registration of API routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateLabResult - Saves new lab result
	CreateLabResult(ctx *gin.Context)

	// GetAllLabResults - Retrieves all lab results
	GetAllLabResults(ctx *gin.Context)

	// GetLabResultById - Retrieves a lab result by ID
	GetLabResultById(ctx *gin.Context)

	// UpdateLabResult - Updates a lab result
	UpdateLabResult(ctx *gin.Context)

	// DeleteLabResult - Deletes a lab result
	DeleteLabResult(ctx *gin.Context)
}

// Partial implementation of LabResultsAPI - all functions must be implemented in add-on files
type implLabResultsAPI struct{}

func NewLabResultsAPI() LabResultsAPI {
	return &implLabResultsAPI{}
}

func (this *implLabResultsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/lab_results", this.CreateLabResult)
	routerGroup.Handle(http.MethodGet, "/lab_results", this.GetAllLabResults)
	routerGroup.Handle(http.MethodGet, "/lab_results/:recordId", this.GetLabResultById)
	routerGroup.Handle(http.MethodPut, "/lab_results/:recordId", this.UpdateLabResult)
	routerGroup.Handle(http.MethodDelete, "/lab_results/:recordId", this.DeleteLabResult)
}
