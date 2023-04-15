package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

// SmartContract provides functions for managing shares
type StockTxContract struct {
	contractapi.Contract
}

var logger = flogging.MustGetLogger("stocktx_cc")

type StockTransaction struct {
	ID        string  `json:"id"`
	TradeDate string  `json:"trade_date"`
	Buyer     string  `json:"buyer"`
	Seller    string  `json:"seller"`
	StockCode string  `json:"stock_code"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// CreateShare adds a new share transaction to the ledger
func (s *StockTxContract) CreateTx(ctx contractapi.TransactionContextInterface, stockTxDate string) (string, error) {
	if len(stockTxDate) == 0 {
		return "", fmt.Errorf("Please pass the correct stock transaction data")
	}

	var stockTransaction StockTransaction
	err := json.Unmarshal([]byte(stockTxDate), &stockTransaction)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling stock transaction. %s", err.Error())
	}

	stockTransactionAsBytes, err := json.Marshal(stockTransaction)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling stock transaction. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(stockTransaction.ID, stockTransactionAsBytes)
}

func (s *StockTxContract) GetStockTxById(ctx contractapi.TransactionContextInterface, stockTransactionID string) (*StockTransaction, error) {
	if len(stockTransactionID) == 0 {
		return nil, fmt.Errorf("Please provide correct  stock transaction data Id")
		// return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	stockTransactionAsBytes, err := ctx.GetStub().GetState(stockTransactionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if stockTransactionAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", stockTransactionID)
	}

	stockTransaction := new(StockTransaction)
	_ = json.Unmarshal(stockTransactionAsBytes, stockTransaction)

	return stockTransaction, nil

}

func main() {

	chaincode, err := contractapi.NewChaincode(new(StockTxContract))
	if err != nil {
		fmt.Printf("Error create stocktx chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincodes: %s", err.Error())
	}

}
