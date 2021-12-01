package tests

import (
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/config"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/routers"
	"github.com/onespacegolib/fabricv2-cckit/testcc"
)

var (
	cc  = testcc.NewMockStub(config.ChaincodeName, routers.New())
	res peer.Response
)
