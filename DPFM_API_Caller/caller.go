package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-exconf-supplier-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-business-partner-exconf-supplier-rmq-kube/database"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
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

func (e *ExistenceConf) Conf(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.BusinessPartnerSupplier {
	businessPartner := *input.BusinessPartnerSupplier.BusinessPartner
	supplier := *input.BusinessPartnerSupplier.Supplier
	notKeyExistence := make([]dpfm_api_output_formatter.BusinessPartnerSupplier, 0, 1)
	KeyExistence := make([]dpfm_api_output_formatter.BusinessPartnerSupplier, 0, 1)

	existData := &dpfm_api_output_formatter.BusinessPartnerSupplier{
		BusinessPartner: businessPartner,
		Supplier:        supplier,
		ExistenceConf:   false,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if !e.confBusinessPartnerSupplier(businessPartner, supplier) {
			notKeyExistence = append(
				notKeyExistence,
				dpfm_api_output_formatter.BusinessPartnerSupplier{businessPartner, supplier, false},
			)
			return
		}
		KeyExistence = append(KeyExistence, dpfm_api_output_formatter.BusinessPartnerSupplier{businessPartner, supplier, true})
	}()

	wg.Wait()

	if len(KeyExistence) == 0 {
		return existData
	}
	if len(notKeyExistence) > 0 {
		return existData
	}

	existData.ExistenceConf = true
	return existData
}

func (e *ExistenceConf) confBusinessPartnerSupplier(businessPartner int, supplier int) bool {
	rows, err := e.db.Query(
		`SELECT Supplier 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data 
		WHERE (BusinessPartner, supplier) = (?, ?);`, businessPartner, supplier,
	)
	if err != nil {
		e.l.Error(err)
		return false
	}
	if err != nil {
		e.l.Error(err)
		return false
	}

	for rows.Next() {
		var businessPartner int
		var supplier int
		err := rows.Scan(&supplier)
		if err != nil {
			e.l.Error(err)
			continue
		}
		if businessPartner == businessPartner {
			return true
		}
	}
	return false
}
