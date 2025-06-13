package request

type UpdateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Mobile      string `json:"mobile" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Pan         string `json:"pan" binding:"required"`
}
