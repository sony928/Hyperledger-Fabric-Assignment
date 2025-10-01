package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Account struct {
    ID      string `json:"id"`
    Balance int    `json:"balance"`
}

type SmartContract struct {
    contractapi.Contract
}

func (s *SmartContract) CreateAccount(ctx contractapi.TransactionContextInterface, id string, balance int) error {
    account := Account{
        ID:      id,
        Balance: balance,
    }
    accountBytes, _ := json.Marshal(account)
    return ctx.GetStub().PutState(id, accountBytes)
}

func (s *SmartContract) GetAccount(ctx contractapi.TransactionContextInterface, id string) (*Account, error) {
    accountBytes, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, err
    }
    if accountBytes == nil {
        return nil, fmt.Errorf("account %s does not exist", id)
    }
    var account Account
    json.Unmarshal(accountBytes, &account)
    return &account, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        panic(err)
    }
    if err := chaincode.Start(); err != nil {
        panic(err)
    }
}

