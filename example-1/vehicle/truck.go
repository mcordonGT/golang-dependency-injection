package vehicle

import "github.com/mcordonGT/golang-dependency-injection/example-1/engine"

type Truck struct {
	engine *engine.DieselEngine
}

// NewTruck function to create a new Truck
// This function uses the engine parameter to create a new Truck.
// The engine parameter is of type *engine.DieselEngine, so only DieselEngine
// can be used with Truck.
// If we want to use a different type of engine with Truck, we need to create
// a new Truck function that accepts the new engine type.
func NewTruck(engine *engine.DieselEngine) *Truck {
	return &Truck{
		engine: engine,
	}
}

func (t *Truck) Start() string {
	// Truck can use any method availabe in engine.DieselEngine.
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
