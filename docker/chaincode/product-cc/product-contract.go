package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
    // "github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	logger "github.com/sirupsen/logrus"
)

// ProductContract contract for managing CRUD for Product
type ProductContract struct {
	contractapi.Contract
}

var _logger = logger.New()

var collectionName string = "collectionProductPrivate-v1"

// CreateAsset creates a new instance of ProductContract
func (c *ProductContract) CreateAsset(ctx contractapi.TransactionContextInterface) (string, error) {

	product := new(Product)
	ledgerproduct := new(LedgerProduct)
	uid := uuid.New()

	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		_logger.Error("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &product)
	if err != nil {
		_logger.Error("failed to unmarshal JSON: %s", err.Error())
	}

	exists, _ = c.MyPrivateAssetExists(ctx, product.CombinedKey)
	if exists {
		_logger.Error("The asset %s already exists", product.CombinedKey)
	}
	ledgerproduct.ObjectType = "PRODUCT"
	ledgerproduct.CombinedKey = product.CombinedKey
	ledgerproduct.GTIN_SN = product.GTIN_SN
	ledgerproduct.BatchID = product.BatchID
	ledgerproduct.ExpiryDate = product.ExpiryDate
	ledgerproduct.ItemStatus = product.ItemStatus
	ledgerproduct.DispositionStatus = product.DispositionStatus
	ledgerproduct.ParentSequenceNo = product.ParentSequenceNo
	ledgerproduct.PartnerCountry = product.PartnerCountry
	ledgerproduct.LastUpdate = product.LastUpdate
	ledgerproduct.SysGenUniqueID = uid.String()       // UUID random
	ledgerproduct.OffChainHash = product.OffChainHash // hash for off chain database record

	productPrivateDetails := new(ProductPrivateDetails)
	productPrivateDetails.ObjectType = "PRIVATE_PRODUCT"
	productPrivateDetails.CombinedKey = product.CombinedKey
	productPrivateDetails.MaterialDescription = product.MaterialDescription
	productPrivateDetails.ManufacturingDate = product.ManufacturingDate
	productPrivateDetails.ShippingDate = product.ShippingDate
	productPrivateDetails.ShippingSiteID = product.ShippingSiteID

	productJSONasBytes, err := json.Marshal(ledgerproduct)
	if err != nil {
		_logger.Error(err.Error())
	}

	// === Save product to state ===
	err = ctx.GetStub().PutState(product.CombinedKey, productJSONasBytes)
	if err != nil {
		_logger.Error("failed to put Product: %s", err.Error())
	}

	bytes, _ := json.Marshal(productPrivateDetails)

	// === Save product to Private state ===
	err = ctx.GetStub().PutPrivateData(collectionName, product.CombinedKey, bytes)
	if err != nil {
		_logger.Error("failed to put Private Product: %s", err.Error())
	}
	// return ctx.GetStub().SetEvent("ADD_COMMISSIONING", productJSONasBytes)
	return ledgerproduct.SysGenUniqueID, nil
}

// CreatePrivateAsset creates a new instance of ProductContract
func (c *ProductContract) CreatePrivateAsset(ctx contractapi.TransactionContextInterface) (string, error) {

	product := new(Product)
	ledgerproduct := new(LedgerProduct)
	// uid := uuid.New()

	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		_logger.Error("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &product)
	if err != nil {
		_logger.Error("failed to unmarshal JSON: %s", err.Error())
	}

	exists, _ = c.MyPrivateAssetExists(ctx, product.CombinedKey)
	if exists {
		_logger.Error("The asset %s already exists", product.CombinedKey)
	}

	productPrivateDetails := new(ProductPrivateDetails)
	productPrivateDetails.ObjectType = "PRIVATE_PRODUCT"
	productPrivateDetails.CombinedKey = product.CombinedKey
	productPrivateDetails.MaterialDescription = product.MaterialDescription
	productPrivateDetails.ManufacturingDate = product.ManufacturingDate
	productPrivateDetails.ShippingDate = product.ShippingDate
	productPrivateDetails.ShippingSiteID = product.ShippingSiteID

	bytes, _ := json.Marshal(productPrivateDetails)

	err = ctx.GetStub().PutPrivateData(collectionName, product.CombinedKey, bytes)
	if err != nil {
		_logger.Error("failed to put Private Product: %s", err.Error())
	}
	return ledgerproduct.SysGenUniqueID, nil
}



// CreateAsset creates a new instance of ProductContract
func (c *ProductContract) CreateBulkAssets(ctx contractapi.TransactionContextInterface) (string, error) {

	products := []Product{}
		
	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		_logger.Error("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &products)
	if err != nil {
		_logger.Error("failed to unmarshal JSON: %s", err.Error())
	}

	for i := 0; i < len(products); i += 1 {
		product := products[i]
		ledgerproduct := new(LedgerProduct)
		uid := uuid.New()

		exists, _ = c.MyPrivateAssetExists(ctx, product.CombinedKey)
		if exists {
			_logger.Error("The asset %s already exists", product.CombinedKey)
		}

		ledgerproduct.ObjectType = "PRODUCT"
		ledgerproduct.CombinedKey = product.CombinedKey
		ledgerproduct.GTIN_SN = product.GTIN_SN
		ledgerproduct.BatchID = product.BatchID
		ledgerproduct.ExpiryDate = product.ExpiryDate
		ledgerproduct.ItemStatus = product.ItemStatus
		ledgerproduct.DispositionStatus = product.DispositionStatus
		ledgerproduct.ParentSequenceNo = product.ParentSequenceNo
		ledgerproduct.PartnerCountry = product.PartnerCountry
		ledgerproduct.LastUpdate = product.LastUpdate
		ledgerproduct.SysGenUniqueID = uid.String()       // UUID random
		ledgerproduct.OffChainHash = product.OffChainHash // hash for off chain database record

		productPrivateDetails := new(ProductPrivateDetails)
		productPrivateDetails.ObjectType = "PRIVATE_PRODUCT"
		productPrivateDetails.CombinedKey = product.CombinedKey
		productPrivateDetails.MaterialDescription = product.MaterialDescription
		productPrivateDetails.ManufacturingDate = product.ManufacturingDate
		productPrivateDetails.ShippingDate = product.ShippingDate
		productPrivateDetails.ShippingSiteID = product.ShippingSiteID

		productJSONasBytes, err := json.Marshal(ledgerproduct)
		if err != nil {
			_logger.Error(err.Error())
		}

		// === Save product to state ===
		err = ctx.GetStub().PutState(product.CombinedKey, productJSONasBytes)
		if err != nil {
			_logger.Error("failed to put Product: %s", err.Error())
		}

		bytes, _ := json.Marshal(productPrivateDetails)

		// === Save product to Private state ===
		err = ctx.GetStub().PutPrivateData(collectionName, product.CombinedKey, bytes)
		if err != nil {
			_logger.Error("failed to put Private Product: %s", err.Error())
		}
	}
	
	return "Success", nil
}


// CreateAsset creates a new instance of ProductContract
func (c *ProductContract) CreateOrUpdateBulkAssets(ctx contractapi.TransactionContextInterface) (string, error) {

	products := []Product{}
		
	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		_logger.Error("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &products)
	if err != nil {
		_logger.Error("failed to unmarshal JSON: %s", err.Error())
	}

	for i := 0; i < len(products); i += 1 {
		product := products[i]
		ledgerproduct := new(LedgerProduct)
		uid := uuid.New()

		ledgerproduct.ObjectType = "PRODUCT"
		ledgerproduct.CombinedKey = product.CombinedKey
		ledgerproduct.GTIN_SN = product.GTIN_SN
		ledgerproduct.BatchID = product.BatchID
		ledgerproduct.ExpiryDate = product.ExpiryDate
		ledgerproduct.ItemStatus = product.ItemStatus
		ledgerproduct.DispositionStatus = product.DispositionStatus
		ledgerproduct.ParentSequenceNo = product.ParentSequenceNo
		ledgerproduct.PartnerCountry = product.PartnerCountry
		ledgerproduct.LastUpdate = product.LastUpdate
		ledgerproduct.SysGenUniqueID = uid.String()       // UUID random
		ledgerproduct.OffChainHash = product.OffChainHash // hash for off chain database record
		ledgerproduct.Correction = true

		productPrivateDetails := new(ProductPrivateDetails)
		productPrivateDetails.ObjectType = "PRIVATE_PRODUCT"
		productPrivateDetails.CombinedKey = product.CombinedKey
		productPrivateDetails.MaterialDescription = product.MaterialDescription
		productPrivateDetails.ManufacturingDate = product.ManufacturingDate
		productPrivateDetails.ShippingDate = product.ShippingDate
		productPrivateDetails.ShippingSiteID = product.ShippingSiteID

		productJSONasBytes, err := json.Marshal(ledgerproduct)
		if err != nil {
			_logger.Error(err.Error())
		}

		// === Save product to state ===
		err = ctx.GetStub().PutState(product.CombinedKey, productJSONasBytes)
		if err != nil {
			_logger.Error("failed to put Product: %s", err.Error())
		}

		bytes, _ := json.Marshal(productPrivateDetails)

		// === Save product to Private state ===
		err = ctx.GetStub().PutPrivateData(collectionName, product.CombinedKey, bytes)
		if err != nil {
			_logger.Error("failed to put Private Product: %s", err.Error())
		}
	}
	
	return "Success", nil
}

// MyPrivateAssetExists returns true when asset with given ID exists in private data collection
func (c *ProductContract) MyPrivateAssetExists(ctx contractapi.TransactionContextInterface, myPrivateAssetID string) (bool, error) {

	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, myPrivateAssetID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// GetAsset retrieves an instance of Asset from the ladger
func (c *ProductContract) GetAsset(ctx contractapi.TransactionContextInterface) (*Product, error) { //no input parameter

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return nil, fmt.Errorf("Error getting transient: " + err.Error())
	}

	// Product properties are private, therefore they get passed in transient field
	transientDeleteJSON, ok := transMap["product_id"] //CombinedKey : asset_id?
	if !ok {
		return nil, fmt.Errorf("product to find not found in the transient map")
	}

	type productData struct {
		CombinedKey string `json:"combinedKey"`
	}

	myAsset := new(productData)
	err = json.Unmarshal(transientDeleteJSON, &myAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	bytes, err := ctx.GetStub().GetState(myAsset.CombinedKey)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	}
	myProductAsset := new(Product)

	err = json.Unmarshal(bytes, myProductAsset)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal private data collection data to type myAsset")
	}
	fmt.Println(myProductAsset, "line134")

	return myProductAsset, nil
}

// UpdateAsset update an existing instance of ProductContract
func (c *ProductContract) UpdateAsset(ctx contractapi.TransactionContextInterface) error {

	product := new(Product)

	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		return fmt.Errorf("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &product)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	bytes, err := ctx.GetStub().GetState(product.CombinedKey)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	}

	myLedgerAsset := new(Product)

	err = json.Unmarshal(bytes, myLedgerAsset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	myLedgerAsset.LastUpdate = product.LastUpdate
	myLedgerAsset.OffChainHash = product.OffChainHash
	if product.DispositionStatus == "SHIPPING" {
		myLedgerAsset.ShippingDate = product.ShippingDate
		myLedgerAsset.ShippingSiteID = product.ShippingSiteID
		myLedgerAsset.DispositionStatus = product.DispositionStatus
		// eventName = "SHIPPING"

	} else if product.DispositionStatus == "AGGREGATION" { // change matching data
		myLedgerAsset.ParentSequenceNo = product.ParentSequenceNo
		// eventName = "AGGREGATION"
	} else {
		return fmt.Errorf("Event type is not valid: %s", product.EventAction)
	}

	productJSONasBytes, err := json.Marshal(myLedgerAsset)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// === Save product to state ===
	// ctx.GetStub().SetEvent(eventName, productJSONasBytes)
	return ctx.GetStub().PutState(product.CombinedKey, productJSONasBytes)
}

// UpdateAsset update an existing instance of ProductContract
func (c *ProductContract) UpdateBulkAsset(ctx contractapi.TransactionContextInterface) (string, error) {

	products := []Product{}


	transientData, _ := ctx.GetStub().GetTransient()

	privateValue, exists := transientData["productValue"]

	if len(transientData) == 0 || !exists {
		return "failed", fmt.Errorf("The privateValue key was not specified in transient data. Please try again")
	}
	err := json.Unmarshal(privateValue, &products)
	if err != nil {
		return "failed",fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	for i := 0; i < len(products); i += 1 {
		product := products[i]
		bytes, err := ctx.GetStub().GetState(product.CombinedKey)
		if err != nil {
			return "failed", fmt.Errorf("Could not read from world state. %s", err)
		}

		myLedgerAsset := new(Product)

		err = json.Unmarshal(bytes, myLedgerAsset)
		if err != nil {
			return "failed", fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
		}

		myLedgerAsset.LastUpdate = product.LastUpdate
		myLedgerAsset.OffChainHash = product.OffChainHash
		if product.DispositionStatus == "SHIPPING" {
			myLedgerAsset.ShippingDate = product.ShippingDate
			myLedgerAsset.ShippingSiteID = product.ShippingSiteID
			myLedgerAsset.DispositionStatus = product.DispositionStatus
			// eventName = "SHIPPING"

		} else if product.DispositionStatus == "AGGREGATION" { // change matching data
			myLedgerAsset.ParentSequenceNo = product.ParentSequenceNo
			// eventName = "AGGREGATION"
		} else if product.DispositionStatus == "INACTIVE" { // change matching data
			myLedgerAsset.ItemStatus = product.ItemStatus
			// eventName = "AGGREGATION"
		} else {
			return "failed", fmt.Errorf("Event type is not valid: %s", product.EventAction)
		}

		productJSONasBytes, err := json.Marshal(myLedgerAsset)
		if err != nil {
			return "failed", fmt.Errorf(err.Error())
		}
		ctx.GetStub().PutState(product.CombinedKey, productJSONasBytes)
	}
	// === Save product to state ===
	// ctx.GetStub().SetEvent(eventName, productJSONasBytes)
	return "Success", nil
}

// GetPrivateAsset retrieves an instance of MyPrivateAsset from the private data collection

func (c *ProductContract) GetPrivateAsset(ctx contractapi.TransactionContextInterface) (*ProductPrivateDetails, error) {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return nil, fmt.Errorf("Error getting transient: " + err.Error())
	}

	// Product properties are private, therefore they get passed in transient field
	transientDeleteJSON, ok := transMap["product_id"]
	if !ok {
		return nil, fmt.Errorf("product to delete not found in the transient map")
	}

	type productData struct {
		CombinedKey string `json:"combinedKey"`
	}

	myAsset := new(productData)
	err = json.Unmarshal(transientDeleteJSON, &myAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	exists, err := c.MyPrivateAssetExists(ctx, myAsset.CombinedKey)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", myAsset.CombinedKey)
	}

	bytes, _ := ctx.GetStub().GetPrivateData(collectionName, myAsset.CombinedKey)

	myPrivateProductAsset := new(ProductPrivateDetails)

	err = json.Unmarshal(bytes, myPrivateProductAsset)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal private data collection data to type MyPrivateAsset")
	}

	return myPrivateProductAsset, nil
}

func (s *ProductContract) ReadPrivateAsset(ctx contractapi.TransactionContextInterface, assetID string) (*ProductPrivateDetails, error) {

	assetJSON, err := ctx.GetStub().GetPrivateData(collectionName, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}
	fmt.Println(string(assetJSON), "line224")

	//No Asset found, return empty response
	if assetJSON == nil {
		return nil, fmt.Errorf("%v does not exist in collection %v", assetID, collectionName)
	}

	asset := new(ProductPrivateDetails)
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	fmt.Println(asset, "line236")

	return asset, nil

}

func (c *ProductContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]Product, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"PRODUCT\"}}")

	queryResults, err := c.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (c *ProductContract) GetAllPrivateAssets(ctx contractapi.TransactionContextInterface) ([]ProductPrivateDetails, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"PRIVATE_PRODUCT\"}}")

	queryResults, err := c.getPrivateQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (c *ProductContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]Product, error) {

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Product{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newProduct := new(Product)

		err = json.Unmarshal(response.Value, newProduct)
		if err != nil {
			return nil, err
		}

		results = append(results, *newProduct)
	}
	return results, nil
}

func (c *ProductContract) getPrivateQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]ProductPrivateDetails, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []ProductPrivateDetails{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newProduct := new(ProductPrivateDetails)

		err = json.Unmarshal(response.Value, newProduct)
		if err != nil {
			return nil, err
		}

		results = append(results, *newProduct)
	}
	return results, nil
}

