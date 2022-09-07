package sap_api_input_reader

import (
	"sap-api-integrations-debit-memo-request-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.DebitMemoRequest
	return &requests.Header{
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
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataDebitMemoRequest := sdc.DebitMemoRequest
	data := sdc.DebitMemoRequest.DebitMemoRequestItem
	return &requests.Item{
		DebitMemoRequest:             dataDebitMemoRequest.DebitMemoRequest,
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
}
