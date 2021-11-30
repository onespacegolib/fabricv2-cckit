package flannel

import (
	"encoding/json"
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/constant"
	"git.onespace.co.th/osgolib/fabricv2-cckit/router"
	"net/http"
)

type RequestDocumentInterface struct {
	Collection string `json:"collection"`
}

func Document(c router.Context) (interface{}, error) {
	isInCollection := false
	var collectionState CollectionInterface
	for _, collection := range consoleCollection.Collections {
		if collection.Name == string(c.GetArgs()[2]) {
			collectionState = collection
			isInCollection = true
			break
		}
	}
	if !isInCollection {
		return c.JSON(constant.ApiResult{
			Result:      nil,
			Message:     `collection name is not exist`,
			MessageCode: http.StatusBadRequest,
		})
	}

	index := fmt.Sprintf(`%s~%s~%s~%s`, collectionState.ObjectType, `primary_key`, `transaction_id`, `transaction_timestamp`)
	switch string(c.GetArgs()[1]) {
	case `find`:
		return documentFind(c, index, collectionState)
	case `count`:
		return documentCount(c, index)
	}
	return nil, nil
}

func documentFind(c router.Context, index string, collectionState CollectionInterface) (interface{}, error) {
	var documentArray []interface{}
	var document interface{}
	valBytes, err := c.Stub().GetStateByPartialCompositeKey(index, []string{})
	if err != nil {
		return nil, err
	}
	kPosition := 0
	for i := 0; valBytes.HasNext(); i++ {
		resRange, err := valBytes.Next()
		if err != nil {
			return nil, err
		}
		_, compositeKeyParts, err := c.Stub().SplitCompositeKey(resRange.Key)
		if err != nil {
			return nil, err
		}
		item, err := c.GetStateByKey(collectionState.ObjectType, compositeKeyParts[kPosition])
		if err != nil {
			return nil, fmt.Errorf("Failed to get state for " + compositeKeyParts[kPosition])
		} else if valBytes == nil {
			return nil, fmt.Errorf("Value does not exist: " + compositeKeyParts[kPosition])
		}
		if err = json.Unmarshal(item, &document); err != nil {
			return nil, err
		}
		documentArray = append(documentArray, document)
	}
	return c.JSON(constant.ApiResult{
		Result:      &documentArray,
		Message:     `OK`,
		MessageCode: http.StatusOK,
	})
}

func documentCount(c router.Context, index string) (interface{}, error) {
	valBytes, err := c.Stub().GetStateByPartialCompositeKey(index, []string{})
	if err != nil {
		return nil, err
	}
	counting := 0
	for i := 0; valBytes.HasNext(); i++ {
		counting++
		valBytes.Next()
	}
	return c.JSON(constant.ApiResult{
		Result:      counting,
		Message:     `OK`,
		MessageCode: http.StatusOK,
	})
}
