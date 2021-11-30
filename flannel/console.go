package flannel

import (
	"encoding/json"
	"git.onespace.co.th/osgolib/fabricv2-cckit/router"
)

func Console(c router.Context) (interface{}, error) {
	var collection ConsoleInterface
	key, _ := c.GetStateByKey(FLANNEL_CONSOLE, `collection`)
	_ = json.Unmarshal(key, &collection)
	return key, nil
}
