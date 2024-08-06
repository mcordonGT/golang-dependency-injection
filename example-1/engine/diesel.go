package engine

// DieselEngine struct with fuelCapacity field
type DieselEngine struct {
	fuelCapacity int
}

// Start method for DieselEngine
func (de *DieselEngine) Start() string {
	return "Engine started"
}

// Stop method for DieselEngine
func (de *DieselEngine) Stop() string {
	return "Engine stopped"
}

// WarmUp method for DieselEngine
func (de *DieselEngine) WarmUp() string {
	return "Engine is warming up"
}

// CoolDown method for DieselEngine
func (de *DieselEngine) CoolDown() string {
	return "Engine is cooling down"
}

// NewDieselEngine function to create a new DieselEngine
func NewDieselEngine(fuelCapacity int) *DieselEngine {
	return &DieselEngine{
		fuelCapacity: fuelCapacity,
	}
}
