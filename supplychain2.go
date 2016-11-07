/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init callback representing the init of a chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {

	_, args := stub.GetFunctionAndParameters()

	var shipmentID, shipmentDocuments string // Entities
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	// Initialize the chaincode
	shipmentID = args[0]
	shipmentDocuments = args[1]
	fmt.Printf("shipmentID = %s, shipmentDocuments = %s\n", shipmentID, shipmentDocuments)

	// Write the state to the ledger
	err = stub.PutState(shipmentID, []byte(shipmentDocuments))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke callback representing the query of a chaincode
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	_, args := stub.GetFunctionAndParameters()

	var shipmentID, shipmentDocuments string // Entities
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	shipmentID = args[0]
	shipmentDocuments = args[1]

	fmt.Printf("Shipment ID submitted with shipment documents.  shipmentID = %s, shipmentDocuments = %s\n", shipmentID, shipmentDocuments)

	err = stub.PutState(shipmentID, []byte(shipmentDocuments))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
