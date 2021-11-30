package tests

import (
	"git.onespace.co.th/osgolib/fabricv2-cckit/example/asset-transfer-ledger-query/config"
	"git.onespace.co.th/osgolib/fabricv2-cckit/example/asset-transfer-ledger-query/routers"
	"git.onespace.co.th/osgolib/fabricv2-cckit/testcc"
	"github.com/hyperledger/fabric-protos-go/peer"
)

var (
	cc  = testcc.NewMockStub(config.ChaincodeName, routers.New())
	res peer.Response
)
