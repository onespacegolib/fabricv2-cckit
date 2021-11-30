package response

import (
	"errors"
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/convert"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func Error(err interface{}) peer.Response {
	return shim.Error(fmt.Sprintf("%s", err))
}

func Success(data interface{}) peer.Response {
	bb, err := convert.ToBytes(data)
	if err != nil {
		return shim.Success(nil)
	}
	return shim.Success(bb)
}

func Create(data interface{}, err interface{}) peer.Response {
	var errObj error

	switch e := err.(type) {
	case nil:
		errObj = nil
	case bool:
		if !e {
			errObj = errors.New(`boolean error: false`)
		}
	case string:
		if e != `` {
			errObj = errors.New(e)
		}
	case error:
		errObj = e
	default:
		panic(fmt.Sprintf(`unknowm error type %s`, err))
	}

	if errObj != nil {
		errRes := peer.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("%s", errObj),
			Payload: nil,
		}
		bb, err := convert.ToBytes(data)
		if err != nil {
			return errRes
		}
		errRes.Payload = bb
		return errRes
	}
	return Success(data)
}
