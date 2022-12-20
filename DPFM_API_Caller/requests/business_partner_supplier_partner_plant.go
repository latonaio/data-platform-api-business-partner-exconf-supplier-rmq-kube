package requests

type BusinessPartnerSupplierPartnerPlant struct {
	BusinessPartner                *int    `json:"BusinessPartner"`
	Supplier                       *int    `json:"Supplier"`
	PartnerCounter                 *int    `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   *int    `json:"PlantCounter"`
}
