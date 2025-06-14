package request

type ValidateOtp struct {
	UserIdentifier string `json:"user_identifier"`
	Slug           string `json:"slug"`
	Otp            string `json:"otp"`
}
