package contract

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Charitycontract contract for managing CRUD for Car
type CharityContract struct {
	contractapi.Contract
}

type Charity struct {
	AssetType string `json:"assetType"`
	CharityId string `json:"charityID"`
	Amount    string `json:"amount"`
	Cause     string `json:"cause"`
}

type HistoryQueryResult struct {
	Record    *Charity `json:"record"`
	TxId      string   `json:"txId"`
	Timestamp string   `json:"timestamp"`
	IsDelete  bool     `json:"isDelete"`
}

type PaginatedQueryResult struct {
	Records             []*Charity `json:"records"`
	FetchedRecordsCount int32      `json:"fetchedRecordsCount"`
	Bookmark            string     `json:"bookmark"`
}

func (c *CharityContract) CarExists(ctx contractapi.TransactionContextInterface, charityID string) (bool, error) {
	data, err := ctx.GetStub().GetState(charityID)

	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)

	}
	return data != nil, nil
}

func (c *CharityContract) CreateCharity(ctx contractapi.TransactionContextInterface, charityID string, amount string, cause string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	// if clientOrgID == "manufacturer-auto-com" {
	// if clientOrgID == "Org1MSP" {
	if clientOrgID == "CharityorgMSP" {
		exists, err := c.CarExists(ctx, charityID)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if exists {
			return "", fmt.Errorf("the car, %s already exists", charityID)
		}

		charity := Charity{
			AssetType: "charity",
			CharityId: charityID,
			Amount:    amount,
			Cause:     cause,
		}

		bytes, _ := json.Marshal(charity)

		err = ctx.GetStub().PutState(charityID, bytes)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("successfully added car %v", charityID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

// ReadCar retrieves an instance of Charity from the world state
func (c *CharityContract) ReadCharity(ctx contractapi.TransactionContextInterface, charityID string) (*Charity, error) {

	bytes, err := ctx.GetStub().GetState(charityID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("the car %s does not exist", charityID)
	}

	var charity Charity

	err = json.Unmarshal(bytes, &charity)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type charity")
	}

	return &charity, nil
}

func (c *CharityContract) DeleteCharity(ctx contractapi.TransactionContextInterface, charityID string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}
	if clientOrgID == "CharityorgMSP" {

		exists, err := c.CarExists(ctx, charityID)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if !exists {
			return "", fmt.Errorf("the charity, %s does not exist", charityID)
		}

		err = ctx.GetStub().DelState(charityID)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("charity with id %v is deleted from the world state.", charityID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

func charityResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Charity, error) {
	var charities []*Charity
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of the result iterator. %s", err)
		}
		var charity Charity
		err = json.Unmarshal(queryResult.Value, &charity)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		charities = append(charities, &charity)
	}

	return charities, nil
}

func (c *CharityContract) GetAllCharities(ctx contractapi.TransactionContextInterface) ([]*Charity, error) {

	queryString := `{"selector":{"assetType":"charity"}}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return charityResultIteratorFunction(resultsIterator)
}

func (c *CharityContract) GetCharitiesByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Charity, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the  data by range. %s", err)
	}
	defer resultsIterator.Close()

	return charityResultIteratorFunction(resultsIterator)
}

func (c *CharityContract) GetCharityHistory(ctx contractapi.TransactionContextInterface, charityID string) ([]*HistoryQueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(charityID)
	if err != nil {
		return nil, fmt.Errorf("could not get the data. %s", err)
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not get the value of resultsIterator. %s", err)
		}

		var charity Charity
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &charity)
			if err != nil {
				return nil, err
			}
		} else {
			charity = Charity{
				CharityId: charityID,
			}
		}

		timestamp := response.Timestamp.AsTime()

		formattedTime := timestamp.Format(time.RFC1123)

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &charity,
			IsDelete:  response.IsDelete,
		}
		records = append(records, &record)
	}

	return records, nil
}

func (c *CharityContract) GetCharityWithPagination(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	queryString := `{"selector":{"assetType":"charity"}}`

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("could not get the charity records. %s", err)
	}
	defer resultsIterator.Close()

	charities, err := charityResultIteratorFunction(resultsIterator)
	if err != nil {
		return nil, fmt.Errorf("could not return the charity records %s", err)
	}

	return &PaginatedQueryResult{
		Records:             charities,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}