func (c *ProductContract) DeleteAsset(ctx contractapi.TransactionContextInterface) error {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}

	// Product properties are private, therefore they get passed in transient field
	transientDeleteProductJSON, ok := transMap["product_delete"]
	if !ok {
		return fmt.Errorf("product to delete not found in the transient map")
	}

	type productDelete struct {
		CombinedKey string `json:"combinedKey"`
	}

	var productDeleteInput productDelete
	err = json.Unmarshal(transientDeleteProductJSON, &productDeleteInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	if len(productDeleteInput.CombinedKey) == 0 {
		return fmt.Errorf("ID field must be a non-empty string")
	}

	valAsbytes, err := ctx.GetStub().GetState(productDeleteInput.CombinedKey) //get the Product from chaincode state
	if err != nil {
		return fmt.Errorf("failed to read product: %s", err.Error())
	}
	if valAsbytes == nil {
		return fmt.Errorf("product private details does not exist: %s", productDeleteInput.CombinedKey)
	}

	var productToDelete Product
	err = json.Unmarshal([]byte(valAsbytes), &productToDelete)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	// delete the product from state
	err = ctx.GetStub().DelState(productDeleteInput.CombinedKey)
	if err != nil {
		return fmt.Errorf("Failed to delete state:" + err.Error())
	}

	// Finally, delete private details of product
	err = ctx.GetStub().DelPrivateData(collectionName, productDeleteInput.CombinedKey)
	if err != nil {
		return err
	}

	return nil

}

