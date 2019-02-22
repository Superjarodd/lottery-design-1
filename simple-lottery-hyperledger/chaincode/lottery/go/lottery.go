package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	sc "github.com/hyperledger/fabric/protos/peer"
)

var n = 1

// SmartContract ...
type SmartContract struct {
}

type Lottery struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createLottery" {
		return s.createLottery(APIstub, args)
	} else if function == "queryAllLotteries" {
		return s.queryAllLotteries(APIstub)
	} else if function == "pickWinner" {
		return s.pickWinner(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	i := 0

	tickets := []Lottery{
		Lottery{Number: strconv.Itoa(i), Name: "Lenovo"},
	}

	for i < len(tickets) {
		fmt.Println("i is ", i)
		ticketAsBytes, _ := json.Marshal(tickets[i])
		APIstub.PutState(strconv.Itoa(i), ticketAsBytes)
		fmt.Println("Added", tickets[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createLottery(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var tickets = Lottery{Number: strconv.Itoa(n), Name: args[0]}

	ticketAsBytes, _ := json.Marshal(tickets)
	APIstub.PutState(strconv.Itoa(n), ticketAsBytes)
	n++

	return shim.Success(nil)
}

func (s *SmartContract) queryAllLotteries(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllLotteries:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) pickWinner(APIstub shim.ChaincodeStubInterface) sc.Response {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := n + 1
	result := rand.Intn(max-min) + min

	queryString := fmt.Sprintf("{\"selector\":{\"number\":\"%d\"}}", result)
	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Winning number\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Winner's name\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
