package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/routers"
)

func main() {
	cc := routers.New()
	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting Contract chaincode: %s", err)
	}
}
