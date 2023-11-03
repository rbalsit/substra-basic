package main

import (
	"encoding/json"
	"fmt"
	"time"
	"log"
	// "github.com/google/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	logger "github.com/sirupsen/logrus"
)

// ProductContract contract for managing CRUD for Product
type ProductContract struct {
	contractapi.Contract
}

// org sequence for validating asset from provenance to distributor
var orgSequence = [...]string{"MSDOrg1MSP","MSDOrg3MSP"}

type HistoryQueryResult struct {
	Record    *scanProduct    `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
	OrgName   string 	 `json:"orgName"`
}


var _logger = logger.New()

const transferAgreementObjectType = "merckData"

func (c *ProductContract) CreateBulkScannedData(ctx contractapi.TransactionContextInterface) (string, error) {

	products := []scanProduct{}
		
	transientData, _ := ctx.GetStub().GetTransient()
	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		_logger.Error("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &products)
	if err != nil {
		_logger.Error("failed to unmarshal JSON: %s", err.Error())
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	for i := 0; i < len(products); i += 1 {
		product := products[i]
		ledgerproduct := new(scanProduct)
		// uid := uuid.New()
		clientMspID, err1 := ctx.GetClientIdentity().GetMSPID()
		if err1 != nil {
			return "",fmt.Errorf("failed to get verified MSP ID: %v",err1)
		}
		// exists, _ = c.MyPrivateAssetExists(ctx, product.CombinedKey)
		// if exists {
		// 	_logger.Error("The asset %s already exists", product.CombinedKey)
		// }
		ledgerproduct.ObjectType = "SCAN"
		ledgerproduct.CombinedKey = product.CombinedKey
		ledgerproduct.GTIN_SN = product.GTIN_SN
		ledgerproduct.SerialID = product.SerialID
		ledgerproduct.BatchID = product.BatchID
		ledgerproduct.MaterialName = product.MaterialName
		ledgerproduct.MaterialDescription = product.MaterialDescription
		ledgerproduct.ManufacturingDate = product.ManufacturingDate
		ledgerproduct.ExpiryDate = product.ExpiryDate
		ledgerproduct.ItemStatus = product.ItemStatus
		ledgerproduct.Timemillies = product.Timemillies
		ledgerproduct.UserId = product.UserId
		ledgerproduct.MSPId = clientMspID
		ledgerproduct.DeviceUuid = product.DeviceUuid
		ledgerproduct.ScannedResult = product.ScannedResult
		ledgerproduct.ScanLocation = product.ScanLocation
		ledgerproduct.ShipToParty = product.ShipToParty
		ledgerproduct.LastUpdatedat = currentTime	
		// ledgerproduct.DispositionStatus = product.DispositionStatus
		// ledgerproduct.ParentSequenceNo = product.ParentSequenceNo
		// ledgerproduct.PartnerCountry = product.PartnerCountry
		// ledgerproduct.LastUpdate = product.LastUpdate
		// ledgerproduct.SysGenUniqueID = uid.String()       // UUID random
		// ledgerproduct.OffChainHash = product.OffChainHash // hash for off chain database record

		// productPrivateDetails := new(ProductPrivateDetails)
		// productPrivateDetails.ObjectType = "PRIVATE_PRODUCT"
		// productPrivateDetails.CombinedKey = product.CombinedKey
		// productPrivateDetails.MaterialDescription = product.MaterialDescription
		// productPrivateDetails.ManufacturingDate = product.ManufacturingDate
		// productPrivateDetails.ShippingDate = product.ShippingDate
		// productPrivateDetails.ShippingSiteID = product.ShippingSiteID

		productJSONasBytes, err := json.Marshal(ledgerproduct)
		if err != nil {
			_logger.Error(err.Error())
		}

		keyParts := []string{product.CombinedKey,currentTime}

		transferAgreeKey, err := ctx.GetStub().CreateCompositeKey(transferAgreementObjectType,keyParts)
		if err != nil {
			return "nil",err
		}

		// === Save product to state ===
		err = ctx.GetStub().PutState(transferAgreeKey, productJSONasBytes)
		if err != nil {
			_logger.Error("failed to put Product: %s", err.Error())
		}
		// bytes, _ := json.Marshal(productPrivateDetails)

		// === Save product to Private state ===
		// err = ctx.GetStub().PutPrivateData(collectionName, product.CombinedKey, bytes)
		// if err != nil {
		// 	_logger.Error("failed to put Private Product: %s", err.Error())
		// }
	}
	
	return "Success", nil
}

// CreateAsset creates a new instance of ProductContract
// func (c *ProductContract) CreateScannedData(ctx contractapi.TransactionContextInterface, combinedKey string ,gtisn string, batchID string,serialID string, expiryDate string,scannedresult string,lat string,lng string,timemillies string,userid string) error {

// 	ledgerproduct := new(scanProduct)

// 	ledgerproduct.ObjectType = "SCAN"
// 	ledgerproduct.CombinedKey = combinedKey
// 	ledgerproduct.GTIN_SN = gtisn
// 	ledgerproduct.BatchID = batchID
// 	ledgerproduct.ExpiryDate = expiryDate
// 	ledgerproduct.Scannedresult = scannedresult
// 	ledgerproduct.Lat = lat
// 	ledgerproduct.Lng = lng
// 	ledgerproduct.Timemillies = timemillies
// 	ledgerproduct.Userid = userid
// 	ledgerproduct.SerialID = serialID


// 	productJSONasBytes, err := json.Marshal(ledgerproduct)
// 	if err != nil {
// 		_logger.Error(err.Error())
// 	}

// 	// === Save product to world state ===
// 	// err = ctx.GetStub().PutState(combinedKey, productJSONasBytes)
// 	// if err != nil {
// 	// 	_logger.Error("failed to put Product: %s", err.Error())
// 	// }

// 	// return ctx.GetStub().SetEvent("ADD_COMMISSIONING", productJSONasBytes)
// 	// return ledgerproduct.SysGenUniqueID, nil
// 	return  ctx.GetStub().PutState(combinedKey, productJSONasBytes)
// }


func (c *ProductContract) UpdateScannedData(ctx contractapi.TransactionContextInterface) (string,error) {
	products := []scanProduct{}
	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		return "",fmt.Errorf("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &products)
	if err != nil {
		return  "",fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	for i := 0; i < len(products); i += 1 {
		product := products[i]
		// ledgerproduct := new(scanProduct)
	// product := new(scanProduct)
	clientMspID, err3 := ctx.GetClientIdentity().GetMSPID()
	if err3 != nil {
		return "",fmt.Errorf("failed to get verified MSP ID: %v",err3)
	}

	bytes,err4 := ctx.GetStub().GetStateByPartialCompositeKey(transferAgreementObjectType, []string{product.CombinedKey})
	if err4 != nil {
		return  "",fmt.Errorf("Could not read from world state. %s", err4)
	}

	var myLedgerAssetCompositeKey scanProduct

	var records []scanProduct

	for bytes.HasNext() {
		responseRange, err := bytes.Next()

		log.Printf("ReadAsset: collection 1 %v",responseRange)

		if len(responseRange.Value) > 0 {
			err = json.Unmarshal(responseRange.Value, &myLedgerAssetCompositeKey)
			if err != nil {
				return "nil", fmt.Errorf("failed to get asset: %v", err)
			}
		} else {
			myLedgerAssetCompositeKey = scanProduct{
				CombinedKey: product.CombinedKey,
			}
		}
		records=append(records, myLedgerAssetCompositeKey)
	}

    myLedgerAsset := records[len(records)-1]

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	myLedgerAsset.ScanLocation = product.ScanLocation
	myLedgerAsset.UserId = product.UserId
	myLedgerAsset.Timemillies = product.Timemillies
	myLedgerAsset.DeviceUuid = product.DeviceUuid
	myLedgerAsset.MSPId = clientMspID
	myLedgerAsset.LastUpdatedat = currentTime	
	
	// myLedgerAsset.Userid = product.OffChainHash
	// currentDate := time.Now().UTC()
	// expireDate := time.Parse("2006-02-17",myLedgerAsset.ExpiryDate)
	// if myLedgerAsset.ItemStatus == "active" {
	// 	myLedgerAsset.ScannedResult = "Verified"

	// 	// eventName = "SHIPPING"
	// }
	fmt.Println("expiry date :",myLedgerAsset.ExpiryDate)
	expiryDateString , _ := time.Parse("2006-01-02",myLedgerAsset.ExpiryDate)
	fmt.Println("expiry date 1:",expiryDateString)
    if (expiryDateString.Before(time.Now().UTC())) {
		myLedgerAsset.ScannedResult = "Expired"
    } else if myLedgerAsset.ShipToParty != "" {
	 if myLedgerAsset.ShipToParty == product.UserId {
		myLedgerAsset.ScannedResult = "Verified"
		myLedgerAsset.ShipToParty = ""
	 } else {
		myLedgerAsset.ScannedResult = "Not Verified by "+myLedgerAsset.ShipToParty
	 }
	} else if myLedgerAsset.ScannedResult == "Verified" {
		result,err := c.VerifyScanHistory(ctx,product.CombinedKey)
		if(err !=nil){
			return  "",fmt.Errorf("failed to get history : %s", err.Error())
		}
		myLedgerAsset.ScannedResult = result
	} else {
		myLedgerAsset.ScannedResult = "Not Verified"
	}
	// else if (expireDate.After(time.Now().UTC()))   { // change matching data
	// 	myLedgerAsset.ScannedResult = "Expired"
	// } 
	// else {
	// 	myLedgerAsset.ScannedResult = "Verified"
	// }

	productJSONasBytes, err6 := json.Marshal(myLedgerAsset)
	if err6 != nil {
		return  "",fmt.Errorf(err6.Error())
	}

	// === Save product to state ===
	// ctx.GetStub().SetEvent(eventName, productJSONasBytes)
	keyParts := []string{product.CombinedKey,currentTime}

	transferAgreeKey, err := ctx.GetStub().CreateCompositeKey(transferAgreementObjectType,keyParts)
	if err != nil {
		return "nil",err
	}
	
	err2 := ctx.GetStub().PutState(transferAgreeKey, productJSONasBytes)
	if err2 != nil {
		_logger.Error("failed to put Product: %s", err2.Error())
		return  "",fmt.Errorf(err2.Error())
	}
}
	return "Success", nil
}

func (c *ProductContract) VerifyScanHistory(ctx contractapi.TransactionContextInterface,combinedKey string) (string, error) { 

	records,err :=c.GetScanHistory(ctx,combinedKey)
	if err != nil{
		return "",err
	}
	jsonbytes,err := json.Marshal(records)
	var orgs []map[string]interface{}
	err1 := json.Unmarshal(jsonbytes,&orgs)
	if err1 != nil {
		panic(err1)
	}
	if (orgs[0]["mspId"] == orgSequence[0]) && (orgs[1]["mspId"] == orgSequence[1] || orgs[1]["mspId"] == "MSDOrg4MSP")  {
		return "Verified",nil
	} else{
		return "Not Verifed",nil
	}

}

// Get scanned data retrieves an instance of Asset from the ledger
func (c *ProductContract) GetScannedData(ctx contractapi.TransactionContextInterface,combinedKey string) (*scanProduct, error) { //no input parameter

	var records []scanProduct

	iDAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(transferAgreementObjectType, []string{combinedKey})
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	log.Printf("ReadAsset: collection 1111 %v",iDAssetResultsIterator)
	defer iDAssetResultsIterator.Close()

	// var id []byte
	for iDAssetResultsIterator.HasNext() {
		responseRange, err := iDAssetResultsIterator.Next()

		log.Printf("ReadAsset: collection 1 %v",responseRange)

		var assetData scanProduct

		if len(responseRange.Value) > 0 {
			err = json.Unmarshal(responseRange.Value, &assetData)
			if err != nil {
				return nil, fmt.Errorf("failed to get asset: %v", err)
			}
		} else {
			assetData = scanProduct{
				CombinedKey: combinedKey,
			}
		}
	
		records=append(records, assetData)
	}
	myLedgerAsset := records[len(records)-1]

	return &myLedgerAsset, nil

}

func (c *ProductContract) GetScanHistory(ctx contractapi.TransactionContextInterface,combinedKey string) ([]scanProduct, error) { //no input parameter

	var records []scanProduct

	iDAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(transferAgreementObjectType, []string{combinedKey})
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	log.Printf("ReadAsset: collection 1111 %v",iDAssetResultsIterator)
	defer iDAssetResultsIterator.Close()

	// var id []byte
	for iDAssetResultsIterator.HasNext() {
		responseRange, err := iDAssetResultsIterator.Next()

		log.Printf("ReadAsset: collection 1 %v",responseRange)

		var assetData scanProduct

		if len(responseRange.Value) > 0 {
			err = json.Unmarshal(responseRange.Value, &assetData)
			if err != nil {
				return nil, fmt.Errorf("failed to get asset: %v", err)
			}
		} else {
			assetData = scanProduct{
				CombinedKey: combinedKey,
			}
		}
	
		records=append(records, assetData)
	}
	return records, nil

}

func (c *ProductContract) VerifyScan(ctx contractapi.TransactionContextInterface,combinedKey string) ([]HistoryQueryResult, error) { //no input parameter


	resultsIterator, err := ctx.GetStub().GetHistoryForKey(combinedKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset scanProduct
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = scanProduct{
				CombinedKey: combinedKey,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
			OrgName:	asset.MSPId,
		}
		records = append(records, record)
	}
	return records, nil

}


func (c *ProductContract) GetByUserId(ctx contractapi.TransactionContextInterface, userId string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_USER, userId)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByDates(ctx contractapi.TransactionContextInterface, startdate string, enddate string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_DATES, startdate,enddate)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByScannedResults(ctx contractapi.TransactionContextInterface,scannedresult string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_SCANNED_RESULTS,scannedresult)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByDatesWithScannedLocation(ctx contractapi.TransactionContextInterface, startdate string, enddate string,scannedlocation string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_DATES_WITH_SCANNED_LOCATION, startdate,enddate,scannedlocation)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByDatesWithScannedResults(ctx contractapi.TransactionContextInterface, startdate string, enddate string,scannedresult string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_DATES_WITH_SCANNED_RESULTS, startdate,enddate,scannedresult)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByDatesWithScannedResultsandScannedLocation(ctx contractapi.TransactionContextInterface, startdate string, enddate string,scannedresult string,scannedlocation string) ([]*scanProduct, error) {
	queryString := fmt.Sprintf(QUERY_BY_DATES_WITH_SCANNED_RESULTS_AND_SCANNED_LOCATION, startdate,enddate,scannedresult,scannedlocation)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (c *ProductContract) GetByQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*scanProduct, error) {
	// queryString := fmt.Sprintf(QUERY_BY_USER, userId)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
func (c *ProductContract) GetBySerialId(ctx contractapi.TransactionContextInterface, serialID string) ([]*scanProduct, error) {
	
	queryString := fmt.Sprintf(QUERY_BY_SERIALID,serialID)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
func (c *ProductContract) GetByUserIdAndSerialId(ctx contractapi.TransactionContextInterface,userId string, serialID string) ([]*scanProduct, error) {

	queryString := fmt.Sprintf(QUERY_BY_USER_AND_SERIALID, userId,serialID)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	var assets []*scanProduct
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset scanProduct
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

