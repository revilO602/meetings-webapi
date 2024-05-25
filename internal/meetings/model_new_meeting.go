/*
 * Meetings Api
 *
 * Online meetings between doctors and patients management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: leoentiev.oliver@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package meetings

type NewMeeting struct {

	// Name of doctor for the meeting
	DoctorName string `json:"doctorName"`

	// Name of patient for the meeting
	PatientName string `json:"patientName"`

	// Date when meeting will take place
	Date string `json:"date"`

	// Start time of meeting
	StartTime string `json:"startTime"`

	// End time of meeting
	EndTime string `json:"endTime"`

	// Whether meeting is important
	Important bool `json:"important"`

	// Online platform for the meeting
	Platform string `json:"platform"`

	// Patients symptoms
	Symptoms string `json:"symptoms"`

	// Diagnosis for the patients problems
	Diagnosis string `json:"diagnosis"`

	// Extra notes from the doctor about the patient
	Notes string `json:"notes"`
}
