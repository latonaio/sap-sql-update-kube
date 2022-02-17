package main

import (
	"context"
	"encoding/json"
	"sap-sql-update-kube/config"
	"sap-sql-update-kube/database"
	"sap-sql-update-kube/database/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/latonaio/golang-logging-library/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	l := logger.NewLogger()

	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), nil)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	ctx := context.Background()

	for msg := range iter {
		data := msg.Data()
		str, _ := json.Marshal(msg.Data()["message"])
		f := data["function"].(string)
		msg.Success()

		switch f {
		case "ClassificationClass":
			pms := &[]models.SapClassificationClassDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ClassificationCharacteristic":
			pms := &[]models.SapClassificationCharacteristicDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ClassificationClassDescription":
			pms := &[]models.SapClassificationClassDescriptionDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CharacteristicCharacteristic":
			pms := &[]models.SapCharacteristicCharacteristicDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CharacteristicValue":
			pms := &[]models.SapCharacteristicValueDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CharacteristicCharcDescription":
			pms := &[]models.SapCharacteristicCharcDescriptionDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CharacteristicValueDescription":
			pms := &[]models.SapCharacteristicValueDescriptionDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BankMasterBank":
			pms := &[]models.SapBankMasterBankDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerGeneral":
			pms := &[]models.SapBusinessPartnerGeneralDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerRole":
			pms := &[]models.SapBusinessPartnerRoleDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerAddress":
			pms := &[]models.SapBusinessPartnerAddressDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerBank":
			pms := &[]models.SapBusinessPartnerBankDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterGeneral":
			pms := &[]models.SapProductMasterGeneralDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterSalesPlant":
			pds := &[]models.SapProductMasterSalesPlantDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterPlant":
			pds := &[]models.SapProductMasterPlantDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterMRPArea":
			mas := &[]models.SapProductMasterMRPAreaDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterProcurement":
			pds := &[]models.SapProductMasterProcurementDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterAccounting":
			pds := &[]models.SapProductMasterAccountingDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterSalesOrganization":
			pds := &[]models.SapProductMasterSalesOrganizationDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterWorkScheduling":
			pds := &[]models.SapProductMasterWorkSchedulingDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterQuality":
			pds := &[]models.SapProductMasterQualityDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterSalesTax":
			pds := &[]models.SapProductMasterSalesTaxDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterProductDescription":
			pds := &[]models.SapProductMasterProductDescriptionDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterClassProductGeneral":
			pds := &[]models.SapProductMasterClassProductGeneralDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterCloassProductClass":
			pds := &[]models.SapProductMasterClassProductClassDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterClassClassDetails":
			pds := &[]models.SapProductMasterClassClassDetailsDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductMasterClassProductCharacteristic":
			pds := &[]models.SapProductMasterClassProductCharacteristicDatum{}
			json.Unmarshal(str, pds)
			for _, pd := range *pds {
				err = pd.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pd.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "Batch":
			mas := &[]models.SapBatchMasterRecord{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaterialStock":
			mas := &[]models.SapMaterialStockDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ReservationDocumentHeader":
			mas := &[]models.SapReservationDocumentHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ReservationDocumentItem":
			mas := &[]models.SapReservationDocumentItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaterialDocumentHeader":
			mas := &[]models.SapMaterialDocumentHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaterialDocumentItem":
			mas := &[]models.SapMaterialDocumentItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PhysicalInventoryDocumentHeader":
			mas := &[]models.SapPhysicalInventoryDocumentHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PhysicalInventoryDocumentItem":
			mas := &[]models.SapPhysicalInventoryDocumentItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerToCustomer":
			sd := &models.SapBusinessPartnerCustomerDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "BusinessPartnerCustomer":
			mas := &[]models.SapBusinessPartnerCustomerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerCustomerSalesArea":
			mas := &[]models.SapBusinessPartnerCustomerSalesAreaDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerCustomerPartnerFunction":
			mas := &[]models.SapBusinessPartnerCustomerPartnerFunctionDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerCustomerCompany":
			mas := &[]models.SapBusinessPartnerCustomerCompanyDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerMaterial":
			mas := &[]models.SapCustomerMaterialDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesPricingConditionValidity":
			mas := &[]models.SapSalesPricingConditionValidityDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesPricingToConditionRecord":
			sd := &models.SapSalesPricingConditionRecordDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "SalesInquiryHeader":
			mas := &[]models.SapSalesInquiryHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesInquiryHeaderPartner":
			mas := &[]models.SapSalesInquiryHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesInquiryItem":
			mas := &[]models.SapSalesInquiryItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesInquiryItemPricingElement":
			mas := &[]models.SapSalesInquiryItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesQuotationHeader":
			mas := &[]models.SapSalesQuotationHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesQuotationHeaderPartner":
			mas := &[]models.SapSalesQuotationHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesQuotationItem":
			mas := &[]models.SapSalesQuotationItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesQuotationItemPricingElement":
			mas := &[]models.SapSalesQuotationItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesOrderHeader":
			mas := &[]models.SapSalesOrderHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesOrderItem":
			mas := &[]models.SapSalesOrderItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesOrderHeaderPartner":
			mas := &[]models.SapSalesOrderHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesOrderItemPricingElement":
			mas := &[]models.SapSalesOrderItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesOrderItemScheduleLine":
			mas := &[]models.SapSalesOrderItemScheduleLineDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesContractHeader":
			mas := &[]models.SapSalesContractHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesContractHeaderPartner":
			mas := &[]models.SapSalesContractHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesContractItem":
			mas := &[]models.SapSalesContractItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesContractItemPricingElement":
			mas := &[]models.SapSalesContractItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementHeader":
			mas := &[]models.SapSalesSchedulingAgreementHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementPartner":
			mas := &[]models.SapSalesSchedulingAgreementPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementItem":
			mas := &[]models.SapSalesSchedulingAgreementItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementItemDeliverySchedule":
			mas := &[]models.SapSalesSchedulingAgreementItemDeliveryScheduleDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementItemScheduleLine":
			mas := &[]models.SapSalesSchedulingAgreementItemScheduleLineDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SalesSchedulingAgreementItemPricingElement":
			mas := &[]models.SapSalesSchedulingAgreementItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "OutboundDeliveryHeader":
			mas := &[]models.SapOutboundDeliveryHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "OutboundDeliveryHeaderPartner":
			mas := &[]models.SapOutboundDeliveryHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "OutboundDeliveryPartnerAddress":
			sd := &models.SapOutboundDeliveryPartnerAddressDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "OutboundDeliveryItem":
			mas := &[]models.SapOutboundDeliveryItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "OutboundDeliveryItemDocumentFlow":
			mas := &[]models.SapOutboundDeliveryItemDocumentFlowDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnHeader":
			mas := &[]models.SapCustomerReturnHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnHeaderPartner":
			mas := &[]models.SapCustomerReturnHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnItem":
			mas := &[]models.SapCustomerReturnItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnItemPricingElement":
			mas := &[]models.SapCustomerReturnItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnItemProcessStep":
			mas := &[]models.SapCustomerReturnItemProcessStepDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CustomerReturnItemScheduleLine":
			mas := &[]models.SapCustomerReturnItemScheduleLineDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillingDocumentHeader":
			mas := &[]models.SapBillingDocumentHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillingDocumentHeaderPartner":
			mas := &[]models.SapBillingDocumentHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillingDocumentItem":
			mas := &[]models.SapBillingDocumentItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillingDocumentItemPartner":
			mas := &[]models.SapBillingDocumentItemPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillingDocumentItemPricingElement":
			mas := &[]models.SapBillingDocumentItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CreditMemoRequestHeader":
			mas := &[]models.SapCreditMemoRequestHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CreditMemoRequestHeaderPartner":
			mas := &[]models.SapCreditMemoRequestHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CreditMemoRequestItem":
			mas := &[]models.SapCreditMemoRequestItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "CreditMemoRequestItemPricingElement":
			mas := &[]models.SapCreditMemoRequestItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "DebitMemoRequestHeader":
			mas := &[]models.SapDebitMemoRequestHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "DebitMemoRequestHeaderPartner":
			mas := &[]models.SapDebitMemoRequestHeaderPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "DebitMemoRequestItem":
			mas := &[]models.SapDebitMemoRequestItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "DebitMemoRequestItemPricingElement":
			mas := &[]models.SapDebitMemoRequestItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerToSupplier":
			sd := &models.SapBusinessPartnerSupplierDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "BusinessPartnerSupplier":
			mas := &[]models.SapBusinessPartnerSupplierDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerSupplierPurchasingOrganization":
			mas := &[]models.SapBusinessPartnerSupplierPurchasingOrganizationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerSupplierPartnerFunction":
			mas := &[]models.SapBusinessPartnerSupplierPartnerFunctionDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BusinessPartnerSupplierCompany":
			mas := &[]models.SapBusinessPartnerSupplierCompanyDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchasingSourceList":
			mas := &[]models.SapPurchasingSourceListDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchasingInfoRecordGeneral":
			mas := &[]models.SapPurchasingInfoRecordGeneralDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchasingInfoRecordOrganizationPlant":
			mas := &[]models.SapPurchasingInfoRecordOrganizationPlantDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchasingInfoRecordPricingConditionValidity":
			mas := &[]models.SapPurchasingInfoRecordPricingConditionValidityDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchasingInfoRecordPricingCondition":
			mas := &[]models.SapPurchasingInfoRecordPricingConditionDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseRequisitionItemDeliveryAddress":
			mas := &[]models.SapPurchaseRequisitionItemDeliveryAddressDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseRequisitionToItemDeliveryAddress":
			sd := &models.SapPurchaseRequisitionItemDeliveryAddressDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "PurchaseRequisitionItemAccount":
			mas := &[]models.SapPurchaseRequisitionItemAccountAssignmentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseOrderHeader":
			mas := &[]models.SapPurchaseOrderHeaderDataWithAddress{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseOrderItem":
			mas := &[]models.SapPurchaseOrderItemDataWithAddress{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseOrderItemScheduleLine":
			mas := &[]models.SapPurchaseOrderItemScheduleLineDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseOrderItemPricingElement":
			mas := &[]models.SapPurchaseOrderItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseOrderItemAccount":
			mas := &[]models.SapPurchaseOrderItemAccountDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseContractHeader":
			mas := &[]models.SapPurchaseContractHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseContractItem":
			mas := &[]models.SapPurchaseContractItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseContractAddress":
			mas := &[]models.SapPurchaseContractAddressDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PurchaseContractItemCondition":
			mas := &[]models.SapPurchaseContractItemConditionDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SupplierInvoiceHeader":
			mas := &[]models.SapSupplierInvoiceHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SupplierInvoiceTax":
			mas := &[]models.SapSupplierInvoiceTaxDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SupplierInvoiceAccountAssignment":
			mas := &[]models.SapSupplierInvoiceAccountAssignmentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "SupplierInvoicePurchaseOrderReference":
			mas := &[]models.SapSupplierInvoicePurchaseOrderReferenceDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "InboundDeliveryHeader":
			mas := &[]models.SapInboundDeliveryHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "InboundDeliveryPartner":
			mas := &[]models.SapInboundDeliveryPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "InboundDeliveryToAddress":
			sd := &models.SapInboundDeliveryAddressDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		case "InboundDeliveryItem":
			mas := &[]models.SapInboundDeliveryItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillOfMaterialHeader":
			mas := &[]models.SapBillOfMaterialHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillOfMaterialItem":
			mas := &[]models.SapBillOfMaterialItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "BillOfMaterialWhereUsedList":
			mas := &[]models.SapBillOfMaterialWhereUsedListDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "WorkCenter":
			mas := &[]models.SapWorkCenterDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionRoutingHeader":
			mas := &[]models.SapProductionRoutingHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionRoutingMaterialAssignment":
			mas := &[]models.SapProductionRoutingMaterialAssignmentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionRoutingSequence":
			mas := &[]models.SapProductionRoutingSequenceDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionRoutingOperation":
			mas := &[]models.SapProductionRoutingOperationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionRoutingComponentAllocation":
			mas := &[]models.SapProductionRoutingComponentAllocationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PlannedIndependentRequirementHeader":
			mas := &[]models.SapPlannedIndependentRequirementHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PlannedIndependentRequirementItem":
			mas := &[]models.SapPlannedIndependentRequirementItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PlannedOrderHeader":
			mas := &[]models.SapPlannedOrderHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "PlannedOrderComponent":
			mas := &[]models.SapPlannedOrderComponentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderGeneral":
			mas := &[]models.SapProductionOrderGeneralDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderComponent":
			mas := &[]models.SapProductionOrderComponentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderItem":
			mas := &[]models.SapProductionOrderItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderOperation":
			mas := &[]models.SapProductionOrderOperationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}

		case "ProductionOrderStatus":
			mas := &[]models.SapProductionOrderStatusDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderConfirmationConfirmation":
			mas := &[]models.SapProductionOrderConfirmationConfirmationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderConfirmationMaterialMovements":
			mas := &[]models.SapProductionOrderConfirmationMaterialMovementsDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProductionOrderConfirmationBatchCharacteristic":
			mas := &[]models.SapProductionOrderConfirmationBatchCharacteristicDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProcessOrderConfirmationConfirmation":
			mas := &[]models.SapProcessOrderConfirmationConfirmationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProcessOrderConfirmationMaterialMovements":
			mas := &[]models.SapProcessOrderConfirmationMaterialMovementsDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ProcessOrderConfirmationBatchCharacteristic":
			mas := &[]models.SapProcessOrderConfirmationBatchCharacteristicDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "FunctionalLocationHeader":
			mas := &[]models.SapFunctionalLocationHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "Equipment":
			mas := &[]models.SapEquipmentMasterEquipmentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "EquipmentPartner":
			mas := &[]models.SapEquipmentMasterBusinessPartnerDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceBillOfMaterialHeader":
			mas := &[]models.SapMaintenanceBillOfMaterialHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceBillOfMaterialItem":
			mas := &[]models.SapMaintenanceBillOfMaterialItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenancePlanHeader":
			mas := &[]models.SapMaintenancePlanHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenancePlanStrategyCycle":
			mas := &[]models.SapMaintenancePlanStrategyCycleDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenancePlanItem":
			mas := &[]models.SapMaintenancePlanItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenancePlanCallObjects":
			mas := &[]models.SapMaintenancePlanCallObjectsDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceItemItem":
			mas := &[]models.SapMaintenanceItemItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceItemCallObjects":
			mas := &[]models.SapMaintenanceItemCallObjectsDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceTaskListHeader":
			mas := &[]models.SapMaintenanceTaskListHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceTaskListStrategyPackage":
			mas := &[]models.SapMaintenanceTaskListStrategyPackageDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceTaskListOperation":
			mas := &[]models.SapMaintenanceTaskListOperationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceTaskListOperationMaterial":
			mas := &[]models.SapMaintenanceTaskListOperationMaterialDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceNotificationHeader":
			mas := &[]models.SapMaintenanceNotificationHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceNotificationItem":
			mas := &[]models.SapMaintenanceNotificationItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceNotificationItemCause":
			mas := &[]models.SapMaintenanceNotificationItemCauseDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceOrderHeader":
			mas := &[]models.SapMaintenanceOrderHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceOrderObjectListItem":
			mas := &[]models.SapMaintenanceOrderObjectListItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceOrderOperation":
			mas := &[]models.SapMaintenanceOrderOperationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceOrderOperationComponent":
			mas := &[]models.SapMaintenanceOrderOperationComponentDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "MaintenanceOrderConfirmationHeader":
			mas := &[]models.SapMaintenanceOrderConfirmationHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "DefectHeader":
			mas := &[]models.SapDefectHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderHeader":
			mas := &[]models.SapServiceOrderHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderItem":
			mas := &[]models.SapServiceOrderItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderConfirmation":
			mas := &[]models.SapServiceOrderConfirmationDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderDefect":
			mas := &[]models.SapServiceOrderDefectDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderPersonResponsible":
			mas := &[]models.SapServiceOrderPersonResponsibleDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderReferenceObject":
			mas := &[]models.SapServiceOrderReferenceObjectDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceOrderItemPricingElement":
			mas := &[]models.SapServiceOrderItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceConfirmationHeader":
			mas := &[]models.SapServiceConfirmationHeaderDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceConfirmationItem":
			mas := &[]models.SapServiceConfirmationItemDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceConfirmationPersonResponsible":
			mas := &[]models.SapServiceConfirmationPersonResponsibleDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceConfirmationReferenceObject":
			mas := &[]models.SapServiceConfirmationReferenceObjectDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "ServiceConfirmationItemPricingElement":
			mas := &[]models.SapServiceConfirmationItemPricingElementDatum{}
			json.Unmarshal(str, mas)
			for _, ma := range *mas {
				err = ma.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = ma.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		default:
			l.Error("catch unknown function %s", f)
		}
	}
}
