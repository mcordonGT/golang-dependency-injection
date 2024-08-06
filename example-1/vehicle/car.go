package vehicle

type vehicleEngine interface {
	Start() string
	Stop() string
}

// Car struct with engine field
type Car struct {
	// Car uses vehicleEngine interface to allow any type of engine
	// to be used with Car. This is an example of dependency injection.
	engine vehicleEngine
}

// NewCar function to create a new Car
func NewCar(engine vehicleEngine) *Car {
	// engine parameter is of type vehicleEngine interface, so any type
	// that implements vehicleEngine interface can be passed to NewCar.
	return &Car{
		engine: engine,
	}
}

// Start method for Car
func (c *Car) Start() string {
	// Car can use any method availabe in vehicleEngine interface.
	return c.engine.Start()
}

// Stop method for Car
func (c *Car) Stop() string {
	return c.engine.Stop()
}

func (c *Car) Drive() string {
	return "Car is driving"
}
