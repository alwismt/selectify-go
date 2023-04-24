package transferobjects

// Create a DTO struct to represent the data required by the auth service layer
type SignInDTO struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	Redirect string `json:"redirect"`
}
