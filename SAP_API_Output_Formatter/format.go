package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-debit-memo-request-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		DebitMemoRequest:              data.DebitMemoRequest,
		DebitMemoRequestType:          data.DebitMemoRequestType,
		SalesOrganization:             data.SalesOrganization,
		DistributionChannel:           data.DistributionChannel,
		OrganizationDivision:          data.OrganizationDivision,
		SalesGroup:                    data.SalesGroup,
		SalesOffice:                   data.SalesOffice,
		SalesDistrict:                 data.SalesDistrict,
		SoldToParty:                   data.SoldToParty,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		LastChangeDateTime:            data.LastChangeDateTime,
		PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderType:     data.CustomerPurchaseOrderType,
		CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
		DebitMemoRequestDate:          data.DebitMemoRequestDate,
		TotalNetAmount:                data.TotalNetAmount,
		TransactionCurrency:           data.TransactionCurrency,
		SDDocumentReason:              data.SDDocumentReason,
		PricingDate:                   data.PricingDate,
		CustomerTaxClassification1:    data.CustomerTaxClassification1,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		IncotermsClassification:       data.IncotermsClassification,
		CustomerPaymentTerms:          data.CustomerPaymentTerms,
		PaymentMethod:                 data.PaymentMethod,
		BillingDocumentDate:           data.BillingDocumentDate,
		ReferenceSDDocument:           data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:   data.ReferenceSDDocumentCategory,
		OverallSDProcessStatus:        data.OverallSDProcessStatus,
		TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
		OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
		OverallOrdReltdBillgStatus:    data.OverallOrdReltdBillgStatus,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		DebitMemoRequest:             data.DebitMemoRequest,
		DebitMemoRequestItem:         data.DebitMemoRequestItem,
		HigherLevelItem:              data.HigherLevelItem,
		DebitMemoRequestItemCategory: data.DebitMemoRequestItemCategory,
		DebitMemoRequestItemText:     data.DebitMemoRequestItemText,
		PurchaseOrderByCustomer:      data.PurchaseOrderByCustomer,
		Material:                     data.Material,
		MaterialByCustomer:           data.MaterialByCustomer,
		PricingDate:                  data.PricingDate,
		RequestedQuantity:            data.RequestedQuantity,
		RequestedQuantityUnit:        data.RequestedQuantityUnit,
		ItemGrossWeight:              data.ItemGrossWeight,
		ItemNetWeight:                data.ItemNetWeight,
		ItemWeightUnit:               data.ItemWeightUnit,
		ItemVolume:                   data.ItemVolume,
		ItemVolumeUnit:               data.ItemVolumeUnit,
		TransactionCurrency:          data.TransactionCurrency,
		NetAmount:                    data.NetAmount,
		MaterialGroup:                data.MaterialGroup,
		MaterialPricingGroup:         data.MaterialPricingGroup,
		ProductTaxClassification1:    data.ProductTaxClassification1,
		Batch:                        data.Batch,
		Plant:                        data.Plant,
		IncotermsClassification:      data.IncotermsClassification,
		CustomerPaymentTerms:         data.CustomerPaymentTerms,
		ItemBillingBlockReason:       data.ItemBillingBlockReason,
		SalesDocumentRjcnReason:      data.SalesDocumentRjcnReason,
		WBSElement:                   data.WBSElement,
		ProfitCenter:                 data.ProfitCenter,
		ReferenceSDDocument:          data.ReferenceSDDocument,
		ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
		SDProcessStatus:              data.SDProcessStatus,
		OrderRelatedBillingStatus:    data.OrderRelatedBillingStatus,
	}

	return item, nil
}
