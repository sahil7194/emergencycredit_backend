package request

type SavePersonalDetails struct {
	FullName      string `json:"full_name" binding:"required"`
	DateOfBirth   string `json:"date_of_birth" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
	PanCardNumber string `json:"pan_card_number" binding:"required"`
}
