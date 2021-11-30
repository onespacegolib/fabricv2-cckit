package router_test
//
//import (
//	"github.com/hyperledger/fabric/core/chaincode/shim"
//	"github.com/hyperledger/fabric/protos/peer"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//
//	"onespace/cckit/router"
//	testcc "onespace/cckit/testing"
//	"testing"
//)
//
//func TestRouter(t *testing.T)  {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Router suite")
//}
//
//func New() *router.Chaincode {
//	r := router.New(`router`)
//
//	r.Init(
//		func(ctx router.Context) (interface{}, error) {
//			return nil, nil
//		})
//	r.Group(`group1`).
//		Invoke(`empty`, func(ctx router.Context) (interface{}, error) {
//			return nil, nil
//		})
//
//	r.Group(`group2`).Invoke(`empty`, func(context router.Context) (i interface{}, e error) {
//		return nil, nil
//	})
//	return router.NewChaincode(r)
//}
//
//var cc *testcc.MockStub
//
//var	_ = Describe(`Router`, func() {
//	BeforeSuite(func() {
//		cc = testcc.NewMockStub(`Router`, New())
//	})
//
//	It(`Allow init response`, func() {
//		Expect(cc.Init()).To(
//			BeEquivalentTo(peer.Response{
//				Status:  shim.OK,
//				Payload: nil,
//				Message: ``,
//			}))
//	})
//	It(`Allow empty response`, func() {
//		Expect(cc.Invoke(`group1empty`)).To(
//			BeEquivalentTo(peer.Response{
//				Status:  shim.OK,
//				Payload: nil,
//				Message: ``,
//			}))
//	})
//
//	It(`Allow empty response`, func() {
//		Expect(cc.Invoke(`group2empty`)).To(
//			BeEquivalentTo(peer.Response{
//				Status:  shim.OK,
//				Payload: nil,
//				Message: ``,
//			}))
//	})
//})
//
