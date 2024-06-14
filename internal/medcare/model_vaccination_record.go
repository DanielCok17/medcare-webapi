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

type VaccinationRecord struct {

	// Unique id of the vaccination record
	Id string `json:"id"`

	// Unique identifier of the patient
	PatientId string `json:"patientId"`

	Vaccine string `json:"vaccine"`

	Date string `json:"date"`
}