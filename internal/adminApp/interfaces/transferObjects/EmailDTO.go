package transferobjects

type EmailDTO struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	To      string      `json:"to,omitempty"`
	Subject string      `json:"subject,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
