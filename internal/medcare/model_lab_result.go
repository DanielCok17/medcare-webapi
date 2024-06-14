/*
 * MedCare API
 *
 * MedCare system API
 *
 * API version: 1.0.0
 * Contact: <your_email>
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package medcare

type LabResult struct {

	// Unique id of the lab result
	Id string `json:"id"`

	// Unique identifier of the patient
	PatientId string `json:"patientId"`

	TestType string `json:"testType"`

	Result string `json:"result"`
}