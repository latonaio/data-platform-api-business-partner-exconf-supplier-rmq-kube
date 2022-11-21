package requests

type BusinessPartnerSupplier struct {
	BusinessPartner       *int    `json:"BusinessPartner"`
	Supplier              *int `json:"Supplier"`
}
