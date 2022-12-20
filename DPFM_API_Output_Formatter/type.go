package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey                                 string                                         `json:"connection_key"`
	Result                                        bool                                           `json:"result"`
	RedisKey                                      string                                         `json:"redis_key"`
	Filepath                                      string                                         `json:"filepath"`
	APIStatusCode                                 int                                            `json:"api_status_code"`
	RuntimeSessionID                              string                                         `json:"runtime_session_id"`
	BusinessPartnerID                             *int                                           `json:"business_partner"`
	ServiceLabel                                  string                                         `json:"service_label"`
	BusinessPartnerSupplier                       *BusinessPartnerSupplier                       `json:"BusinessPartnerSupplier,omitempty"`
	BusinessPartnerSupplierPartnerFunctionContact *BusinessPartnerSupplierPartnerFunctionContact `json:"BusinessPartnerSupplierPartnerFunctionContact,omitempty"`
	BusinessPartnerSupplierPartnerPlant           *BusinessPartnerSupplierPartnerPlant           `json:"BusinessPartnerSupplierPartnerPlant,omitempty"`
	BusinessPartnerSupplierFinInst                *BusinessPartnerSupplierFinInst                `json:"BusinessPartnerSupplierFinInst,omitempty"`
	APISchema                                     string                                         `json:"api_schema"`
	Accepter                                      []string                                       `json:"accepter"`
	Deleted                                       bool                                           `json:"deleted"`
}

type BusinessPartnerSupplier struct {
	BusinessPartner int  `json:"BusinessPartner"`
	Supplier        int  `json:"Supplier"`
	ExistenceConf   bool `json:"ExistenceConf"`
}

type BusinessPartnerSupplierPartnerFunctionContact struct {
	BusinessPartner int  `json:"BusinessPartner"`
	Supplier        int  `json:"Supplier"`
	PartnerCounter  int  `json:"PartnerCounter"`
	ContactID       int  `json:"ContactID"`
	ExistenceConf   bool `json:"ExistenceConf"`
}

type BusinessPartnerSupplierPartnerPlant struct {
	BusinessPartner                int    `json:"BusinessPartner"`
	Supplier                       int    `json:"Supplier"`
	PartnerCounter                 int    `json:"PartnerCounter"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner int    `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   int    `json:"PlantCounter"`
	ExistenceConf                  bool   `json:"ExistenceConf"`
}

type BusinessPartnerSupplierFinInst struct {
	BusinessPartner       int    `json:"BusinessPartner"`
	Supplier              int    `json:"Supplier"`
	FinInstIdentification int    `json:"FinInstIdentification"`
	ValidityEndDate       string `json:"ValidityEndDate"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	ExistenceConf         bool   `json:"ExistenceConf"`
}
