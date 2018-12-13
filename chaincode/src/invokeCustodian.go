package main

import (
    //"fmt"
    "encoding/json"
    //"strings"
    //"time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO CREATE INVESTOR
func onboardInvestor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    logger.Info("***************************************\n")
    logger.Info("---------- IN ONBOARDINVESTOR----------\n")
    logger.Info("***************************************\n")

    // RETURN ERROR IF ARGS IS NOT 7 IN NUMBER
    if len(args) != 7 {
        return shim.Error("Invalid argument count. Expecting 7.")
    }
/**
    // CREATE A TEMP STRUCTURE TO RECEIVE INVESOR DATA FROM API
    dto := struct {
        userName     string  `json:"user_name"`
        userFName    string  `json:"user_fname"`
        userLName    string  `json:"user_lname"`
        userIdentity string  `json:"user_identity"`
        kycStatus    string  `json:"kyc_status"`
        depositoryAC string  `json:"depository_ac"`
        bankAC       string  `json:"bank_ac"`
    } {}

    // CHECK FOR ERROR IN PARSING INPUT
    err := json.Unmarshal([]byte(args[0]), &dto)
    if err != nil {
        return shim.Error(err.Error())
    }
**/
    // MOVE THE INPUT ARGS INTO A STRUCTURE
    // PREPARE THE INPUT VALUES TO WRITE
    _investor := investor {
        userName:     args[0], //dto.userName,
        userFName:    args[1], //dto.userFName,
        userLName:    args[2], //dto.userLName,
        userIdentity: args[3], //dto.userIdentity,
        kycStatus:    args[4], //dto.kycStatus,
        depositoryAC: args[5], //dto.depositoryAC,
        bankAC:       args[6], //dto.bankAC,
    }

    // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
    _investorKey, err := stub.CreateCompositeKey(prefixCustodian, []string{_investor.userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }

    // MARSHAL INVESTOR RECORD
    _investorBytes, err := json.Marshal(_investor)
    // CHECK FOR ERROR IN MARSHALING 
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE INVESTOR RECORD
    err = stub.PutState(_investorKey, _investorBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }

    logger.Info("****************************************\n")
    logger.Info("---------- OUT ONBOARDINVESTOR----------\n")
    logger.Info("****************************************\n")

    // RETURN SUCCESS
    return shim.Success(nil)
}

