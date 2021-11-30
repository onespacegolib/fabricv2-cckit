package testcc

type (
	Factory interface {
	}
	factory struct {
		stub       *MockStub
		objectType string
		payload    map[string]interface{}
		chain      map[string]string
	}
)

func InitFactory(mockStub *MockStub, objectType string, payload map[string]interface{}, chain map[string]string) Factory {
	return &factory{
		stub:       mockStub,
		objectType: objectType,
		payload:    payload,
		chain:      chain,
	}
}
