package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type serverConfig struct {
    CCID string
    Address string
}
func main() {
	productContract := new(ProductContract)
	productContract.Info.Version = "0.0.1"
	productContract.Info.Description = "Merck Product Smart Contract -- New Chaincode Upgrade build 17"
	productContract.Info.License = new(metadata.LicenseMetadata)
	productContract.Info.License.Name = "Apache-2.0"
	productContract.Info.Contact = new(metadata.ContactMetadata)
	productContract.Info.Contact.Name = "Aditya Saha"

	config := serverConfig{
	    CCID: os.Getenv("CHAINCODE_CCID"),
	    Address: os.Getenv("CHAINCODE_ADDRESS"),
	}

	chaincode, err := contractapi.NewChaincode(productContract)
	chaincode.Info.Title = "merck product chaincode"
	chaincode.Info.Version = os.Getenv("CHAINCODE_CCID")

	if err != nil {
		panic("Could not create chaincode from ProductContract." + err.Error())
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: getTLSProperties(),
	}

	fmt.Printf("CCID: %s\n", config.CCID)
	fmt.Printf("Addr: %s\n", config.Address)
	fmt.Printf("%+v\n", productContract.Info.Description)
	fmt.Printf("tls: %s\n",getEnvOrDefault("TLS_DISABLED", "true"))
	fmt.Printf("%v\n", chaincode.Info)

	if err := server.Start(); err != nil {
		log.Panicf("error starting asset-transfer-basic chaincode: %s", err)
	}else{
		fmt.Printf("Productchaincode started listening on %s\n",config.Address )
	}
}

func getTLSProperties() shim.TLSProperties {
	// Check if chaincode is TLS enabled
	tlsDisabledStr := getEnvOrDefault("TLS_DISABLED", "true")
	key := getEnvOrDefault("TLS_KEY_FILE", "")
	cert := getEnvOrDefault("TLS_CERT_FILE", "")
	clientCACert := getEnvOrDefault("TLS_ROOTCERT_FILE", "")

	// convert tlsDisabledStr to boolean
	tlsDisabled := getBoolOrDefault(tlsDisabledStr, false)
	var keyBytes, certBytes, clientCACertBytes []byte
	var err error

	if !tlsDisabled {
		keyBytes, err = ioutil.ReadFile(key)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
		certBytes, err = ioutil.ReadFile(cert)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
	}
	// Did not request for the peer cert verification
	if clientCACert != "" {
		clientCACertBytes, err = ioutil.ReadFile(clientCACert)
		if err != nil {
			log.Panicf("error while reading the crypto file: %s", err)
		}
	}

	return shim.TLSProperties{
		Disabled: tlsDisabled,
		Key: keyBytes,
		Cert: certBytes,
		ClientCACerts: clientCACertBytes,
	}
}

func getEnvOrDefault(env, defaultVal string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		value = defaultVal
	}
	return value
}

// Note that the method returns default value if the string
// cannot be parsed!
func getBoolOrDefault(value string, defaultVal bool) bool {
	parsed, err := strconv.ParseBool(value)
	if err!= nil {
		return defaultVal
	}
	return parsed
}