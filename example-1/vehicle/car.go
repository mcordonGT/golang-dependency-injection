package vehicle

type vehicleEngine interface {
	Start() string
	Stop() string
}

// Car struct with engine field
type Car struct {
	engine vehicleEngine
}

// NewCar function to create a new Car
func NewCar(engine vehicleEngine) *Car {
	return &Car{
		engine: engine,
	}
}

// Start method for Car
func (c *Car) Start() string {
	return c.engine.Start()
}

// Stop method for Car
func (c *Car) Stop() string {
	return c.engine.Stop()
}

func (c *Car) Drive() string {
	return "Car is driving"
}
