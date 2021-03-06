// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

var TableNames = struct {
	SapBankMasterBankData                                 string
	SapBatchMasterRecord                                  string
	SapBillOfMaterialHeaderData                           string
	SapBillOfMaterialItemData                             string
	SapBillOfMaterialWhereUsedListData                    string
	SapBillingDocumentHeaderData                          string
	SapBillingDocumentHeaderPartnerData                   string
	SapBillingDocumentItemData                            string
	SapBillingDocumentItemPartnerData                     string
	SapBillingDocumentItemPricingElementData              string
	SapBusinessPartnerAddressData                         string
	SapBusinessPartnerBankData                            string
	SapBusinessPartnerCustomerCompanyData                 string
	SapBusinessPartnerCustomerData                        string
	SapBusinessPartnerCustomerPartnerFunctionData         string
	SapBusinessPartnerCustomerSalesAreaData               string
	SapBusinessPartnerGeneralData                         string
	SapBusinessPartnerRoleData                            string
	SapBusinessPartnerSupplierCompanyData                 string
	SapBusinessPartnerSupplierData                        string
	SapBusinessPartnerSupplierPartnerFunctionData         string
	SapBusinessPartnerSupplierPurchasingOrganizationData  string
	SapCharacteristicCharacteristicData                   string
	SapCharacteristicCharcDescriptionData                 string
	SapCharacteristicValueData                            string
	SapCharacteristicValueDescriptionData                 string
	SapClassificationCharacteristicData                   string
	SapClassificationClassData                            string
	SapClassificationClassDescriptionData                 string
	SapCreditManagementMasterBusinessPartnerData          string
	SapCreditManagementMasterCreditAccountData            string
	SapCreditMemoRequestHeaderData                        string
	SapCreditMemoRequestHeaderPartnerData                 string
	SapCreditMemoRequestItemData                          string
	SapCreditMemoRequestItemPricingElementData            string
	SapCustomerMaterialData                               string
	SapCustomerReturnHeaderData                           string
	SapCustomerReturnHeaderPartnerData                    string
	SapCustomerReturnItemData                             string
	SapCustomerReturnItemPricingElementData               string
	SapCustomerReturnItemProcessStepData                  string
	SapCustomerReturnItemScheduleLineData                 string
	SapDebitMemoRequestHeaderData                         string
	SapDebitMemoRequestHeaderPartnerData                  string
	SapDebitMemoRequestItemData                           string
	SapDebitMemoRequestItemPricingElementData             string
	SapDefectHeaderData                                   string
	SapEquipmentMasterBusinessPartnerData                 string
	SapEquipmentMasterEquipmentData                       string
	SapFunctionalLocationHeaderData                       string
	SapInboundDeliveryAddressData                         string
	SapInboundDeliveryHeaderData                          string
	SapInboundDeliveryItemData                            string
	SapInboundDeliveryPartnerData                         string
	SapMaintenanceBillOfMaterialHeaderData                string
	SapMaintenanceBillOfMaterialItemData                  string
	SapMaintenanceItemCallObjectsData                     string
	SapMaintenanceItemItemData                            string
	SapMaintenanceNotificationHeaderData                  string
	SapMaintenanceNotificationItemCauseData               string
	SapMaintenanceNotificationItemData                    string
	SapMaintenanceOrderConfirmationHeaderData             string
	SapMaintenanceOrderHeaderData                         string
	SapMaintenanceOrderObjectListItemData                 string
	SapMaintenanceOrderOperationComponentData             string
	SapMaintenanceOrderOperationData                      string
	SapMaintenancePlanCallObjectsData                     string
	SapMaintenancePlanHeaderData                          string
	SapMaintenancePlanItemData                            string
	SapMaintenancePlanStrategyCycleData                   string
	SapMaintenanceTaskListHeaderData                      string
	SapMaintenanceTaskListOperationData                   string
	SapMaintenanceTaskListOperationMaterialData           string
	SapMaintenanceTaskListStrategyPackageData             string
	SapMaterialDocumentHeaderData                         string
	SapMaterialDocumentItemData                           string
	SapMaterialStockData                                  string
	SapMeasurementDocumentHeaderData                      string
	SapOutboundDeliveryHeaderData                         string
	SapOutboundDeliveryHeaderPartnerData                  string
	SapOutboundDeliveryItemData                           string
	SapOutboundDeliveryItemDocumentFlowData               string
	SapOutboundDeliveryPartnerAddressData                 string
	SapPhysicalInventoryDocumentHeaderData                string
	SapPhysicalInventoryDocumentItemData                  string
	SapPlannedIndependentRequirementHeaderData            string
	SapPlannedIndependentRequirementItemData              string
	SapPlannedOrderComponentData                          string
	SapPlannedOrderHeaderData                             string
	SapProcessOrderConfirmationBatchCharacteristicData    string
	SapProcessOrderConfirmationConfirmationData           string
	SapProcessOrderConfirmationMaterialMovementsData      string
	SapProductMasterAccountingData                        string
	SapProductMasterClassClassDetailsData                 string
	SapProductMasterClassProductCharacteristicData        string
	SapProductMasterClassProductClassData                 string
	SapProductMasterClassProductGeneralData               string
	SapProductMasterGeneralData                           string
	SapProductMasterMRPAreaData                           string
	SapProductMasterPlantData                             string
	SapProductMasterProcurementData                       string
	SapProductMasterProductDescriptionData                string
	SapProductMasterQualityData                           string
	SapProductMasterSalesOrganizationData                 string
	SapProductMasterSalesPlantData                        string
	SapProductMasterSalesTaxData                          string
	SapProductMasterWorkSchedulingData                    string
	SapProductionOrderComponentData                       string
	SapProductionOrderConfirmationBatchCharacteristicData string
	SapProductionOrderConfirmationConfirmationData        string
	SapProductionOrderConfirmationMaterialMovementsData   string
	SapProductionOrderGeneralData                         string
	SapProductionOrderItemData                            string
	SapProductionOrderOperationData                       string
	SapProductionOrderStatusData                          string
	SapProductionRoutingComponentAllocationData           string
	SapProductionRoutingHeaderData                        string
	SapProductionRoutingMaterialAssignmentData            string
	SapProductionRoutingOperationData                     string
	SapProductionRoutingSequenceData                      string
	SapPurchaseContractAddressData                        string
	SapPurchaseContractHeaderData                         string
	SapPurchaseContractItemConditionData                  string
	SapPurchaseContractItemData                           string
	SapPurchaseOrderHeaderDataWithAddress                 string
	SapPurchaseOrderItemAccountData                       string
	SapPurchaseOrderItemDataWithAddress                   string
	SapPurchaseOrderItemPricingElementData                string
	SapPurchaseOrderItemScheduleLineData                  string
	SapPurchaseRequisitionHeaderData                      string
	SapPurchaseRequisitionItemAccountAssignmentData       string
	SapPurchaseRequisitionItemData                        string
	SapPurchaseRequisitionItemDeliveryAddressData         string
	SapPurchasingInfoRecordGeneralData                    string
	SapPurchasingInfoRecordOrganizationPlantData          string
	SapPurchasingInfoRecordPricingConditionData           string
	SapPurchasingInfoRecordPricingConditionValidityData   string
	SapPurchasingSourceListData                           string
	SapReservationDocumentHeaderData                      string
	SapReservationDocumentItemData                        string
	SapSalesContractHeaderData                            string
	SapSalesContractHeaderPartnerData                     string
	SapSalesContractItemData                              string
	SapSalesContractItemPricingElementData                string
	SapSalesInquiryHeaderData                             string
	SapSalesInquiryHeaderPartnerData                      string
	SapSalesInquiryItemData                               string
	SapSalesInquiryItemPricingElementData                 string
	SapSalesOrderHeaderData                               string
	SapSalesOrderHeaderPartnerData                        string
	SapSalesOrderItemData                                 string
	SapSalesOrderItemPricingElementData                   string
	SapSalesOrderItemScheduleLineData                     string
	SapSalesPricingConditionRecordData                    string
	SapSalesPricingConditionValidityData                  string
	SapSalesQuotationHeaderData                           string
	SapSalesQuotationHeaderPartnerData                    string
	SapSalesQuotationItemData                             string
	SapSalesQuotationItemPricingElementData               string
	SapSalesSchedulingAgreementHeaderData                 string
	SapSalesSchedulingAgreementItemData                   string
	SapSalesSchedulingAgreementItemDeliveryScheduleData   string
	SapSalesSchedulingAgreementItemPricingElementData     string
	SapSalesSchedulingAgreementItemScheduleLineData       string
	SapSalesSchedulingAgreementPartnerData                string
	SapServiceConfirmationHeaderData                      string
	SapServiceConfirmationItemData                        string
	SapServiceConfirmationItemPricingElementData          string
	SapServiceConfirmationPersonResponsibleData           string
	SapServiceConfirmationReferenceObjectData             string
	SapServiceOrderConfirmationData                       string
	SapServiceOrderDefectData                             string
	SapServiceOrderHeaderData                             string
	SapServiceOrderItemData                               string
	SapServiceOrderItemPricingElementData                 string
	SapServiceOrderPersonResponsibleData                  string
	SapServiceOrderReferenceObjectData                    string
	SapSupplierInvoiceAccountAssignmentData               string
	SapSupplierInvoiceHeaderData                          string
	SapSupplierInvoicePurchaseOrderReferenceData          string
	SapSupplierInvoiceTaxData                             string
	SapWarehouseResourceHeaderData                        string
	SapWorkCenterData                                     string
}{
	SapBankMasterBankData:                                 "sap_bank_master_bank_data",
	SapBatchMasterRecord:                                  "sap_batch_master_record",
	SapBillOfMaterialHeaderData:                           "sap_bill_of_material_header_data",
	SapBillOfMaterialItemData:                             "sap_bill_of_material_item_data",
	SapBillOfMaterialWhereUsedListData:                    "sap_bill_of_material_where_used_list_data",
	SapBillingDocumentHeaderData:                          "sap_billing_document_header_data",
	SapBillingDocumentHeaderPartnerData:                   "sap_billing_document_header_partner_data",
	SapBillingDocumentItemData:                            "sap_billing_document_item_data",
	SapBillingDocumentItemPartnerData:                     "sap_billing_document_item_partner_data",
	SapBillingDocumentItemPricingElementData:              "sap_billing_document_item_pricing_element_data",
	SapBusinessPartnerAddressData:                         "sap_business_partner_address_data",
	SapBusinessPartnerBankData:                            "sap_business_partner_bank_data",
	SapBusinessPartnerCustomerCompanyData:                 "sap_business_partner_customer_company_data",
	SapBusinessPartnerCustomerData:                        "sap_business_partner_customer_data",
	SapBusinessPartnerCustomerPartnerFunctionData:         "sap_business_partner_customer_partner_function_data",
	SapBusinessPartnerCustomerSalesAreaData:               "sap_business_partner_customer_sales_area_data",
	SapBusinessPartnerGeneralData:                         "sap_business_partner_general_data",
	SapBusinessPartnerRoleData:                            "sap_business_partner_role_data",
	SapBusinessPartnerSupplierCompanyData:                 "sap_business_partner_supplier_company_data",
	SapBusinessPartnerSupplierData:                        "sap_business_partner_supplier_data",
	SapBusinessPartnerSupplierPartnerFunctionData:         "sap_business_partner_supplier_partner_function_data",
	SapBusinessPartnerSupplierPurchasingOrganizationData:  "sap_business_partner_supplier_purchasing_organization_data",
	SapCharacteristicCharacteristicData:                   "sap_characteristic_characteristic_data",
	SapCharacteristicCharcDescriptionData:                 "sap_characteristic_charc_description_data",
	SapCharacteristicValueData:                            "sap_characteristic_value_data",
	SapCharacteristicValueDescriptionData:                 "sap_characteristic_value_description_data",
	SapClassificationCharacteristicData:                   "sap_classification_characteristic_data",
	SapClassificationClassData:                            "sap_classification_class_data",
	SapClassificationClassDescriptionData:                 "sap_classification_class_description_data",
	SapCreditManagementMasterBusinessPartnerData:          "sap_credit_management_master_business_partner_data",
	SapCreditManagementMasterCreditAccountData:            "sap_credit_management_master_credit_account_data",
	SapCreditMemoRequestHeaderData:                        "sap_credit_memo_request_header_data",
	SapCreditMemoRequestHeaderPartnerData:                 "sap_credit_memo_request_header_partner_data",
	SapCreditMemoRequestItemData:                          "sap_credit_memo_request_item_data",
	SapCreditMemoRequestItemPricingElementData:            "sap_credit_memo_request_item_pricing_element_data",
	SapCustomerMaterialData:                               "sap_customer_material_data",
	SapCustomerReturnHeaderData:                           "sap_customer_return_header_data",
	SapCustomerReturnHeaderPartnerData:                    "sap_customer_return_header_partner_data",
	SapCustomerReturnItemData:                             "sap_customer_return_item_data",
	SapCustomerReturnItemPricingElementData:               "sap_customer_return_item_pricing_element_data",
	SapCustomerReturnItemProcessStepData:                  "sap_customer_return_item_process_step_data",
	SapCustomerReturnItemScheduleLineData:                 "sap_customer_return_item_schedule_line_data",
	SapDebitMemoRequestHeaderData:                         "sap_debit_memo_request_header_data",
	SapDebitMemoRequestHeaderPartnerData:                  "sap_debit_memo_request_header_partner_data",
	SapDebitMemoRequestItemData:                           "sap_debit_memo_request_item_data",
	SapDebitMemoRequestItemPricingElementData:             "sap_debit_memo_request_item_pricing_element_data",
	SapDefectHeaderData:                                   "sap_defect_header_data",
	SapEquipmentMasterBusinessPartnerData:                 "sap_equipment_master_business_partner_data",
	SapEquipmentMasterEquipmentData:                       "sap_equipment_master_equipment_data",
	SapFunctionalLocationHeaderData:                       "sap_functional_location_header_data",
	SapInboundDeliveryAddressData:                         "sap_inbound_delivery_address_data",
	SapInboundDeliveryHeaderData:                          "sap_inbound_delivery_header_data",
	SapInboundDeliveryItemData:                            "sap_inbound_delivery_item_data",
	SapInboundDeliveryPartnerData:                         "sap_inbound_delivery_partner_data",
	SapMaintenanceBillOfMaterialHeaderData:                "sap_maintenance_bill_of_material_header_data",
	SapMaintenanceBillOfMaterialItemData:                  "sap_maintenance_bill_of_material_item_data",
	SapMaintenanceItemCallObjectsData:                     "sap_maintenance_item_call_objects_data",
	SapMaintenanceItemItemData:                            "sap_maintenance_item_item_data",
	SapMaintenanceNotificationHeaderData:                  "sap_maintenance_notification_header_data",
	SapMaintenanceNotificationItemCauseData:               "sap_maintenance_notification_item_cause_data",
	SapMaintenanceNotificationItemData:                    "sap_maintenance_notification_item_data",
	SapMaintenanceOrderConfirmationHeaderData:             "sap_maintenance_order_confirmation_header_data",
	SapMaintenanceOrderHeaderData:                         "sap_maintenance_order_header_data",
	SapMaintenanceOrderObjectListItemData:                 "sap_maintenance_order_object_list_item_data",
	SapMaintenanceOrderOperationComponentData:             "sap_maintenance_order_operation_component_data",
	SapMaintenanceOrderOperationData:                      "sap_maintenance_order_operation_data",
	SapMaintenancePlanCallObjectsData:                     "sap_maintenance_plan_call_objects_data",
	SapMaintenancePlanHeaderData:                          "sap_maintenance_plan_header_data",
	SapMaintenancePlanItemData:                            "sap_maintenance_plan_item_data",
	SapMaintenancePlanStrategyCycleData:                   "sap_maintenance_plan_strategy_cycle_data",
	SapMaintenanceTaskListHeaderData:                      "sap_maintenance_task_list_header_data",
	SapMaintenanceTaskListOperationData:                   "sap_maintenance_task_list_operation_data",
	SapMaintenanceTaskListOperationMaterialData:           "sap_maintenance_task_list_operation_material_data",
	SapMaintenanceTaskListStrategyPackageData:             "sap_maintenance_task_list_strategy_package_data",
	SapMaterialDocumentHeaderData:                         "sap_material_document_header_data",
	SapMaterialDocumentItemData:                           "sap_material_document_item_data",
	SapMaterialStockData:                                  "sap_material_stock_data",
	SapMeasurementDocumentHeaderData:                      "sap_measurement_document_header_data",
	SapOutboundDeliveryHeaderData:                         "sap_outbound_delivery_header_data",
	SapOutboundDeliveryHeaderPartnerData:                  "sap_outbound_delivery_header_partner_data",
	SapOutboundDeliveryItemData:                           "sap_outbound_delivery_item_data",
	SapOutboundDeliveryItemDocumentFlowData:               "sap_outbound_delivery_item_document_flow_data",
	SapOutboundDeliveryPartnerAddressData:                 "sap_outbound_delivery_partner_address_data",
	SapPhysicalInventoryDocumentHeaderData:                "sap_physical_inventory_document_header_data",
	SapPhysicalInventoryDocumentItemData:                  "sap_physical_inventory_document_item_data",
	SapPlannedIndependentRequirementHeaderData:            "sap_planned_independent_requirement_header_data",
	SapPlannedIndependentRequirementItemData:              "sap_planned_independent_requirement_item_data",
	SapPlannedOrderComponentData:                          "sap_planned_order_component_data",
	SapPlannedOrderHeaderData:                             "sap_planned_order_header_data",
	SapProcessOrderConfirmationBatchCharacteristicData:    "sap_process_order_confirmation_batch_characteristic_data",
	SapProcessOrderConfirmationConfirmationData:           "sap_process_order_confirmation_confirmation_data",
	SapProcessOrderConfirmationMaterialMovementsData:      "sap_process_order_confirmation_material_movements_data",
	SapProductMasterAccountingData:                        "sap_product_master_accounting_data",
	SapProductMasterClassClassDetailsData:                 "sap_product_master_class_class_details_data",
	SapProductMasterClassProductCharacteristicData:        "sap_product_master_class_product_characteristic_data",
	SapProductMasterClassProductClassData:                 "sap_product_master_class_product_class_data",
	SapProductMasterClassProductGeneralData:               "sap_product_master_class_product_general_data",
	SapProductMasterGeneralData:                           "sap_product_master_general_data",
	SapProductMasterMRPAreaData:                           "sap_product_master_mrp_area_data",
	SapProductMasterPlantData:                             "sap_product_master_plant_data",
	SapProductMasterProcurementData:                       "sap_product_master_procurement_data",
	SapProductMasterProductDescriptionData:                "sap_product_master_product_description_data",
	SapProductMasterQualityData:                           "sap_product_master_quality_data",
	SapProductMasterSalesOrganizationData:                 "sap_product_master_sales_organization_data",
	SapProductMasterSalesPlantData:                        "sap_product_master_sales_plant_data",
	SapProductMasterSalesTaxData:                          "sap_product_master_sales_tax_data",
	SapProductMasterWorkSchedulingData:                    "sap_product_master_work_scheduling_data",
	SapProductionOrderComponentData:                       "sap_production_order_component_data",
	SapProductionOrderConfirmationBatchCharacteristicData: "sap_production_order_confirmation_batch_characteristic_data",
	SapProductionOrderConfirmationConfirmationData:        "sap_production_order_confirmation_confirmation_data",
	SapProductionOrderConfirmationMaterialMovementsData:   "sap_production_order_confirmation_material_movements_data",
	SapProductionOrderGeneralData:                         "sap_production_order_general_data",
	SapProductionOrderItemData:                            "sap_production_order_item_data",
	SapProductionOrderOperationData:                       "sap_production_order_operation_data",
	SapProductionOrderStatusData:                          "sap_production_order_status_data",
	SapProductionRoutingComponentAllocationData:           "sap_production_routing_component_allocation_data",
	SapProductionRoutingHeaderData:                        "sap_production_routing_header_data",
	SapProductionRoutingMaterialAssignmentData:            "sap_production_routing_material_assignment_data",
	SapProductionRoutingOperationData:                     "sap_production_routing_operation_data",
	SapProductionRoutingSequenceData:                      "sap_production_routing_sequence_data",
	SapPurchaseContractAddressData:                        "sap_purchase_contract_address_data",
	SapPurchaseContractHeaderData:                         "sap_purchase_contract_header_data",
	SapPurchaseContractItemConditionData:                  "sap_purchase_contract_item_condition_data",
	SapPurchaseContractItemData:                           "sap_purchase_contract_item_data",
	SapPurchaseOrderHeaderDataWithAddress:                 "sap_purchase_order_header_data_with_address",
	SapPurchaseOrderItemAccountData:                       "sap_purchase_order_item_account_data",
	SapPurchaseOrderItemDataWithAddress:                   "sap_purchase_order_item_data_with_address",
	SapPurchaseOrderItemPricingElementData:                "sap_purchase_order_item_pricing_element_data",
	SapPurchaseOrderItemScheduleLineData:                  "sap_purchase_order_item_schedule_line_data",
	SapPurchaseRequisitionHeaderData:                      "sap_purchase_requisition_header_data",
	SapPurchaseRequisitionItemAccountAssignmentData:       "sap_purchase_requisition_item_account_assignment_data",
	SapPurchaseRequisitionItemData:                        "sap_purchase_requisition_item_data",
	SapPurchaseRequisitionItemDeliveryAddressData:         "sap_purchase_requisition_item_delivery_address_data",
	SapPurchasingInfoRecordGeneralData:                    "sap_purchasing_info_record_general_data",
	SapPurchasingInfoRecordOrganizationPlantData:          "sap_purchasing_info_record_organization_plant_data",
	SapPurchasingInfoRecordPricingConditionData:           "sap_purchasing_info_record_pricing_condition_data",
	SapPurchasingInfoRecordPricingConditionValidityData:   "sap_purchasing_info_record_pricing_condition_validity_data",
	SapPurchasingSourceListData:                           "sap_purchasing_source_list_data",
	SapReservationDocumentHeaderData:                      "sap_reservation_document_header_data",
	SapReservationDocumentItemData:                        "sap_reservation_document_item_data",
	SapSalesContractHeaderData:                            "sap_sales_contract_header_data",
	SapSalesContractHeaderPartnerData:                     "sap_sales_contract_header_partner_data",
	SapSalesContractItemData:                              "sap_sales_contract_item_data",
	SapSalesContractItemPricingElementData:                "sap_sales_contract_item_pricing_element_data",
	SapSalesInquiryHeaderData:                             "sap_sales_inquiry_header_data",
	SapSalesInquiryHeaderPartnerData:                      "sap_sales_inquiry_header_partner_data",
	SapSalesInquiryItemData:                               "sap_sales_inquiry_item_data",
	SapSalesInquiryItemPricingElementData:                 "sap_sales_inquiry_item_pricing_element_data",
	SapSalesOrderHeaderData:                               "sap_sales_order_header_data",
	SapSalesOrderHeaderPartnerData:                        "sap_sales_order_header_partner_data",
	SapSalesOrderItemData:                                 "sap_sales_order_item_data",
	SapSalesOrderItemPricingElementData:                   "sap_sales_order_item_pricing_element_data",
	SapSalesOrderItemScheduleLineData:                     "sap_sales_order_item_schedule_line_data",
	SapSalesPricingConditionRecordData:                    "sap_sales_pricing_condition_record_data",
	SapSalesPricingConditionValidityData:                  "sap_sales_pricing_condition_validity_data",
	SapSalesQuotationHeaderData:                           "sap_sales_quotation_header_data",
	SapSalesQuotationHeaderPartnerData:                    "sap_sales_quotation_header_partner_data",
	SapSalesQuotationItemData:                             "sap_sales_quotation_item_data",
	SapSalesQuotationItemPricingElementData:               "sap_sales_quotation_item_pricing_element_data",
	SapSalesSchedulingAgreementHeaderData:                 "sap_sales_scheduling_agreement_header_data",
	SapSalesSchedulingAgreementItemData:                   "sap_sales_scheduling_agreement_item_data",
	SapSalesSchedulingAgreementItemDeliveryScheduleData:   "sap_sales_scheduling_agreement_item_delivery_schedule_data",
	SapSalesSchedulingAgreementItemPricingElementData:     "sap_sales_scheduling_agreement_item_pricing_element_data",
	SapSalesSchedulingAgreementItemScheduleLineData:       "sap_sales_scheduling_agreement_item_schedule_line_data",
	SapSalesSchedulingAgreementPartnerData:                "sap_sales_scheduling_agreement_partner_data",
	SapServiceConfirmationHeaderData:                      "sap_service_confirmation_header_data",
	SapServiceConfirmationItemData:                        "sap_service_confirmation_item_data",
	SapServiceConfirmationItemPricingElementData:          "sap_service_confirmation_item_pricing_element_data",
	SapServiceConfirmationPersonResponsibleData:           "sap_service_confirmation_person_responsible_data",
	SapServiceConfirmationReferenceObjectData:             "sap_service_confirmation_reference_object_data",
	SapServiceOrderConfirmationData:                       "sap_service_order_confirmation_data",
	SapServiceOrderDefectData:                             "sap_service_order_defect_data",
	SapServiceOrderHeaderData:                             "sap_service_order_header_data",
	SapServiceOrderItemData:                               "sap_service_order_item_data",
	SapServiceOrderItemPricingElementData:                 "sap_service_order_item_pricing_element_data",
	SapServiceOrderPersonResponsibleData:                  "sap_service_order_person_responsible_data",
	SapServiceOrderReferenceObjectData:                    "sap_service_order_reference_object_data",
	SapSupplierInvoiceAccountAssignmentData:               "sap_supplier_invoice_account_assignment_data",
	SapSupplierInvoiceHeaderData:                          "sap_supplier_invoice_header_data",
	SapSupplierInvoicePurchaseOrderReferenceData:          "sap_supplier_invoice_purchase_order_reference_data",
	SapSupplierInvoiceTaxData:                             "sap_supplier_invoice_tax_data",
	SapWarehouseResourceHeaderData:                        "sap_warehouse_resource_header_data",
	SapWorkCenterData:                                     "sap_work_center_data",
}
