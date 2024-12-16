package contract

// first write collection file

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DonarContract struct {
	contractapi.Contract
}

type Donar struct {
	AssetType     string `json:"donar"`
	Name          string `json:"name"`
	DonarID       string `json:"donarID"`
	CharityID     string `json:"charityID"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

const collectionName string = "DonarCollection"

// OrderExists returns true when asset with given ID exists in private data collection
func (d *DonarContract) DonarExists(ctx contractapi.TransactionContextInterface, donarID string) (bool, error) {

	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, donarID)

	if err != nil {
		return false, fmt.Errorf("could not fetch the private data hash. %s", err)
	}

	return data != nil, nil
}

// CreateOrder creates a new instance of Order
func (d *DonarContract) CreateDonar(ctx contractapi.TransactionContextInterface, donarID string) (string, error) {

	fmt.Println("--------------INSIDE CREATE--------------------")
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "DonarMSP" {
		exists, err := d.DonarExists(ctx, donarID)
		if err != nil {
			return "", fmt.Errorf("could not read from world state. %s", err)
		} else if exists {
			return "", fmt.Errorf("the asset %s already exists", donarID)
		}

		var donar Donar

		transientData, err := ctx.GetStub().GetTransient()
		if err != nil {
			return "", fmt.Errorf("could not fetch transient data. %s", err)
		}

		if len(transientData) == 0 {
			return "", fmt.Errorf("please provide the private data of make, model, color, dealerName")
		}

		name, exists := transientData["name"]
		if !exists {
			return "", fmt.Errorf("the name was not specified in transient data. Please try again")
		}
		donar.Name = string(name)

		charityID, exists := transientData["charityID"]
		if !exists {
			return "", fmt.Errorf("the charityid was not specified in transient data. Please try again")
		}
		donar.CharityID = string(charityID)

		amount, exists := transientData["amount"]
		if !exists {
			return "", fmt.Errorf("the amount was not specified in transient data. Please try again")
		}
		donar.Amount = string(amount)

		txnsID, exists := transientData["txnsID"]
		if !exists {
			return "", fmt.Errorf("the txnsID was not specified in transient data. Please try again")
		}
		donar.TransactionID = string(txnsID)

		status, exists := transientData["status"]
		if !exists {
			return "", fmt.Errorf("the status was not specified in transient data. Please try again")
		}
		donar.Status = string(status)

		donar.AssetType = "Donar"
		//donar.Status = "Pending"
		donar.DonarID = donarID

		bytes, _ := json.Marshal(donar)
		err = ctx.GetStub().PutPrivateData(collectionName, donarID, bytes)
		if err != nil {
			return "", fmt.Errorf("could not able to write the data")
		}
		return fmt.Sprintf("donar with id %v added successfully", donarID), nil
	} else {
		return fmt.Sprintf("donar cannot be created by organisation with MSPID %v ", clientOrgID), nil
	}
}

// ReadOrder retrieves an instance of Order from the private data collection
func (d *DonarContract) ReadDonar(ctx contractapi.TransactionContextInterface, orderID string) (*Donar, error) {

	fmt.Println("--------------INSIDE Read--------------------")
	exists, err := d.DonarExists(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", orderID)
	}

	bytes, err := ctx.GetStub().GetPrivateData(collectionName, orderID)
	if err != nil {
		return nil, fmt.Errorf("could not get the private data. %s", err)
	}
	var donar Donar

	err = json.Unmarshal(bytes, &donar)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal private data collection data to type Donar")
	}

	return &donar, nil

}

// DeleteOrder deletes an instance of Order from the private data collection
func (d *DonarContract) DeleteDonar(ctx contractapi.TransactionContextInterface, donarID string) error {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("could not read the client identity. %s", err)
	}

	if clientOrgID == "DonarCollection" {

		exists, err := d.DonarExists(ctx, donarID)

		if err != nil {
			return fmt.Errorf("could not read from world state. %s", err)
		} else if !exists {
			return fmt.Errorf("the asset %s does not exist", donarID)
		}

		return ctx.GetStub().DelPrivateData(collectionName, donarID)
	} else {
		return fmt.Errorf("organisation with %v cannot delete the donar", clientOrgID)
	}
}

func (d *DonarContract) GetAllDonars(ctx contractapi.TransactionContextInterface) ([]*Donar, error) {
	queryString := `{"selector":{"assetType":"Donar"}}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return DonarResultIteratorFunction(resultsIterator)
}

func (d *DonarContract) GetDonarsByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]*Donar, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(collectionName, startKey, endKey)

	if err != nil {
		return nil, fmt.Errorf("could not fetch the private data by range. %s", err)
	}
	defer resultsIterator.Close()

	return DonarResultIteratorFunction(resultsIterator)

}

// iterator function

func DonarResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Donar, error) {
	var donars []*Donar
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of result iterator. %s", err)
		}
		var donar Donar
		err = json.Unmarshal(queryResult.Value, &donar)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		donars = append(donars, &donar)
	}

	return donars, nil
}
