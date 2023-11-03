**SCAN CHAINCODE**

**HOW THE DATA IS STORED IN BLOCKCHAIN AND TRACKING DATA HISTORY** :

This is a quick explanation of how data is stored in a blockchain, including the functions used to call and store bulk data, how the data is sent to the blockchain, the function used to update the data, and how to track the data's history.

You can store bulk data as a JSON array on the blockchain using the CreateBulkScannedData function. You can update the data with the UpdateScannedData function and track its history with the GetScanHistory function. Additionally, you can track the updated data with the GetScannedData function.

**CODE EXPLANATION:**

Four functions are there in the chaincode,

1)CreateBulkScannedData

2)UpdateScannedData

3)GetScannedData

4)GetScanHistory

**CreateBulkScannedData:**

This function will only be used by the MSD organisation. It is designed to store bulk data in DLT, and takes input in the form of a JSON array here is an example,

e.g:

[

{

"combinedKey": "ca87b167277e6ecce0fa69f5ca1018f463210b227bd32e913a00d28e689f87b1",

"gtinsn": "046022100020781000220077030",

"serialID": "1000220077030",

"batchId": "EXT3PLTST11",

"materialName": "TestMaterialname",

"materialDescription": "undefined",

"manufacturingDate": "2021-02-08",

"expiryDate": "2026-02-08",

"itemStatus": "active",

"timemillies": "45454345667",

"userId": "MSDuser",

"deviceUuid": "75C84E94-9E79-45BF-B84D-D9852F7CFD3F",

"scannedResult": "",

"scanLocation": "NewJersey,US",

"shipToParty": "Distributor1user",

}

]

**UpdateScannedData:**

This function will be used by the Distributor organisation. It is designed to store the scanned events at distributors, sub-distributors, and hospital users' data in DLT, and it takes input in the form of a JSON array here is an example,

e.g:

[

{

"combinedKey": "ca87b167277e6ecce0fa69f5ca1018f463210b227bd32e913a00d28e689f87b1",

"gtinsn": "046022100020781000220077030",

"batchId": "EXT3PLTST11",

"expiryDate": "2026-02-08",

"timemillies": "55454345667",

"userId": "Distributor1user",

"deviceUuid": "75C84E94-9E79-45BF-B84D-D9852F7CFD3F",

"scanLocation": "Jalico,mexico",

}

]

During this update process it will go through the following validation,

- It will validate the expiry date.
- Ship to party validation
- Check the rule for chronological scanned events (MSD -\> Distributor1 ).

**GetScannedData:**

This function will be used to retrieve the latest scanned data. It takes a combined key as input and can be used by both MSD and Distributor org.

**GetScanHistory:**

This function will be used to retrieve the history of a particular scanned data. It takes a combined key as input and can be used by both MSD and Distributor org.

**Function for report generation:**

1)GetByUserId

2)GetByDates

3)GetByScannedResults

4)GetByDatesWithScannedLocation

5)GetByDatesWithScannedResults

6)GetByDatesWithScannedResultsandScannedLocation

7)GetBySerialId

8)GetByUserIdAndSerialId