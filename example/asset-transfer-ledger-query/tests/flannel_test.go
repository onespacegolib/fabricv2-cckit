package tests

import (
	"git.onespace.co.th/osgolib/fabricv2-cckit/helper"
	"git.onespace.co.th/osgolib/fabricv2-cckit/testcc"
	"testing"
)

const FLANNEL = `flannel`

func TestModels(t *testing.T) {
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset1", "blue", "5", "tom", "35")
	helper.PlayRes(res)
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset2", "blue", "5", "tom", "35")
	helper.PlayRes(res)
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset3", "blue", "5", "tom", "35")
	helper.PlayRes(res)
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset4", "blue", "5", "tom", "35")
	helper.PlayRes(res)
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset5", "blue", "5", "tom", "35")
	helper.PlayRes(res)

	res = cc.Query(testcc.MakeFunction(FLANNEL, `console`))
	helper.PlayRes(res)
	res = cc.Query(testcc.MakeFunction(FLANNEL, `document`), "find", "asset")
	//res = cc.Query(testcc.MakeFunction(FLANNEL, `document`), "count", "asset")
	helper.PlayRes(res)

	res = cc.Query(testcc.MakeFunction(FLANNEL, `document`), "count", "asset")
	//res = cc.Query(testcc.MakeFunction(FLANNEL, `document`), "count", "asset")
	helper.PlayRes(res)
}
