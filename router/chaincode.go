package router

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type Chaincode struct {
	router *Group
}

func NewChaincode(r *Group) *Chaincode {
	return &Chaincode{r}
}

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return cc.router.HandleInit(stub)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return cc.router.Handle(stub)
}