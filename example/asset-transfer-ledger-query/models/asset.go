package models

import (
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/config"
	"github.com/onespacegolib/fabricv2-cckit/flannel"
	"github.com/onespacegolib/fabricv2-cckit/router"
)

type Asset struct {
	flannel.Schema
	ID             string `json:"id" flannel:"primary_key"` //the field tags are needed to keep case from bouncing around
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraised_value"`
}

func AssetModel(c router.Context) flannel.Flannel {
	inti := &flannel.InitStubInterface{
		ModelOrg:  config.FabricORG,
		ModelType: flannel.TypeAsset,
	}
	f := flannel.InitFlannel(c, inti).Model(&Asset{})

	return f
}

func AssetChainValue(c router.Context) flannel.ChainCompositeKey {
	return AssetModel(c).InitChainCompositeKey(`asset_id`, `color`, `owner`)
}
