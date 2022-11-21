# data-platform-api-business-partner-exconf-supplier-rmq-kube
data-platform-api-business-partner-exconf-supplier-rmq-kube は、データ連携基盤において、API でビジネスパートナ仕入先の存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、ビジネスパートナ仕入先の存在確認が行われます。

* data-platform-business-partner-sql-supplier-data.sql（データ連携基盤 ビジネスパートナ - 仕入先データ）

## caller.go による存在性確認
Input で取得されたファイルに基づいて、caller.go で、 API がコールされます。
caller.go の 以下の箇所が、指定された API をコールするソースコードです。

```
func (e *ExistenceConf) Conf(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.BusinessPartnerSupplier {
	businessPartner := *input.BusinessPartnerSupplier.BusinessPartner
	supplier := *input.BusinessPartnerSupplier.supplier
	notKeyExistence := make([]int, 0, 1)
	KeyExistence := make([]int, 0, 1)

	existData := &dpfm_api_output_formatter.BusinessPartnerSupplier{
		BusinessPartner: businessPartner,
		Supplier:        supplier,
		ExistenceConf:   false,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if !e.confBusinessPartnerSupplier(businessPartner) {
			notKeyExistence = append(notKeyExistence, businessPartner)
			return
		}
		KeyExistence = append(KeyExistence, businessPartner)
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
```

## Input
data-platform-api-business-partner-exconf-Supplier-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "ORDERS",
	"BusinessPartnerSupplier": {
		"BusinessPartner": 101,
		"Supplier": 201
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
}
```

## Output
data-platform-api-business-partner-exconf-supplier-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。ビジネスパートナ仕入先の対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"service_label": "ORDERS",
	"BusinessPartnerSupplier": {
		"BusinessPartner": 101,
		"Supplier": 201,
		"ExistenceConf": true
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": [
		"All"
	],
	"order_id": null,
	"deleted": false
}
```


