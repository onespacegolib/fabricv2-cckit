package request

import (
	"encoding/json"
	"github.com/go-playground/validator"
)

var err	 error

func Validate(args interface{}, valiArgs []byte) (interface{}, error) {
	if err = json.Unmarshal(valiArgs, args); err != nil { return nil, err }
	validate := validator.New()
	if err := validate.Struct(args); err != nil { return nil, err }
	return args, err
}
