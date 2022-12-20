package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SupplierSDC) ConvertToBusinessPartnerSupplier() *requests.BusinessPartnerSupplier {
	data := sdc.BusinessPartnerSupplier
	return &requests.BusinessPartnerSupplier{
		BusinessPartner: data.BusinessPartner,
		Supplier:        data.Supplier,
	}
}

func (sdc *PartnerFunctionContactSDC) ConvertToBusinessPartnerSupplierPartnerFunctionContact() *requests.BusinessPartnerSupplierPartnerFunctionContact {
	data := sdc.BusinessPartnerSupplierPartnerFunctionContact
	return &requests.BusinessPartnerSupplierPartnerFunctionContact{
		BusinessPartner: data.BusinessPartner,
		Supplier:        data.Supplier,
		PartnerCounter:  data.PartnerCounter,
		ContactID:       data.ContactID,
	}
}

func (sdc *PartnerPlantSDC) ConvertToBusinessPartnerSupplierPartnerPlant() *requests.BusinessPartnerSupplierPartnerPlant {
	data := sdc.BusinessPartnerSupplierPartnerPlant
	return &requests.BusinessPartnerSupplierPartnerPlant{
		BusinessPartner:                data.BusinessPartner,
		Supplier:                       data.Supplier,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		PlantCounter:                   data.PlantCounter,
	}
}

func (sdc *FinInstSDC) ConvertToBusinessPartnerSupplierFinInst() *requests.BusinessPartnerSupplierFinInst {
	data := sdc.BusinessPartnerSupplierFinInst
	return &requests.BusinessPartnerSupplierFinInst{
		BusinessPartner:       data.BusinessPartner,
		Supplier:              data.Supplier,
		FinInstIdentification: data.FinInstIdentification,
		ValidityEndDate:       data.ValidityEndDate,
		ValidityStartDate:     data.ValidityStartDate,
	}
}
