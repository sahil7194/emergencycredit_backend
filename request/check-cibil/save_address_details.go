package request

type SaveAddressDetails struct {
	Slug    string `json:"slug"`
	Address string `json:"address"`
	PinCode string `json:"pin_code"`
	CityID  string `json:"city_id"`
	StateID string `json:"state_id"`
}
