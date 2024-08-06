package engine

// GasolineEngine struct with fuelCapacity field
type GasolineEngine struct {
	fuelCapacity int
}

// Start method for GasolineEngine
func (ge *GasolineEngine) Start() string {
	return "Engine started"
}

// Stop method for GasolineEngine
func (ge *GasolineEngine) Stop() string {
	return "Engine stopped"
}

// NewGasolineEngine function to create a new GasolineEngine
func NewGasolineEngine(fuelCapacity int) *GasolineEngine {
	return &GasolineEngine{
		fuelCapacity: fuelCapacity,
	}
}
