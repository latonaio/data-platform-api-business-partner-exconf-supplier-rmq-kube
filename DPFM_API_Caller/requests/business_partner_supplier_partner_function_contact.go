package requests

type BusinessPartnerSupplierPartnerFunctionContact struct {
	BusinessPartner *int `json:"BusinessPartner"`
	Supplier        *int `json:"Supplier"`
	PartnerCounter  *int `json:"PartnerCounter"`
	ContactID       *int `json:"ContactID"`
}
