package flannel

import (
	"errors"
	"fmt"
	"git.onespace.co.th/osgolib/fabricv2-cckit/convert"
	"github.com/fatih/structs"
	"strconv"
)

var (
	ErrCompileStruct     = errors.New(`error to compile interface is not a struct type In Create Function`)
	ErrTagPrimaryKey     = errors.New(`error to get a primary key in the struct. primary key is important to make a transaction`)
	ErrModelDefineSchema = errors.New(`error to mapping flannel 'Schema' to Model structure please check your model type`)
)

func (f flannel) mapFlannelScheme() Flannel {
	schemaName := structs.Name(&Schema{})
	if _, ok := f.modelDefine[schemaName]; !ok {
		f.err = ErrModelDefineSchema
		return f
	}
	if ok := structs.IsStruct(f.payload); !ok {
		f.err = ErrCompileStruct
		return f
	}
	s := f.modelStruct
	fields := s.Field(schemaName).Fields()

	schema := structs.Map(&f.schema)

	for _, field := range fields {
		if err := field.Set(schema[field.Name()]); err != nil {
			f.err = fmt.Errorf(`error to set '%s' in Flannel Schema : %v`, field.Name(), err)
		}
	}
	return f
}

func (f flannel) InvokeLedger(key string, arg interface{}) Flannel {
	var j []byte
	if j, err = convert.ToBytes(arg); err != nil {
		f.err = err
	}
	if f.err != nil {
		return f
	}
	if err := f.context.PutStateByKey(
		f.schema.ObjectType,
		key,
		j,
	); err != nil {
		f.err = err
	}
	return f
}

func (f flannel) createMainChainComposite(key string) error {
	// args
	index := fmt.Sprintf(`%s~%s~%s~%s`, f.schema.ObjectType, `primary_key`, `transaction_id`, `transaction_timestamp`)
	chain := make([]string, 0, 2)
	chain = []string{key, f.schema.TransactionID, strconv.FormatInt(f.schema.TransactionTimestamp, 10)}
	if err := f.context.PutStateByCompositeKey(index, chain); err != nil {
		return err
	}
	return nil
}

func (f flannel) Create(arg interface{}) Flannel {
	if ok := structs.IsStruct(arg); !ok {
		f.err = ErrCompileStruct
		return f
	}
	f.Model(arg)
	f.payload = arg
	f.modelStruct = structs.New(f.payload)

	var key string
	if _, key, ok = f.getValueOfTag(`primary_key`); !ok || key == "" {
		f.err = ErrTagPrimaryKey
		return f
	}
	f.mapFlannelScheme()

	if exist, _ := f.Exist(key); exist {
		f.err = ErrExistKeyInLedger
		return f
	}

	f.InvokeLedger(key, f.payload)
	if err := f.createMainChainComposite(key); err != nil {
		f.err = err
		return f
	}

	return f
}
