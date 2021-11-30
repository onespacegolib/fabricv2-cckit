package flannel

import (
	"github.com/fatih/structs"
)

func (f flannel) UpdateOne(arg interface{}) Flannel {
	if ok := structs.IsStruct(arg); !ok {
		panic(ErrCompileStruct)
	}
	f.Model(arg)
	f.payload = arg
	f.modelStruct = structs.New(f.payload)

	var key string
	if _, key, ok = f.getValueOfTag(`primary_key`); !ok {
		panic(ErrTagPrimaryKey)
	}

	if exist, _ := f.Exist(key); !exist {
		f.err = ErrRecordNotFound
		return f
	}
	return f.InvokeLedger(key, f.payload)
}
