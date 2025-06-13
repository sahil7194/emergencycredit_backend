package request

type CreateBankRequest struct {
	Name    string `json:"name" binding:"required"`
	Details string `json:"details" binding:"required"`
}
