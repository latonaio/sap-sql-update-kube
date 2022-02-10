# sap-sql-update-kube  

sap-sql-update-kube は、主にエッジコンピューティング環境において SAP API で取得したデータに SQL で保持された静的データとの差分がある場合に、SQLを更新するマイクロサービスです。  
本マイクロサービスには、SQLテーブルへのレコードの登録機能、更新機能が含まれます
本マイクロサービスは、SAP API の Runtime で取得されたデータを RabbitMQ のキューで受け取ります。  

## 動作環境  

* OS: Linux OS  
* CPU: ARM/AMD/Intel  
* Kubernetes  
* [sap-sandbox](https://github.com/latonaio/sap-sandbox) に含まれる API Runtimes

## クラウド環境での利用  
sap-sql-update-kube は、クラウド環境においても、利用可能なように設計されています。  

## SQLの登録更新 モジュール SQLBOILER の利用
sap-sql-update-kube は、SQLの登録更新モジュールとして、[SQLBOILER](https://github.com/volatiletech/sqlboiler)を利用しています。  

## 事前準備（SQLデータベースからのソースコードの生成）
sap-sql-update-kube のランタイム環境を利用するための事前準備として、  

```
/database/model_generate.sh
```
  
の実行をします。  
これにより、対象のSQLデータベースを読み込み、当該SQLデータベースの登録更新モジュールとして必要なソースコードが、/database/models/ 内に自動生成されます。  
（同一環境内での２回目以降の実行ではソースコードが上書きされます）  

model_generate.sh の 対象データベース接続時の情報は、/database/sqlboiler.toml に書かれる必要があります。    

sap-sql-update-kube には、ソースコードの初期値として、下記のSAP SQL テーブルを対象として生成されたソースコード群が、/database/models/ 内に含まれています。  

## SQLBoiler のバージョン
本レポジトリは、[SQLBoiler](https://github.com/volatiletech/sqlboiler) 4.8.6 のバージョンを使用しています。  
model_generate.sh のコマンドで、/database/models/ 内のソースコードを生成すると、当該ソースコード内の SQLBoiler のバージョンが最新になるのでご注意ください。      
バージョンを最新にした際は、併せて go modのバージョンを更新し、go mod tidy を行なってください。  

## 対象とする SAP SQL テーブル
本レポジトリには、以下の SAP SQLテーブル のレコードの登録・更新機能が含まれます。  

### Central Functions ###

#### Classification ####

* [ClassificationCharacteristic](https://github.com/latonaio/sap-classification-sql/blob/main/sap-classification-sql-characteristic-data.sql)    
* [ClassificationClass](https://github.com/latonaio/sap-classification-sql/blob/main/sap-classification-sql-class-data.sql)   
* [ClassificationClassDescription](https://github.com/latonaio/sap-classification-sql/blob/main/sap-classification-sql-class-description-data.sql) 

#### Characteristic ####

* [CharacteristicCharacteristic](https://github.com/latonaio/sap-characteristic-sql/blob/main/sap-characteristic-sql-characteristic-data.sql) 
* [CharacteristicValue](https://github.com/latonaio/sap-characteristic-sql/blob/main/sap-characteristic-sql-value-data.sql)   
* [CharacteristicCharcDescription](https://github.com/latonaio/sap-characteristic-sql/blob/main/sap-characteristic-sql-charc-description-data.sql)   
* [CharacteristicValueDescription](https://github.com/latonaio/sap-characteristic-sql/blob/main/sap-characteristic-sql-value-description-data.sql)   

#### Bank Master ####

* [BankMasterBank](https://github.com/latonaio/sap-bank-master-sql/blob/main/sap-bank-master-sql-bank-data.sql)
#### Business Partner ####

* [BusinessPartnerGeneral](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-general-data.sql)
* [BusinessPartnerRole](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-role-data.sql)
* [BusinessPartnerAddress](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-address-data.sql)
* [BusinessPartnerBank](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-bank-data.sql)

### Logistics  

#### Product Master ####

* [ProductMasterGeneral](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_basic_data.sql)
* [ProductMasterSalesPlant](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_sales_plant_data.sql)
* [ProductMasterPlant](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_plant_data.sql)
* [ProductMasterMRPArea](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_mrp_area_data.sql)
* [ProductMasterProcurement](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_procurement_data.sql)
* [ProductMasterAccounting](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_accounting_data.sql)
* [ProductMasterSalesOrganization](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_sales_organization_data.sql)
* [ProductMasterSalesTax](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_sales_tax_data.sql)
* [ProductMasterWorkScheduling](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_work_scheduling_data.sql)
* [ProductMasterQuality](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_quality_data.sql)
* [ProductMasterProductDesc](https://github.com/latonaio/sap-product-master-sql/blob/main/sap_product_master_sql_product_description_data.sql)

#### Batch Master Record ####

* [BatchMasterRecord](https://github.com/latonaio/sap-batch-master-record-sql/blob/main/sap-batch-master-record.sql)

### Inventory Management

#### Material Stock ####

* [MaterialStock](https://github.com/latonaio/sap-material-stock-sql/blob/main/sap-material-stock-sql.sql)

#### Reservation Document ####

* [ReservationDocumentHeader](https://github.com/latonaio/sap-reservation-document-sql/blob/main/sap-reservation-document-sql-header-data.sql)
* [ReservationDocumentItem](https://github.com/latonaio/sap-reservation-document-sql/blob/main/sap-reservation-document-sql-item-data.sql)

#### Inbound Delivery ####

* [InboundDeliveryHeader](https://github.com/latonaio/sap-inbound-delivery-sql/blob/main/sap-inbound-delivery-sql-header-data.sql)
* [InboundDeliveryItem](https://github.com/latonaio/sap-inbound-delivery-sql/blob/main/sap-inbound-delivery-sql-item-data.sql)
* [InboundDeliveryPartner](https://github.com/latonaio/sap-inbound-delivery-sql/blob/main/sap-inbound-delivery-sql-partner-data.sql)
* [InboundDeleveryAddress](https://github.com/latonaio/sap-inbound-delivery-sql/blob/main/sap-inbound-delivery-sql-address-data.sql)

#### Material Document ####

* [MaterialDocumentHeader](https://github.com/latonaio/sap-material-document-sql/blob/main/sap-material-document-sql-header-data.sql)
* [MaterialDocumentItem](https://github.com/latonaio/sap-material-document-sql/blob/main/sap-material-document-sql-item-data.sql)

#### Physical Inventory Document ####

* [PhysicalInventoryDocumentHeader](https://github.com/latonaio/sap-physical-inventory-document-sql/blob/main/sap-physical-inventory-document-sql-header-data.sql)
* [PhysicalInventoryDocumentItem](https://github.com/latonaio/sap-physical-inventory-document-sql/blob/main/sap-physical-inventory-document-sql-item-data.sql)

### Sales Management ###

#### Customer Master ####

* [BusinessPartnerCustomer](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-customer-data.sql)
* [BusinessPartnerCustomerPartnerFunction](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-customer-partner-function-data.sql)
* [BusinessPartnerCustomerSalesArea](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-customer-sales-area-data.sql)
* [BusinessPartnerCustomerCompany](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-customer-company-data.sql)

#### Customer Material ####

* [CustomerMaterial](https://github.com/latonaio/sap-customer-material-sql/blob/main/sap-customer-material-sql.sql)

#### Sales Pricing ####

* [SalesPricingConditionValidity](https://github.com/latonaio/sap-sales-pricing-sql/blob/main/sap-sales-pricing-sql-condition-validity-data.sql)
* [SalesPricingConditionRecord](https://github.com/latonaio/sap-sales-pricing-sql/blob/main/sap-sales-pricing-sql-condition-record-data.sql)

#### Sales Inquiry ####

* [SalesInquiryHeader](https://github.com/latonaio/sap-sales-inquiry-sql/blob/main/sap-sales-inquiry-sql-header-data.sql)
* [SalesInquiryHeaderPartner](https://github.com/latonaio/sap-sales-inquiry-sql/blob/main/sap-sales-inquiry-sql-header-partner-data.sql)
* [SalesInquiryItem](https://github.com/latonaio/sap-sales-inquiry-sql/blob/main/sap-sales-inquiry-sql-item-data.sql)
* [SalesInquiryItemPricingElement](https://github.com/latonaio/sap-sales-inquiry-sql/blob/main/sap-sales-inquiry-sql-item-pricing-element-data.sql)

#### Sales Quotation ####

* [SalesQuotationHeader](https://github.com/latonaio/sap-sales-quotation-sql/blob/main/sap-sales-quotation-sql-header-data.sql)  
* [SalesQuotationHeaderPartner](https://github.com/latonaio/sap-sales-quotation-sql/blob/main/sap-sales-quotation-sql-header-partner-data.sql)
* [SalesQuotationItem](https://https://github.com/latonaio/sap-sales-quotation-sql/blob/main/sap-sales-quotation-sql-item-data.sql)  
* [SalesQuotationItemPricingElement](https://github.com/latonaio/sap-sales-quotation-sql/blob/main/sap-sales-quotation-sql-item-pricing-element-data.sql)  

#### Sales Order ####

* [SalesOrderHeader](https://github.com/latonaio/sap-sales-order-sql/blob/main/sap-sales-order-sql-header-data.sql)
* [SalesOrderItem](https://github.com/latonaio/sap-sales-order-sql/blob/main/sap-sales-order-sql-item-data.sql)
* [SalesOrderHeaderPartner](https://github.com/latonaio/sap-sales-order-sql/blob/main/sap-sales-order-sql-header-partner-data.sql)
* [SalesOrderItemPricingElement](https://github.com/latonaio/sap-sales-order-sql/blob/main/sap-sales-order-sql-item-pricing-element-data.sql)
* [SalesOrderItemScheduleLine](https://github.com/latonaio/sap-sales-order-sql/blob/main/sap-sales-order-sql-item-schedule-line-data.sql)

#### Sales Contract ####

* [SalesContractHeader](https://github.com/latonaio/sap-sales-contract-sql/blob/main/sap-sales-contract-sql-header-data.sql)
* [SalesContractHeaderPartner](https://github.com/latonaio/sap-sales-contract-sql/blob/main/sap-sales-contract-sql-partner-data.sql)
* [SalesContractItem](https://github.com/latonaio/sap-sales-contract-sql/blob/main/sap-sales-contract-sql-item-data.sql)
* [SalesContractItemPricingElement](https://github.com/latonaio/sap-sales-contract-sql/blob/main/sap-sales-contract-sql-item-pricing-element-data.sql)

#### Sales Scheduling Agreement ####

* [SalesSchedulingAgreementHeader](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-header-data.sql)　　
* [SalesSchedulingAgreementPartner](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-partner-data.sql)　　
* [SalesSchedulingAgreementItem](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-item-data.sql)　　　　
* [SalesSchedulingAgreementItemDeliverySchedule](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-item-delivery-schedule-data.sql)　　　　
* [SalesSchedulingAgreementItemPricingElement](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-item-pricing-element-data.sql)  　　　
* [SalesSchedulingAgreementItemScheduleLine](https://github.com/latonaio/sap-sales-scheduling-agreement-sql/blob/main/sap-sales-scheduling-agreement-sql-item-schedule-line-data.sql)  　　

#### Supplier Invoice ####

* [SupplierInvoiceHeader](https://github.com/latonaio/sap-supplier-invoice-sql/blob/main/sap-supplier-invoice-sql-header-data.sql)  
* [SupplierInvoiceTax](https://github.com/latonaio/sap-supplier-invoice-sql/blob/main/sap-supplier-invoice-sql-tax-data.sql)  
* [SupplierInvoiceAccountAssignment](https://github.com/latonaio/sap-supplier-invoice-sql/blob/main/sap-supplier-invoice-sql-account-assignment-data.sql)  
* [SupplierInvoicePurchaseOrderReference](https://github.com/latonaio/sap-supplier-invoice-sql/blob/main/sap-supplier-invoice-sql-purchase-order-reference-data.sql)

#### Outbound Delivery ####

* [OutboundDeliveryHeader](https://github.com/latonaio/sap-outbound-delivery-sql/blob/main/sap-outbound-delivery-sql-header-data.sql)
* [OutboundDeliveryHeaderPartner](https://github.com/latonaio/sap-outbound-delivery-sql/blob/main/sap-outbound-delivery-sql-header-partner-data.sql)
* [OutboundDeliveryPartnerAddress](https://github.com/latonaio/sap-outbound-delivery-sql/blob/main/sap-outbound-delivery-sql-partner-address-data.sql)
* [OutboundDeliveryItem](https://github.com/latonaio/sap-outbound-delivery-sql/blob/main/sap-outbound-delivery-sql-item-data.sql)
* [OutboundDeliveryItemDocumentFlow](https://github.com/latonaio/sap-outbound-delivery-sql/blob/main/sap-outbound-delivery-sql-item-document-flow-data.sql)

#### Customer Return ####

* [CustomerReturnHeader](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-header-data.sql)  
* [CustomerReturnHeaderPartner](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-header-partner-data.sql)  
* [CustomerReturnItem](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-item-data.sql)  
* [CustomerReturnItemPricingElement](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-item-pricing-element-data.sql)  
* [CustomerReturnItemProcessStep](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-item-process-step-data.sql)  
* [CustomerReturnItemScheduleData](https://github.com/latonaio/sap-customer-return-sql/blob/main/sap-customer-return-sql-item-schedule-line-data.sql)

#### Billing Document ####

* [BillingDocumentHeader](https://github.com/latonaio/sap-billing-document-sql/blob/main/sap-billing-document-sql-header-data.sql)
* [BillingDocumentHeaderPartner](https://github.com/latonaio/sap-billing-document-sql/blob/main/sap-billing-document-sql-header-partner-data.sql)
* [BillingDocumentItem](https://github.com/latonaio/sap-billing-document-sql/blob/main/sap-billing-document-sql-item-data.sql)
* [BillingDocumentItemPartner](https://github.com/latonaio/sap-billing-document-sql/blob/main/sap-billing-document-sql-item-partner-data.sql)
* [BillingDocumentItemPricingElement](https://github.com/latonaio/sap-billing-document-sql/blob/main/sap-billing-document-sql-item-pricing-element-data.sql)

#### Credit Memo Request ####

* [CreditMemoRequestHeader](https://github.com/latonaio/sap-credit-memo-request-sql/blob/main/sap-credit-memo-request-sql-header-data.sql)
* [CreditMemoRequestHeaderPartner](https://github.com/latonaio/sap-credit-memo-request-sql/blob/main/sap-credit-memo-request-sql-header-partner-data.sql)
* [CreditMemoRequestItem](https://github.com/latonaio/sap-credit-memo-request-sql/blob/main/sap-credit-memo-request-sql-item-data.sql)
* [CreditMemoRequestItemPricingElement](https://github.com/latonaio/sap-credit-memo-request-sql/blob/main/sap-credit-memo-request-sql-item-pricing-element-data.sql)

#### Debit Memo Request ####

* [DebitMemoRequestHeader](https://github.com/latonaio/sap-debit-memo-request-sql/blob/main/sap-debit-memo-request-sql-header-data.sql)  
* [DebitMemoRequestHeaderPartner](https://github.com/latonaio/sap-debit-memo-request-sql/blob/main/sap-debit-memo-request-sql-header-partner-data.sql)    
* [DebitMemoRequestItem](https://github.com/latonaio/sap-debit-memo-request-sql/blob/main/sap-debit-memo-request-sql-item-data.sql)    
* [DebitMemoRequestItemPricingElement](https://github.com/latonaio/sap-debit-memo-request-sql/blob/main/sap-debit-memo-request-sql-item-pricing-element-data.sql)   

### Production Management

#### Bill Of Material ####

* [BillOfMaterialHeader](https://github.com/latonaio/sap-bill-of-material-sql/blob/main/sap-bill-of-material-sql-header-data.sql)
* [BillOfMaterialItem](https://github.com/latonaio/sap-bill-of-material-sql/blob/main/sap-bill-of-material-sql-item-data.sql)
  
#### Bill Of Material Where Used List 

* [BillOfMaterialWhereUsedList](https://github.com/latonaio/sap-bill-of-material-where-used-list-sql/blob/main/sap-bill-of-material-where-used-list.sql)
  
#### Production Routing ####

* [ProductionRoutingHeader](https://github.com/latonaio/sap-production-routing-sql/blob/main/sap-production-routing-sql-header-data.sql)
* [ProductionRoutingMaterialAssignment](https://github.com/latonaio/sap-production-routing-sql/blob/main/sap-production-routing-sql-material-assignment-data.sql)
* [ProductionRoutingSequence](https://github.com/latonaio/sap-production-routing-sql/blob/main/sap-production-routing-sql-sequence-data.sql)
* [ProductionRoutingOperation](https://github.com/latonaio/sap-production-routing-sql/blob/main/sap-production-routing-sql-operation-data.sql)
* [ProductionRoutingComponentAllocation](https://github.com/latonaio/sap-production-routing-sql/blob/main/sap-production-routing-sql-component-allocation-data.sql)

#### Planned Independent Requirement ####

* [PlannedIndependentRequirementHeader](https://github.com/latonaio/sap-planned-independent-requirement-sql/blob/main/sap-planned-independent-requirement-sql-header-data.sql)
* [PlannedIndependentRequirementItem](https://github.com/latonaio/sap-planned-independent-requirement-sql/blob/main/sap-planned-independent-requirement-sql-item-data.sql)

#### Production Order ####

* [ProductionOrderGeneral](https://github.com/latonaio/sap-production-order-sql/blob/main/sap-production-order-sql-general-data.sql)
* [ProductionOrderComponent](https://github.com/latonaio/sap-production-order-sql/blob/main/sap-production-order-sql-component-data.sql)
* [ProductionOrderItem](https://github.com/latonaio/sap-production-order-sql/blob/main/sap-production-order-sql-item-data.sql)
* [ProductionOrderOperation](https://github.com/latonaio/sap-production-order-sql/blob/main/sap-production-order-sql-operation-data.sql)
* [ProductionOrderStatus](https://github.com/latonaio/sap-production-order-sql/blob/main/sap-production-order-sql-status-data.sql)

#### Production Order Confirmation ####
* [ProductionOrderConfirmationConfirmation](https://github.com/latonaio/sap-production-order-confirmation-sql/blob/main/sap-production-order-confirmation-sql-confirmation-data.sql)
* [ProductionOrderConfirmationMaterialMovements](https://github.com/latonaio/sap-production-order-confirmation-sql/blob/main/sap-production-order-confirmation-sql-material-movements-data.sql)
* [ProductionOrderConfirmationBatchCharacteristic](https://github.com/latonaio/sap-production-order-confirmation-sql/blob/main/sap-production-order-confirmation-sql-batch-characteristic-data.sql)

### Process Management

#### Process Order Confirmation ####

* [ProcessOrderConfirmationConfirmation](https://github.com/latonaio/sap-process-order-confirmation-sql/blob/main/sap-process-order-confirmation-sql-confirmation-data.sql)
* [ProcessOrderConfirmationMaterialMovements](https://github.com/latonaio/sap-process-order-confirmation-sql/blob/main/sap-process-order-confirmation-sql-material-movements-data.sql)
* [ProcessOrderConfirmationBatchCharacteristic](https://github.com/latonaio/sap-process-order-confirmation-sql/blob/main/sap-process-order-confirmation-sql-batch-characteristics-data.sql)
### Procurement Management

#### Supplier Master ####

* [BusinessPartnerSupplier](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-supplier-data.sql)
* [BusinessPartnerSupplierPartnerFunction](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-supplier-partner-function-data.sql)
* [BusinessPartnerSupplierCompany](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-supplier-company-data.sql)
* [BusinessPartnerSupplierPurchasingOrganization](https://github.com/latonaio/sap-business-partner-sql/blob/main/sap-business-partner-sql-supplier-purchasing-organization-data.sql)

#### Purchasing Source List

* [PurchasingSourceList](https://github.com/latonaio/sap-purchasing-source-list-sql/blob/main/sap-purchasing-source-list.sql)

#### Purchasing Info Record ####

* [PurchasingInfoRecordPurchasingOrganizationPlant](https://github.com/latonaio/sap-purchasing-info-record-sql/blob/main/sap-purchasing-info-record-purchasing-organization-plant-data.sql)  
* [PurchasingInfoRecordGeneral](https://github.com/latonaio/sap-purchasing-info-record-sql/blob/main/sap-purchasing-info-record-sql-general-data.sql)  
* [PurchasingInfoRecordPricingCondition](https://github.com/latonaio/sap-purchasing-info-record-sql/blob/main/sap-purchasing-info-record-sql-pricing-condition-data.sql)  
* [PurchasingInfoRecordPricingConditionValidity](https://github.com/latonaio/sap-purchasing-info-record-sql/blob/main/sap-purchasing-info-record-sql-pricing-condition-validity-data.sql)


#### Purchase Requisition ####

* [PurchaseRequisitionHeader](https://github.com/latonaio/sap-purchase-requisition-sql/blob/main/sap-purchase-requisition-sql-header-data.sql)
* [PurchaseRequisitionItem](https://github.com/latonaio/sap-purchase-requisition-sql/blob/main/sap-purchase-requisition-sql-item-data.sql)
* [PurchaseRequisitionItemDeliveryAddress](https://github.com/latonaio/sap-purchase-requisition-sql/blob/main/sap-purchase-requisition-sql-item-delivery-address-data.sql)
* [PurchaseRequisitionItemAccountAssignment](https://github.com/latonaio/sap-purchase-requisition-sql/blob/main/sap-purchase-requisition-sql-item-account-assignment-data.sql)

#### Purchase Order ####

* [PurchaseOrderHeaderWithAddress](https://github.com/latonaio/sap-purchase-order-sql/blob/main/sap-purchase-order-sql-header-data-with-address.sql)
* [PurchaseOrderItemWithAddress](https://github.com/latonaio/sap-purchase-order-sql/blob/main/sap-purchase-order-sql-item-data-with-address.sql)
* [PurchaseOrderItemScheduleLine](https://github.com/latonaio/sap-purchase-order-sql/blob/main/sap-purchase-order-sql-item-schedule-line-data.sql)
* [PurchaseOrderItemPricingElement](https://github.com/latonaio/sap-purchase-order-sql/blob/main/sap-purchase-order-sql-item-pricing-element-data.sql)
* [PurchaseOrderItemAccount](https://github.com/latonaio/sap-purchase-order-sql/blob/main/sap-purchase-order-sql-item-account-data.sql)

#### Purchase Contract ####

* [PurchaseContractHeader](https://github.com/latonaio/sap-purchase-contract-sql/blob/main/sap-purchase-contract-sql-header-data.sql)
* [PurchaseContractItem](https://github.com/latonaio/sap-purchase-contract-sql/blob/main/sap-purchase-contract-sql-item-data.sql)
* [PurchaseContractAddress](https://github.com/latonaio/sap-purchase-contract-sql/blob/main/sap-purchase-contract-sql-address-data.sql)
* [PurchaseComtractItemCondition](https://github.com/latonaio/sap-purchase-contract-sql/blob/main/sap-purchase-contract-sql-item-condition-data.sql)

### Plant Maintenance

#### Functional Location ####

* [FunctionalLocationHeader](https://github.com/latonaio/sap-functional-location-sql/blob/main/sap-functional-location-sql-header-data.sql)

#### Equipment Master ####

* [Equipment](https://github.com/latonaio/sap-equipment-master-sql/blob/main/sap-equipment-master-sql-equpment-data.sql)
* [EquipmentPartner](https://github.com/latonaio/sap-equipment-master-sql/blob/main/sap-equipment-master-sql-business-partner-data.sql)

#### Maintenance Bill Of Material ####

* [MaintenanceBillOfMaterialHeader](https://github.com/latonaio/sap-maintenance-bill-of-material-sql/blob/main/sap-maintenance-bill-of-material-sql-header-data.sql)
* [MaintenanceBillOfMaterialItem](https://github.com/latonaio/sap-maintenance-bill-of-material-sql/blob/main/sap-maintenance-bill-of-material-sql-item-data.sql)

#### Maintenance Plan ####

* [MaintenancePlanHeader](https://github.com/latonaio/sap-maintenance-plan-sql/blob/main/sap-maintenance-plan-sql-header-data.sql)
* [MaintenancePlanStrategyCycle](https://github.com/latonaio/sap-maintenance-plan-sql/blob/main/sap-maintenance-plan-sql-strategy-cycle-data.sql)
* [MaintenancePlanItem](https://github.com/latonaio/sap-maintenance-plan-sql/blob/main/sap-maintenance-plan-sql-item-data.sql)
* [MaintenancePlanCallObjects](https://github.com/latonaio/sap-maintenance-plan-sql/blob/main/sap-maintenance-plan-sql-call-objects-data.sql)

#### Maintenance Item ####

* [MaintenanceItemItem](https://github.com/latonaio/sap-maintenance-item-sql/blob/main/sap-maintenance-item-sql-item-data.sql)
* [MaintenanceItemCallObjects](https://github.com/latonaio/sap-maintenance-item-sql/blob/main/sap-maintenance-item-sql-call-objects-data.sql)

#### Maintenance Notification ####

* [MaintenanceNotificationHeader](https://github.com/latonaio/sap-maintenance-notification-sql/blob/main/sap-maintenance-notification-sql-header-data.sql)
* [MaintenanceNotificationItem](https://github.com/latonaio/sap-maintenance-notification-sql/blob/main/sap-maintenance-notification-sql-item-data.sql)
* [MaintenanceNotificationItemCause](https://github.com/latonaio/sap-maintenance-notification-sql/blob/main/sap-maintenance-notification-sql-item-cause-data.sql)

#### Maintenance Order ####

* [MaintenanceOrderHeader](https://github.com/latonaio/sap-maintenance-order-sql/blob/main/sap-maintenance-order-sql-header-data.sql)
* [MaintenanceOrderObjectListItem](https://github.com/latonaio/sap-maintenance-order-sql/blob/main/sap-maintenance-order-sql-object-list-item-data.sql)
* [MaintenanceOrderOperation](https://github.com/latonaio/sap-maintenance-order-sql/blob/main/sap-maintenance-order-sql-operation-data.sql)
* [MaintenanceOrderOperationComponent](https://github.com/latonaio/sap-maintenance-order-sql/blob/main/sap-maintenance-order-sql-operation-component-data.sql)

#### Maintenance Order Confirmation ####

* [MaintenanceOrderConfirmationHeader](https://github.com/latonaio/sap-maintenance-order-confirmation-sql/blob/main/sap-maintenance-order-confirmation-header-data.sql)


#### Defect ####

* [DefectHeader](https://github.com/latonaio/sap-defect-sql/blob/main/sap-defect-sql-header-data.sql)

#### Maintenance Task List ####

* [MaintenanceTaskListHeader](https://github.com/latonaio/sap-maintenance-task-list-sql/blob/main/sap-maintenance-task-list-sql-header-data.sql)
* [MaintenanceTaskListStrategyPackage](https://github.com/latonaio/sap-maintenance-task-list-sql/blob/main/sap-maintenance-task-list-sql-strategy-package-data.sql)
* [MaintenanceTaskListOperation](https://github.com/latonaio/sap-maintenance-task-list-sql/blob/main/sap-maintenance-task-list-sql-operation-data.sql)
* [MaintenanceTaskListOperationMaterial](https://github.com/latonaio/sap-maintenance-task-list-sql/blob/main/sap-maintenance-task-list-sql-operation-material-data.sql)

### Customer Service  

#### Service Order ####

* [ServiceOrderHeder](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-header-data.sql)
* [ServiceOrderItem](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-item-data.sql)
* [ServiceOrderConfirmation](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-confirmation-data.sql)
* [ServiceOrderDefect](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-defect-data.sql)
* [ServiceOrderPersonResponsible](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-person-responsible-data.sql)
* [ServiceOrderReferenceObject](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-reference-object-data.sql)
* [ServiceOrderItemPricingElement](https://github.com/latonaio/sap-service-order-sql/blob/main/sap-service-order-sql-item-pricing-element-data.sql)

#### Service Confirmation ####

* [ServiceConfirmationHeder](https://github.com/latonaio/sap-service-confirmation-sql/blob/main/sap-service-confirmation-sql-header-data.sql)
* [ServiceConfirmationItem](https://github.com/latonaio/sap-service-confirmation-sql/blob/main/sap-service-confirmation-sql-item-data.sql)
* [ServiceConfirmationPersonResponsible](https://github.com/latonaio/sap-service-confirmation-sql/blob/main/sap-service-confirmation-sql-person-responsible-data.sql)
* [ServiceConfirmationReferenceObject](https://github.com/latonaio/sap-service-confirmation-sql/blob/main/sap-service-confirmation-sql-reference-object-data.sql)
* [ServiceConfirmationItemPricingElement](https://github.com/latonaio/sap-service-confirmation-sql/blob/main/sap-service-confirmation-sql-item-pricing-element-data.sql)

## RabbitMQ からの JSON Input

sap-sql-update-kube は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sap-sql-update-kube は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ の マスタサーバ環境

sap-sql-update-kube が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sap-sql-update-kube は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sap-sql-update-kube の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```

## Input JSON の例
例えば、[sap-api-integrations-product-master-reads-rmq-kube](https://github.com/latonaio/sap-api-integrations-production-order-reads-rmq-kube)で出力されたデータは、以下の形式でインプットされます。  
```
{
    "cursor": "/Users/latona2/bitbucket/sap-api-integrations-product-master-reads-rmq-kube/SAP_API_Caller/caller.go#L93",
    "function": "sap-api-integrations-product-master-reads-rmq-kube/SAP_API_Caller.(*SAPAPICaller).General",
    "level": "INFO",
    "message": [
        {
            "Material": "A001",
            "Product_desc": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description",
            "BaseUnit": "AU",
            "ValidityStartDate": "",
            "ProductGroup": "A001",
            "Division": "00",
            "GrossWeight": "0.000",
            "WeightUnit": "KG",
            "SizeOrDimensionText": "",
            "ProductStandardID": ""
        }
    ],
    "time": "2021-12-22T10:19:54.310555+09:00"
}
```
