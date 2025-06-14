package request

type RequestOtp struct {
	UserIdentifier string `json:"user_identifier" binding:"required" `
	Slug           string `json:"slug" binding:"required"`
}
