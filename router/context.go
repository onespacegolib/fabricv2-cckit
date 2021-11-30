package router

import (
	"encoding/json"
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/constant"
	"git.onespace.co.th/osgolib/fabricv2-cckit/request"
	"github.com/fatih/structs"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type (
	Context interface {
		Stub() shim.ChaincodeStubInterface
		Request() interface{}
		Path() string
		GetArgs() [][]byte
		SetHandler(*HandlerMeta)
		ReplaceArgs(args [][]byte) Context
		RequestArgs(interface{}) error
		FirstArgs() []byte
		GetStateByKey(string, string) ([]byte, error)
		PutStateByKey(string, string, []byte) error
		DelStateByKey(string, string) error
		PutStateByIndex(string, string, []byte, string, []string) error
		PutStateByCompositeKey(string, []string) error
		DelStateByCompositeKey(string, []string) error
		TimeUnixAndTx() (int64, string, error)

		// Json Response
		JSON(constant.ApiResult) (interface{}, error)
	}

	context struct {
		stub    shim.ChaincodeStubInterface
		args    [][]byte
		handler *HandlerMeta
		request interface{}
	}
)

func (c *context) JSON(r constant.ApiResult) (interface{}, error) {
	if r.MessageCode >= 300 {
		marshal, _ := json.Marshal(r)
		return nil, fmt.Errorf(string(marshal))
	}
	rMap := structs.Map(r)
	if _, ok := rMap[`path`]; !ok {
		rMap[`path`] = c.Path()
	}
	return rMap, nil
}

func (c *context) Stub() shim.ChaincodeStubInterface {
	return c.stub
}

func (c *context) Request() interface{} {
	return c.request
}

func (c *context) Path() string {
	if len(c.GetArgs()) == 0 {
		return ``
	}
	return string(c.GetArgs()[0])
}

func (c *context) GetArgs() [][]byte {
	if c.args != nil {
		return c.args
	}
	return c.stub.GetArgs()
}

func (c *context) SetHandler(h *HandlerMeta) {
	c.handler = h
}

func (c *context) ReplaceArgs(args [][]byte) Context {
	c.args = args
	return c
}

func (c *context) RequestArgs(args interface{}) error {
	requestValidated, err := request.Validate(args, c.FirstArgs())
	c.request = requestValidated
	return err
}

func (c *context) FirstArgs() []byte {
	return c.stub.GetArgs()[1]
}

func (c *context) GetStateByKey(lead string, key string) ([]byte, error) {
	return c.stub.GetState(lead + "#" + key)
}

func (c *context) PutStateByKey(lead string, key string, b []byte) error {
	var err error
	err = c.stub.PutState(lead+"#"+key, b)
	if err != nil {
		return err
	}
	return nil
}

func (c *context) DelStateByKey(lead string, key string) error {
	var err error
	err = c.stub.DelState(lead + "#" + key)
	if err != nil {
		return err
	}
	return nil
}

func (c *context) PutStateByIndex(lead string, key string, b []byte, index string, m []string) error {
	var err error

	err = c.PutStateByKey(lead, key, b)
	if err != nil {
		return err
	}

	indexKey, err := c.stub.CreateCompositeKey(index, m)
	if err != nil {
		return err
	}

	value := []byte{0x00}
	err = c.stub.PutState(indexKey, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *context) PutStateByCompositeKey(index string, m []string) error {
	indexKey, err := c.stub.CreateCompositeKey(index, m)
	if err != nil {
		return err
	}

	value := []byte{0x00}
	err = c.stub.PutState(indexKey, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *context) DelStateByCompositeKey(index string, m []string) error {
	indexKey, err := c.stub.CreateCompositeKey(index, m)
	if err != nil {
		return err
	}
	err = c.stub.DelState(indexKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *context) TimeUnixAndTx() (int64, string, error) {
	txTimestamp, err := c.stub.GetTxTimestamp()
	if err != nil {
		return 0, "", err
	}
	return txTimestamp.GetSeconds(), c.stub.GetTxID(), nil
}
