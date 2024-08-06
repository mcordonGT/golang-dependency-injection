package vehicle

import "github.com/mcordonGT/golang-dependency-injection/example-1/engine"

type EngineType int

const (
	Gasoline EngineType = iota
	Diesel   EngineType = iota
	Electric EngineType = iota
)

type MotorCycle struct {
	engine vehicleEngine
}

// NewMotorCycle function to create a new MotorCycle
// This function uses the engineType and capacity to create a new engine
// for the MotorCycle.
// The engineType is used to determine the type of engine to create, every
// time we add a new engine type, we need to update this function.
func NewMotorCycle(engineType EngineType, capacity int) *MotorCycle {
	var motorcycleEngine vehicleEngine
	switch engineType {
	case Gasoline:
		motorcycleEngine = engine.NewGasolineEngine(capacity)
	case Diesel:
		motorcycleEngine = engine.NewDieselEngine(capacity)
	case Electric:
		motorcycleEngine = engine.NewElectricEngine(capacity)
	}
	return &MotorCycle{
		engine: motorcycleEngine,
	}
}

func (m *MotorCycle) Start() string {
	return m.engine.Start()
}

func (m *MotorCycle) Drive() string {
	return "MotorCycle is driving"
}

func (m *MotorCycle) Stop() string {
	return m.engine.Stop()
}
