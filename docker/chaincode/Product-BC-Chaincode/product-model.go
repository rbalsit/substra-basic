package main



type scanProduct struct {
	ObjectType          string `json:"docType"` // this field is useful when there are multiple docTypes for queries
	CombinedKey         string `json:"combinedKey"`
	GTIN_SN             string `json:"gtinsn"`   
	SerialID		    string `json:"serialID"`  
	BatchID             string `json:"batchId"`    
	MaterialName        string `json:"materialName"`
	MaterialDescription string `json:"materialDescription"`
	ManufacturingDate   string `json:"manufacturingDate"`
	ExpiryDate          string `json:"expiryDate"`
	ItemStatus          string `json:"itemStatus"`
	Timemillies 	    string `json:"timemillies"`
	UserId 			    string `json:"userId"`
	MSPId 			    string `json:"mspId"`
	DeviceUuid 		    string `json:"deviceUuid"`
	ScannedResult       string `json:"scannedResult"`
	ScanLocation        string `json:"scanLocation"`
	ShipToParty        	string `json:"shipToParty"`
	LastUpdatedat		string `json:"lastUpdatedat"`
}



type Product struct {
	CombinedKey         string `json:"combinedKey"`
	GTIN_SN             string `json:"gtinsn"`
	BatchID             string `json:"batchId"`
	ExpiryDate          string `json:"expiryDate"`
	ItemStatus          string `json:"itemStatus"`
	DispositionStatus   string `json:"dispositionStatus"`
	ParentSequenceNo    string `json:"parentSequenceNo"`
	PickSlipNumber      string `json:"pickSlipNumber"`
	PrincipalID         string `json:"principalId"`
	PartnerCountry      string `json:"partnerCountry"`
	MaterialNumber      string `json:"materialNumber"`
	LastUpdate          string `json:"lastUpdate"`
	DocumentType        string `json:"documentType"`
	DocumentID          string `json:"documentId"`
	SysGenUniqueID      string `json:"sysGenUniqueID"`
	EventAction         string `json:"eventAction"`
	OffChainHash        string `json:"offChainHash"`
	MaterialDescription string `json:"materialDescription"`
	ManufacturingDate   string `json:"manufacturingDate"`
	ShippingDate        string `json:"shippingDate"`
	ShippingSiteID      string `json:"shippingSiteID"`
	MaterialName        string `json:"materialName"`
	InvoiceNumber       string `json:"invoiceNumber"`
	ShipToParty         string `json:"shipToParty"`
	ShipToPartyAddress  string `json:"shipToPartyAddress"`
	CustomerName        string `json:"customerName"`
	CustomerID          string `json:"customerID"`
	CustomerAddress     string `json:"customerAddress"`
	scannedresult		string `json:"scannedresult"`
	lat 				string `json:"lat"`
	lng 				string `json:"lng"`
	timemillies 		string `json:"timemillies"`
	userid 				string `json:"userid"`
	LastUpdatedat		string `json:"lastUpdatedat"`
}

type LedgerProduct struct {
	ObjectType        string `json:"docType"` // this field is useful when there are multiple docTypes for queries
	CombinedKey       string `json:"combinedKey"`
	GTIN_SN           string `json:"gtinsn"`     // in hashed
	BatchID           string `json:"batchId"`    // in hashed
	ExpiryDate        string `json:"expiryDate"` // in hashed
	ItemStatus        string `json:"itemStatus"`
	DispositionStatus string `json:"dispositionStatus"`
	ParentSequenceNo  string `json:"parentSequenceNo"` // in hashed
	PartnerCountry    string `json:"partnerCountry"`
	LastUpdate        string `json:"lastUpdate"`
	SysGenUniqueID    string `json:"sysGenUniqueID"` // generate UUID
	OffChainHash      string `json:"offChainHash"`
}

type QueryProduct struct {
	Gtin         string `json:"gtin"`
	SerialNumber string `json:"serialNumber"`
	LotNumber    string `json:"lotNumber"`
	ExpiryDate   string `json:"expiryDate"`
}

// ProductKey - this struct represents the product key
type ProductKey struct {
	Key          string
	Gtin         string
	SerialNumber string
	LotNumber    string
	ExpiryDate   string
}

type ProductPrivateDetails struct {
	ObjectType          string `json:"docType"`
	CombinedKey         string `json:"combinedKey"`
	MaterialDescription string `json:"materialDescription"`
	ManufacturingDate   string `json:"manufacturingDate"`
	ShippingDate        string `json:"shippingDate"`
	ShippingSiteID      string `json:"shippingSiteID"`
}
type ProductData struct {
	ObjectType string `json:"docType"`
	ID         string `json:"id"`
	Gtin       string `json:"gtin"`
	SerialID   string `json:"serialId"`
	BatchID    string `json:"batchId"`
	ExpiryDate string `json:"expiryDate"`
}
