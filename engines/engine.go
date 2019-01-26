package engines

type Engine interface {
	UserEngine
}

type genericEngine struct {
	UserEngine
}

func NewGenericEngine(userEngine UserEngine) Engine {
	return &genericEngine{userEngine}
}
