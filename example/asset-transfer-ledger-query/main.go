package main

import (
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/example/asset-transfer-ledger-query/routers"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	cc := routers.New()
	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting Contract chaincode: %s", err)
	}
}
