package asset_controller

import (
	"github.com/onespacegolib/fabricv2-cckit/constant"
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/models"
	"github.com/onespacegolib/fabricv2-cckit/flannel"
	"github.com/onespacegolib/fabricv2-cckit/router"
	"net/http"
)

func Create(c router.Context) (interface{}, error) {
	asset := models.Asset{
		ID:             string(c.GetArgs()[1]),
		Color:          "blue",
		Size:           5,
		Owner:          "tom",
		AppraisedValue: 35,
	}

	asset2 := models.Asset{
		ID:             "asset2",
		Color:          "blue",
		Size:           5,
		Owner:          "tom",
		AppraisedValue: 35,
	}

	if err := models.AssetModel(c).Create(&asset).Error(); err != nil {
		return nil, err
	}
	if err := models.AssetModel(c).Create(&asset2).Error(); err != nil {
		return nil, err
	}
	if err := models.AssetChainValue(c).Create(asset.ID, asset.Color, asset.Owner).Error(); err != nil {
		return nil, err
	}
	if err := models.AssetChainValue(c).Create(asset2.ID, asset2.Color, asset2.Owner).Error(); err != nil {
		return nil, err
	}

	var assetRes flannel.ChainList
	if err := models.AssetChainValue(c).Find([]string{}, &assetRes).Error(); err != nil {
		return nil, err
	}

	return c.JSON(constant.ApiResult{
		Result:      assetRes,
		Message:     ``,
		MessageCode: http.StatusBadRequest,
	})
}
