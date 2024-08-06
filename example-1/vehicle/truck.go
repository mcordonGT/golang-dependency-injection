package vehicle

import "github.com/mcordonGT/golang-dependency-injection/example-1/engine"

type Truck struct {
	engine *engine.DieselEngine
}

func NewTruck(engine *engine.DieselEngine) *Truck {
	return &Truck{
		engine: engine,
	}
}

func (t *Truck) Start() string {
	t.engine.Start()
	t.engine.WarmUp()
	return "Truck is started"
}

func (t *Truck) Drive() string {
	return "Truck is driving"
}

func (t *Truck) Stop() string {
	t.engine.CoolDown()
	t.engine.Stop()
	return "Truck is stopped"
}
