/*
 * MedCare API
 *
 * MedCare system API
 *
 * API version: 1.0.0
 * Contact: <your_email>
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package medcare_api

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type MedicalRecordsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // GetAllMedicalRecords - Get all medical records
   GetAllMedicalRecords(ctx *gin.Context)

 }

 // partial implementation of MedicalRecordsAPI - all functions must be implemented in add on files
type implMedicalRecordsAPI struct {

}

func newMedicalRecordsAPI() MedicalRecordsAPI {
  return &implMedicalRecordsAPI{}
}

func (this *implMedicalRecordsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodGet, "/medical-records", this.GetAllMedicalRecords)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // GetAllMedicalRecords - Get all medical records
// func (this *implMedicalRecordsAPI) GetAllMedicalRecords(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

