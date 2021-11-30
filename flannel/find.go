package flannel

import (
	"encoding/json"
	"github.com/fatih/structs"
)

func (f flannel) findByIdByte(id string) ([]byte, error) {
	state, err := f.context.GetStateByKey(f.schema.ObjectType, id)
	if err != nil {
		return nil, err
	}
	if state == nil {
		return nil, ErrRecordNotFound
	}
	return state, nil
}

func (f flannel) FindById(id string, arg interface{}) Flannel {
	argByte, err := f.findByIdByte(id)
	if err != nil {
		f.err = err
		return f
	}

	if err := json.Unmarshal(argByte, &arg); err != nil {
		f.err = err
	}
	return f
}

func (f flannel) FindByIdAndUpdate(id string, modify interface{}) Flannel {
	// check modify is a struct
	if ok := structs.IsStruct(modify); !ok {
		panic(ErrCompileStruct)
	}

	// mapping struct in to model fuction
	f.Model(modify)
	f.payload = modify
	f.modelStruct = structs.New(f.payload)

	// find a primary key of this model
	var name, _ string
	if name, _, ok = f.getValueOfTag(`primary_key`); !ok {
		panic(ErrTagPrimaryKey)
	}

	// find replacement value interface from modify interface
	modifyMap := map[string]interface{}{}
	for _, field := range f.modelStruct.Fields() {
		fieldName := field.Name()
		if fieldName != name && !field.IsZero() {
			if _, ok := modifyMap[fieldName]; !ok {
				modifyMap[fieldName] = field.Value()
			}
		}
	}

	// create copy modify in to new argument
	arg := modify

	// find a argument in a ledger
	if err := f.FindById(id, &arg).Error(); err != nil {
		return f
	}

	// bind modify map to change value interface of argument
	structArg := structs.New(arg)
	for s, v := range modifyMap {
		if err := structArg.Field(s).Set(v); err != nil {
			f.err = err
			return f
		}
	}

	// store and return
	return f.InvokeLedger(id, arg)
}
