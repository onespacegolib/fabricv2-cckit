package flannel

import (
	"encoding/json"
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/router"
	"github.com/fatih/structs"
	"strings"
)

const (
	TypeTransaction = `transaction`
	TypeAsset       = `asset`
	TypeParticipant = `participant`
)

var (
	ok  bool
	err error
)

type InitStubInterface struct {
	ModelOrg  string
	ModelType string
	ModelName string
	Debug     *bool
}

type Schema struct {
	ObjectType           string `json:"object_type"`
	TransactionTimestamp int64  `json:"transaction_timestamp"`
	TransactionID        string `json:"transaction_id"`
}

const FLANNEL_CONSOLE = `flannel_console`

type CollectionInterface struct {
	Name       string `json:"name"`
	ObjectType string `json:"object_type"`
}

type ConsoleInterface struct {
	Collections []CollectionInterface `json:"collections"`
}

type (
	Flannel interface {
		Error() error

		Chain(string, string) Flannel

		Model(interface{}) Flannel

		// InvokeLedger : force to create state in ledger
		InvokeLedger(key string, arg interface{}) Flannel
		// Create : new create a ledger transaction
		Create(interface{}) Flannel
		// Exist : checking a state is existing
		Exist(id string) (bool, Flannel)
		// FindById : query a ledger by using key state
		FindById(id string, arg interface{}) Flannel
		FindByIdAndUpdate(id string, modify interface{}) Flannel

		//FindByIdAndUpdate(id string, arg interface{}) Flannel

		// UpdateOne : update ledger state using struct in the same query
		UpdateOne(arg interface{}) Flannel

		// InitChainCompositeKey :
		InitChainCompositeKey(index ...string) ChainCompositeKey
	}
	flannel struct {
		context     router.Context
		init        *InitStubInterface
		schema      *Schema
		model       interface{}
		modelDefine map[string]interface{}
		modelStruct *structs.Struct
		payload     interface{}
		chain       map[string]string
		err         error
	}
)

func ObjectType(stub *InitStubInterface) string {
	return fmt.Sprintf(`%s.%s.%s`, stub.ModelOrg, stub.ModelType, stub.ModelName)
}

func InitFlannel(c router.Context, stub *InitStubInterface) Flannel {

	return &flannel{
		context:     c,
		init:        stub,
		err:         nil,
		modelDefine: map[string]interface{}{},
		payload:     nil,
		schema:      &Schema{},
	}
}

func (f flannel) Error() error {
	return f.err
}

var consoleCollection ConsoleInterface

func (f flannel) Model(arg interface{}) Flannel {
	if f.init.ModelName == "" {
		f.init.ModelName = strings.ToLower(structs.Name(arg))
	}
	f.model = arg
	f.modelDefine = structs.Map(arg)

	timeStamp, err := f.context.Stub().GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf(`error to fine transaction timestamp`)
	}
	f.schema = &Schema{
		ObjectType:           ObjectType(f.init),
		TransactionTimestamp: timeStamp.GetSeconds(),
		TransactionID:        f.context.Stub().GetTxID(),
	}

	isInCollection := false

	if len(consoleCollection.Collections) == 0 {
		key, _ := f.context.GetStateByKey(FLANNEL_CONSOLE, `collection`)
		_ = json.Unmarshal(key, &consoleCollection)
	}
	for _, collection := range consoleCollection.Collections {
		if collection.Name == f.init.ModelName {
			isInCollection = true
			break
		}
	}
	if !isInCollection {
		consoleCollection.Collections = append(consoleCollection.Collections, CollectionInterface{
			Name:       f.init.ModelName,
			ObjectType: f.schema.ObjectType,
		})

		marshal, _ := json.Marshal(&consoleCollection)
		_ = f.context.PutStateByKey(FLANNEL_CONSOLE, `collection`, marshal)
	}
	return f
}

func (f flannel) Chain(name string, index string) Flannel {
	if _, ok := f.chain[name]; !ok {
		f.chain[name] = index
	}
	return f
}
