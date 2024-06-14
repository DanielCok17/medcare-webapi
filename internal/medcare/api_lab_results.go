package medcare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LabResultsAPI interface {
	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// GetAllLabResults - Get all lab results
	GetAllLabResults(ctx *gin.Context)

	// GetLabResultById - Get lab result by ID
	GetLabResultById(ctx *gin.Context)

	// UpdateLabResult - Update a lab result
	UpdateLabResult(ctx *gin.Context)
}

// partial implementation of LabResultsAPI - all functions must be implemented in add on files
type implLabResultsAPI struct {
}

func newLabResultsAPI() LabResultsAPI {
	return &implLabResultsAPI{}
}

func (this *implLabResultsAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodGet, "/lab-results", this.GetAllLabResults)
	routerGroup.Handle(http.MethodGet, "/lab-results/:resultId", this.GetLabResultById)
	routerGroup.Handle(http.MethodPut, "/lab-results/:resultId", this.UpdateLabResult)
}

// GetAllLabResults - Get all lab results
func (this *implLabResultsAPI) GetAllLabResults(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetLabResultById - Get lab result by ID
func (this *implLabResultsAPI) GetLabResultById(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateLabResult - Update a lab result
func (this *implLabResultsAPI) UpdateLabResult(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}