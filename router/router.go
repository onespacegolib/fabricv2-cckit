package router

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/onespacegolib/fabricv2-cckit/response"
)

// v0.1.0 OSP
type (
	MethodType         string
	ContextHandlerFunc func(Context) peer.Response
	HandlerFunc        func(Context) (interface{}, error)
	HandlerMeta        struct {
		Hdl         HandlerFunc
		Type        MethodType
		ReqValidate interface{}
	}

	RequestMiddleware func(HandlerFunc, ...int) HandlerFunc

	Group struct {
		prefix        string
		handlers      map[string]*HandlerMeta
		chaincodeName string

		requestMiddleware []RequestMiddleware
	}

	Router interface {
		HandleInit(stubInterface shim.ChaincodeStubInterface)
		Handle(stubInterface shim.ChaincodeStubInterface)
		Query(path string, handle HandlerFunc, middleware ...RequestMiddleware) Router
		Invoke(path string, handle HandlerFunc, middleware ...RequestMiddleware) Router
	}
)

func (g *Group) Group(path string) *Group {
	return &Group{
		prefix:   g.prefix + path,
		handlers: g.handlers,
	}
}

func New(name string) *Group {
	g := new(Group)
	g.handlers = make(map[string]*HandlerMeta)
	return g
}

//func (g *Group) Init(handler HandlerFunc) *Group {
//	fmt.Println(1)
//	return g.Invoke(`init`, handler, nil)
//}
//

func (g *Group) Init(handler HandlerFunc) *Group {
	return g.addHandler(`invoke`, "init", handler, nil)
}

func (g *Group) Invoke(path string, handler HandlerFunc, requestValidate interface{}) *Group {
	return g.addHandler(`invoke`, `@`+path, handler, requestValidate)
}

func (g *Group) Query(path string, handler HandlerFunc, requestValidate interface{}) *Group {
	return g.addHandler(`query`, `@`+path, handler, requestValidate)
}

func (g *Group) addHandler(t MethodType, path string, handler HandlerFunc, requestValidate interface{}) *Group {
	g.handlers[g.prefix+path] = &HandlerMeta{
		Type: t,
		Hdl: func(context Context) (interface{}, error) {
			h := handler
			return h(context)
		},
		ReqValidate: requestValidate,
	}
	return g
}

func (g *Group) HandleInit(stub shim.ChaincodeStubInterface) peer.Response {
	h := g.buildHandler()
	return h(g.Context(stub).ReplaceArgs(append([][]byte{[]byte(`init`)}, stub.GetArgs()...)))
}

func (g *Group) Handle(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()
	if len(args) == 0 {
		return response.Error(errors.New(`empty args`))
	}
	h := g.buildHandler()
	return h(g.Context(stub))
}

func (g *Group) buildHandler() ContextHandlerFunc {
	return func(context Context) peer.Response {
		return g.handleContext(context)
	}
}

func (g *Group) handleContext(c Context) peer.Response {
	if handlerMeta, ok := g.handlers[c.Path()]; ok {
		if handlerMeta.ReqValidate != nil {
			err := c.RequestArgs(handlerMeta.ReqValidate)
			if err != nil {
				err := fmt.Errorf(`%s: %s`, err, c.Path())
				return shim.Error(err.Error())
			}
		}
		h := func(c Context) (interface{}, error) {
			c.SetHandler(handlerMeta)
			h := handlerMeta.Hdl
			return h(c)
		}
		resp := response.Create(h(c))
		return resp
	}
	err := fmt.Errorf(`%s: %s`, errors.New(`router handler error`), c.Path())
	return shim.Error(err.Error())
}

func (g *Group) Context(stub shim.ChaincodeStubInterface) Context {
	return NewContext(stub)
}

func NewContext(stub shim.ChaincodeStubInterface) *context {
	return &context{
		stub: stub,
	}
}
