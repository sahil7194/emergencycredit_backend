package request

type UpdateBankRequest struct {
	Name    string `json:"name" binding:"required"`
	Details string `json:"details" binding:"required"`
}