// GetAssetHistory retrieves an instance of Asset from the ladger
// func (c *ProductContract) GetAssetHistory(ctx contractapi.TransactionContextInterface) ([]Product, error) { //no input parameter

// 	transMap, err := ctx.GetStub().GetTransient()
// 	if err != nil {
// 		return nil, fmt.Errorf("Error getting transient: " + err.Error())
// 	}

// 	// Product properties are private, therefore they get passed in transient field
// 	transientDeleteJSON, ok := transMap["product_id"] //CombinedKey : asset_id?
// 	if !ok {
// 		return nil, fmt.Errorf("product to find not found in the transient map")
// 	}

// 	type productData struct {
// 		CombinedKey string `json:"combinedKey"`
// 	}

// 	myAsset := new(productData)
// 	err = json.Unmarshal(transientDeleteJSON, &myAsset)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
// 	}

// 	products := []Product{}

// 	resultsIterator, err := ctx.GetStub().GetHistoryForKey(myAsset.CombinedKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("Could not read from history state. %s", err)
// 	}
// 	defer resultsIterator.Close()
// 	for resultsIterator.HasNext() {
// 		response, err := resultsIterator.Next()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return shim.Error(err.Error())
// 		}
// 		myProductAsset := new(Product)
// 		err = json.Unmarshal(bytes, myProductAsset)

// 		if err != nil {
// 			return nil, fmt.Errorf("Could not unmarshal history GetAssetHistory")
// 		}
// 		products = append(products, *newProduct)
// 	}

// 	fmt.Println(products, "line681")

// 	return products, nil
// }