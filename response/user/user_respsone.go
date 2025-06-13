package response

type UserResponse struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	Pan         string `json:"pan"`
}
