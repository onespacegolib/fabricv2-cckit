package convert

import (
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

func ArgsToBytes(iargs ...interface{}) (aa [][]byte, err error) {
	args := make([][]byte, len(iargs))

	for i, arg := range iargs {
		val, err := ToBytes(arg)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(`unable to convert invoke arg[%d]`, i))
		}
		args[i] = val
	}

	return args, nil
}

func ToBytes(value interface{}) ([]byte, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case ToByter:
		return v.ToBytes()
	case proto.Message:
		return proto.Marshal(proto.Clone(v))
	case bool:
		return []byte(strconv.FormatBool(v)), nil
	case string:
		return []byte(v), nil
	case uint, int, int32:
		return []byte(fmt.Sprint(v)), nil
	case []byte:
		return v, nil

	default:
		valueType := reflect.TypeOf(value).Kind()

		switch valueType {
		case reflect.Ptr, reflect.Struct, reflect.Array, reflect.Map, reflect.Slice:
			return json.Marshal(value)
		case reflect.String:
			return []byte(reflect.ValueOf(value).String()), nil

		default:
			return nil, fmt.Errorf(
				`toBytes converting supports ToByter interface,struct,array,slice,bool and string, current type is %s`,
				valueType)
		}

	}
}