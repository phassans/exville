package engines

import "github.com/phassans/exville/engines/database"

type Engine interface {
	database.DatabaseEngine
}

type genericEngine struct {
	database.DatabaseEngine
}

func NewGenericEngine(dbEngine database.DatabaseEngine) Engine {
	return &genericEngine{dbEngine}
}
