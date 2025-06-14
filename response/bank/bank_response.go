package response

type BankResponse struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Details    string `json:"details"`
	Image      string `json:"image"`
	VendorCode string `json:"vendor_code"`
}
