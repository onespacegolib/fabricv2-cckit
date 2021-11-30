package asset_controller

import (
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/example/asset-transfer-ledger-query/models"
	"git.onespace.co.th/osgolib/fabricv2-cckit/flannel"
	"git.onespace.co.th/osgolib/fabricv2-cckit/router"
	"log"
)

func GetID(c router.Context) (interface{}, error) {

	var asset models.Asset

	if err := models.AssetModel(c).
		FindById(`111`, &asset).
		Error(); err != nil {
		log.Println(err)
		if !flannel.IsRecordNotFoundError(err) {
			return nil, err
		} else {
			return nil, fmt.Errorf(`ไม่สามารถ`)
		}
	}

	return asset, nil
}
