package asset_controller

import (
	"github.com/onespacegolib/fabricv2-cckit/example/asset-transfer-ledger-query/models"
	"github.com/onespacegolib/fabricv2-cckit/router"
)

func Put(c router.Context) (interface{}, error) {
	//var asset models.Asset
	//assetModel = models.AssetModel(c)
	//if err := models.AssetModel(c).FindByID(`111`, &asset).Error(); err != nil {
	//	return nil, err
	//}
	//asset.Owner = `tom2`
	//if err := models.AssetModel(c).UpdateOne(&asset).Error(); err != nil {
	//	return nil, err
	//}

	if err := models.AssetModel(c).UpdateOne(&models.Asset{ID: `111`, Owner: `tom2`}).Error(); err != nil {
		return nil, err
	}

	//if err := models.AssetModel(c).FindByIdAndUpdate(`111`, &models.Asset{Owner: `tom2`, Size: 8}).Error(); err != nil {
	//	return nil, err
	//}
	return nil, nil
}
