package flannel

import (
	"encoding/json"
	"strings"
)

type (
	ChainList         []string
	ChainCompositeKey interface {
		Create(args ...string) ChainCompositeKey
		Error() error
		Find(filter []string, val *ChainList) ChainCompositeKey
	}
	chainCompositeKey struct {
		f        *flannel
		indexKey string
		indexMap map[string]int
		indexLen int
		err      error
	}
)

func (cck chainCompositeKey) Error() error {
	return cck.err
}

func (f flannel) InitChainCompositeKey(index ...string) ChainCompositeKey {
	indexMap := map[string]int{}
	for i, s := range index {
		if _, ok := indexMap[s]; !ok {
			indexMap[s] = i
		}
	}

	return &chainCompositeKey{
		f:        &f,
		indexKey: strings.Join(index[:], "~"),
		indexMap: indexMap,
		indexLen: len(index),
		err:      nil,
	}
}

func (cck chainCompositeKey) Create(args ...string) ChainCompositeKey {
	if ok := cck.Exist(args); ok {
		cck.err = ErrExistKeyInLedger
		return cck
	}
	if err = cck.f.context.PutStateByCompositeKey(cck.indexKey, args); err != nil {
		cck.err = err
		return cck
	}
	return cck
}

func (cck chainCompositeKey) Exist(filter []string) bool {
	var val ChainList
	if err := cck.Find(filter, &val).Error(); err != nil {
		return false
	}
	return len(val) >= 1
}

func (cck chainCompositeKey) Find(filter []string, val *ChainList) ChainCompositeKey {
	var valItem []interface{}

	valBytes, err := cck.f.context.Stub().GetStateByPartialCompositeKey(cck.indexKey, filter)
	if err != nil {
		cck.err = err
		return cck
	}
	for i := 0; valBytes.HasNext(); i++ {
		resRange, err := valBytes.Next()
		if err != nil {
			cck.err = err
			return cck
		}
		_, compositeKeyParts, err := cck.f.context.Stub().SplitCompositeKey(resRange.Key)
		if err != nil {
			cck.err = err
			return cck
		}
		valItem = append(valItem, compositeKeyParts)
	}
	bytes, err := json.Marshal(valItem)
	if err != nil {
		cck.err = err
		return cck
	}
	if err := json.Unmarshal(bytes, &val); err != nil {
		cck.err = err
		return cck
	}
	return cck
}

func (cck chainCompositeKey) UpdateOne(filter []string, new []string) ChainCompositeKey {
	var val ChainList
	if err := cck.Find(filter, &val).Error(); err != nil {
		cck.err = err
		return cck
	}
	if len(val) < 1 {
		cck.err = ErrRecordNotFound
		return cck
	}

	if err := cck.DeleteOne(filter).Error(); err != nil {
		cck.err = err
		return cck
	}

	if err := cck.Create(new...).Error(); err != nil {
		cck.err = err
		return cck
	}
	return cck
}

func (cck chainCompositeKey) DeleteOne(filter []string) ChainCompositeKey {
	if err := cck.f.context.DelStateByCompositeKey(cck.indexKey, filter); err != nil {
		cck.err = err
		return cck
	}
	return cck
}
