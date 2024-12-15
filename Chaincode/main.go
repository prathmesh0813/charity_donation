package main

import (
	"charity/contract"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	charityContract := new(contract.CharityContract)
	donarContract := new(contract.DonarContract)

	chaincode, err := contractapi.NewChaincode(charityContract, donarContract)

	if err != nil {
		log.Panicf("Could not create chaincode : %v", err)
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode : %v", err)
	}

}
