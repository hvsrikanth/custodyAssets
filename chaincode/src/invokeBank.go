package main

import (
    "fmt"
    "encoding/json"
    //"strings"
    //"time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


// METHOD TO INITIATE RECORDS IN BANK MASTER
func initBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("******************************")
    fmt.Println("---------- INIT BANK----------")
    fmt.Println("******************************")

    // INITIALIZE BANK MASTER WITH DATA
    _bankMaster := []bankMaster {
                     bankMaster {userName: "hvsrikanth",bankAC: "HDFC00001", balance:10000},
                     bankMaster {userName: "madhurib",  bankAC: "HDFC00002", balance:20000},
                     bankMaster {userName: "naveens",   bankAC: "HDFC00003", balance:30000},
                    }
    i := 0
    for (i<len(_bankMaster)) {

        fmt.Println("i is : ", i)

        // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
        _bankKey, err := stub.CreateCompositeKey(prefixBank, []string{_bankMaster[i].userName})
        // CHECK FOR ERROR IN CREATING COMPOSITE KEY
        if err != nil {
            return shim.Error(err.Error())
        }

        // MARSHAL THE BANK MASTER RECORD 
        _bankMasterAsBytes, err := json.Marshal(_bankMaster[i])
        // CHECK FOR ERROR IN MARSHALING
        if err != nil {
            return shim.Error(err.Error())
        }

        // NOW WRITE THE BANK MASTER RECORD
        err = stub.PutState(_bankKey, _bankMasterAsBytes)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }
    }

    fmt.Println("***********************************")
    fmt.Println("---------- OUT INIT BANK ----------")
    fmt.Println("***********************************")

    // RETURN SUCCESS
    return shim.Success(nil)
}

