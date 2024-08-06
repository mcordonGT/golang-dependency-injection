package engine

// ElectricEngine struct with batteryCapacity field
type ElectricEngine struct {
	batteryCapacity int
}

// Start method for ElectricEngine
func (ee *ElectricEngine) Start() string {
	return "Engine started"
}

// Stop method for ElectricEngine
func (ee *ElectricEngine) Stop() string {
	return "Engine stopped"
}

// NewElectricEngine function to create a new ElectricEngine
func NewElectricEngine(batteryCapacity int) *ElectricEngine {
	return &ElectricEngine{
		batteryCapacity: batteryCapacity,
	}
}
