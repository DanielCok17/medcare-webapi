package medcare

type Condition struct {
	Value string `json:"value"`

	Code string `json:"code,omitempty"`

	// Link to encyclopedical explanation of the patient's condition
	string `json:"reference,omitempty"`

	TypicalDurationMinutes int32 `json:"typicalDurationMinutes,omitempty"`
}
