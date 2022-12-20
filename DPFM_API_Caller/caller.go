package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["BusinessPartnerSupplier"]
	if ok {
		input := &dpfm_api_input_reader.SupplierSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confBusinessPartnerSupplier(input)
		goto endProcess
	}
	_, ok = input["BusinessPartnerSupplierPartnerFunctionContact"]
	if ok {
		input := &dpfm_api_input_reader.PartnerFunctionContactSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfBusinessPartnerSupplierPartnerFunctionContact(input)
		goto endProcess
	}
	_, ok = input["BusinessPartnerSupplierPartnerPlant"]
	if ok {
		input := &dpfm_api_input_reader.PartnerPlantSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfBusinessPartnerSupplierPartnerPlant(input)
		goto endProcess
	}
	_, ok = input["BusinessPartnerSupplierFinInst"]
	if ok {
		input := &dpfm_api_input_reader.FinInstSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfBusinessPartnerSupplierFinInst(input)
		goto endProcess
	}

	err = xerrors.Errorf("can not get exconf check target")
endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confBusinessPartnerSupplier(input *dpfm_api_input_reader.SupplierSDC) *dpfm_api_output_formatter.BusinessPartnerSupplier {
	exconf := dpfm_api_output_formatter.BusinessPartnerSupplier{
		ExistenceConf: false,
	}
	if input.BusinessPartnerSupplier.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplier.Supplier == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerSupplier{
		BusinessPartner: *input.BusinessPartnerSupplier.BusinessPartner,
		Supplier:        *input.BusinessPartnerSupplier.Supplier,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerSupplier
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data 
		WHERE (businessPartner, supplier) = (?, ?);`, exconf.BusinessPartner, exconf.Supplier,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerSupplierPartnerFunctionContact(input *dpfm_api_input_reader.PartnerFunctionContactSDC) *dpfm_api_output_formatter.BusinessPartnerSupplierPartnerFunctionContact {
	exconf := dpfm_api_output_formatter.BusinessPartnerSupplierPartnerFunctionContact{
		ExistenceConf: false,
	}
	if input.BusinessPartnerSupplierPartnerFunctionContact.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerFunctionContact.Supplier == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerFunctionContact.PartnerCounter == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerFunctionContact.ContactID == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerSupplierPartnerFunctionContact{
		BusinessPartner: *input.BusinessPartnerSupplierPartnerFunctionContact.BusinessPartner,
		Supplier:        *input.BusinessPartnerSupplierPartnerFunctionContact.Supplier,
		PartnerCounter:  *input.BusinessPartnerSupplierPartnerFunctionContact.PartnerCounter,
		ContactID:       *input.BusinessPartnerSupplierPartnerFunctionContact.ContactID,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerSupplierPartnerFunctionContact
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_function_contact_data 
		WHERE (businessPartner, supplier, partnerCounter, contactID) = (?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Supplier, exconf.PartnerCounter, exconf.ContactID,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerSupplierPartnerPlant(input *dpfm_api_input_reader.PartnerPlantSDC) *dpfm_api_output_formatter.BusinessPartnerSupplierPartnerPlant {
	exconf := dpfm_api_output_formatter.BusinessPartnerSupplierPartnerPlant{
		ExistenceConf: false,
	}
	if input.BusinessPartnerSupplierPartnerPlant.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerPlant.Supplier == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerPlant.PartnerCounter == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerPlant.PartnerFunction == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerPlant.PartnerFunctionBusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierPartnerPlant.PlantCounter == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerSupplierPartnerPlant{
		BusinessPartner:                *input.BusinessPartnerSupplierPartnerPlant.BusinessPartner,
		Supplier:                       *input.BusinessPartnerSupplierPartnerPlant.Supplier,
		PartnerCounter:                 *input.BusinessPartnerSupplierPartnerPlant.PartnerCounter,
		PartnerFunction:                *input.BusinessPartnerSupplierPartnerPlant.PartnerFunction,
		PartnerFunctionBusinessPartner: *input.BusinessPartnerSupplierPartnerPlant.PartnerFunctionBusinessPartner,
		PlantCounter:                   *input.BusinessPartnerSupplierPartnerPlant.PlantCounter,
		ExistenceConf:                  false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerSupplierPartnerPlant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_plant_data 
		WHERE (businessPartner, supplier, partnerCounter, partnerFunction, partnerFunctionBusinessPartner, plantCounter) = (?, ?, ?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Supplier, exconf.PartnerCounter, exconf.PartnerFunction, exconf.PartnerFunctionBusinessPartner, exconf.PlantCounter,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerSupplierFinInst(input *dpfm_api_input_reader.FinInstSDC) *dpfm_api_output_formatter.BusinessPartnerSupplierFinInst {
	exconf := dpfm_api_output_formatter.BusinessPartnerSupplierFinInst{
		ExistenceConf: false,
	}
	if input.BusinessPartnerSupplierFinInst.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierFinInst.Supplier == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierFinInst.FinInstIdentification == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierFinInst.ValidityEndDate == nil {
		return &exconf
	}
	if input.BusinessPartnerSupplierFinInst.ValidityStartDate == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerSupplierFinInst{
		BusinessPartner:       *input.BusinessPartnerSupplierFinInst.BusinessPartner,
		Supplier:              *input.BusinessPartnerSupplierFinInst.Supplier,
		FinInstIdentification: *input.BusinessPartnerSupplierFinInst.FinInstIdentification,
		ValidityEndDate:       *input.BusinessPartnerSupplierFinInst.ValidityEndDate,
		ValidityStartDate:     *input.BusinessPartnerSupplierFinInst.ValidityStartDate,
		ExistenceConf:         false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerSupplierFinInst
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_fin_inst_data 
		WHERE (businessPartner, supplier, finInstIdentification, validityEndDate, validityStartDate) = (?, ?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Supplier, exconf.FinInstIdentification, exconf.ValidityEndDate, exconf.ValidityStartDate,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
