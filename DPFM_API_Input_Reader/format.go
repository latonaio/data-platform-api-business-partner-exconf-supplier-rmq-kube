package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBusinessPartnerSupplier() *requests.BusinessPartnerSupplier {
	data := sdc.BusinessPartnerSupplier
	return &requests.BusinessPartnerSupplier{
		BusinessPartner:    data.BusinessPartner,
		Supplier:           data.Supplier,
	}
}
