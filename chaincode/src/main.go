package main

import (
    //"fmt"
    //"encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

const prefixCustodian = "CUSTDN"
//const prefixBank = "BNK"
//const prefixExchange = "XCHNG"
//const prefixDepository = "DPSTRY"

var logger = shim.NewLogger("main")

type SmartContract struct {
}

// MAPPING BETWEEN FUNCTION NAMES IN APIs and GO METHODS
var bcFunctions = map[string] func(shim.ChaincodeStubInterface, []string) pb.Response{
    // CUSTODIAN PEER
    "onboard_investor":       onboardInvestor,
    //"check_kyc":              checkKYC,
    //"buy_share":              buyShare,
    //"sell_share":             sellShare,
    //"get_investor_dashboard": getInvestorDashboards,

    // BANK PEER
    //"execute_transaction": executeTransaction,

    // EXCHANGE PEER
    //"init_exchange": initExchange,
    //"execute_trade": executeTrade,

    // DEPOSITORY PEER
    //"init_depository": initDepository,
    //"record_trade":    recordTrade,
}

// INIT CALLBACK REPRESENTING INVOCATION OF CHAINCODE
// INITIALIZE EXCHANGE STRUCTURE WITH MASTER DATA
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
    //_, args := stub.GetFunctionAndParameters()
    //t.initExchange(stub)
    logger.Info("**********************************\n")
    logger.Info("----------IN INIT METHOD----------\n")
    logger.Info("**********************************\n")
    return shim.Success(nil)
}

// INVOKE FUNCTION ACCEPS BLOCKCHAIN CODE INVOCATIONS
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    logger.Info("************************************\n")
    logger.Info("----------IN INVOKE METHOD----------\n")
    logger.Info("************************************\n")
    function, args := stub.GetFunctionAndParameters()
    if function == "init" {
        return t.Init(stub)
    }
    bcFunc := bcFunctions[function]
    if bcFunc == nil {
        return shim.Error("Invalid invoke function.")
        }
    return bcFunc(stub, args)
}

// MAIN METHOD
func main() {
    logger.SetLevel(shim.LogInfo)
    err := shim.Start(new(SmartContract))
    logger.Info("**********************************\n")
    logger.Info("----------In MAIN METHOD----------\n")
    logger.Info("**********************************\n")
    if err != nil {
        logger.Info("Error starting Simple chaincode: %s", err)
    }
}
