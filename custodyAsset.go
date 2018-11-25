package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type CustodyAsset struct {
}

// Define the asset structure, with initial set of properties.  Structure tags are used by encoding/json library
type EquityAsset struct {
	investorName        string `json:"investorName"`
	investorPAN         string `json:"investorPAN"`
	equityName          string `json:"equityName"`
	equityNos           string `json:"equityNos"`
}

/*
 * The Init method is called when the Smart Contract "CustodyAsset" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *CustodyAsset) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "CustodyAsset"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *CustodyAsset) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "query" {
		return s.queryEquity(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "create" {
		return s.createEquity(APIstub, args)
	} else if function == "queryAll" {
		return s.queryAllEquities(APIstub)
	} else if function == "change" {
		return s.changeEquity(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *CustodyAsset) queryEquity(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	investorAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(investorAsBytes)
}

func (s *CustodyAsset) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	equityAssets := []EquityAsset{
		EquityAsset{investorName: "John Doe", investorPAN: "AAAPH1111H", equityName:"BHEL", equityNos:"20"},
		EquityAsset{investorName: "Mary Doe", investorPAN: "BBBPH1111H", equityName:"Infosys", equityNos:"30"},
	}

	i := 0
	for i < len(equityAssets) {
		fmt.Println("i is ", i)
		equityAsBytes, _ := json.Marshal(equityAssets[i])
		APIstub.PutState("EQUITY "+strconv.Itoa(i), equityAsBytes)
		fmt.Println("Added", equityAssets[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *CustodyAsset) createEquity(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var equityAsset = EquityAsset{investorName: args[1], investorPAN: args[2], equityName: args[3], equityNos: args[4]}

	equityAsBytes, _ := json.Marshal(equityAsset)
	APIstub.PutState(args[0], equityAsBytes)

	return shim.Success(nil)
}

func (s *CustodyAsset) queryAllEquities(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "EQUITY 0"
	endKey := "EQUITY 999"

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

	fmt.Printf("- queryAllEquities:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *CustodyAsset) changeEquity(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	equityAsBytes, _ := APIstub.GetState(args[0])
	equity := EquityAsset{}

	json.Unmarshal(equityAsBytes, &equity)
	equity.investorName = args[1]

	equityAsBytes, _ = json.Marshal(equity)
	APIstub.PutState(args[0], equityAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(CustodyAsset))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

