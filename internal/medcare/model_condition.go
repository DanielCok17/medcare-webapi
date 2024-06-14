package medcare

type Condition struct {
	Value string `json:"value"`

	Code string `json:"code,omitempty"`

	// Link to encyclopedic explanation of the patient's condition
	Reference string `json:"reference,omitempty"`

	TypicalDurationMinutes int32 `json:"typicalDurationMinutes,omitempty"`
}
