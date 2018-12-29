package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

// Define Asset

// Define Chaincode
type SampleChaincode struct {}

/*

Method
add data into ledger, send event to blockchain system

*/
func (sc *SampleChaincode) setData(stub shim.ChaincodeStubInterface, args []string ) peer.Response{

	/*
		當收到資料的時候，發出customEvent，將更新內容也一並通知
	*/
	stub.SetEvent("dataUpdate",[]byte("just test"))
	stub.PutState(args[0],[]byte(args[1]))
	return shim.Success([]byte("update"))

}

func (sc *SampleChaincode) getData(stub shim.ChaincodeStubInterface, args []string ) peer.Response{

	/*
		當收到資料的時候，發出customEvent，將更新內容也一並通知
	*/
	stub.SetEvent("getData",[]byte("get data test"))
	res,_ := stub.GetState(args[0])
	return shim.Success(res)

}

func (sc *SampleChaincode) Init( stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)

}

func (sc *SampleChaincode) Invoke( stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "set" {
		return sc.setData(stub, args)
	}else if function =="get" {
		return sc.getData(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")

}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}

}
