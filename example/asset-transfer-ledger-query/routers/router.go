package routers

import (
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/config"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/controllers/asset_controller"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/controllers/init_controller"
	"github.com/onespacegolib/fabricv2-cckit/flannel"
	"github.com/onespacegolib/fabricv2-cckit/router"
)

var r = router.New(config.ChaincodeName)

func New() *router.Chaincode {

	r.Init(init_controller.Init)

	flannel.Router(r)

	r.Group(`asset`).
		// สำหรับลงทะเบียน Service ID
		Invoke(`create`, asset_controller.Create, nil).
		Query(`get[id]`, asset_controller.GetID, nil).
		Invoke(`put`, asset_controller.Put, nil)

	return router.NewChaincode(r)
}
