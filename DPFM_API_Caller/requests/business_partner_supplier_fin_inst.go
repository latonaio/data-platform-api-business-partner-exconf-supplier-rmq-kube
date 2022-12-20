package requests

type BusinessPartnerSupplierFinInst struct {
	BusinessPartner       *int    `json:"BusinessPartner"`
	Supplier              *int    `json:"Supplier"`
	FinInstIdentification *int    `json:"FinInstIdentification"`
	ValidityEndDate       *string `json:"ValidityEndDate"`
	ValidityStartDate     *string `json:"ValidityStartDate"`
}
