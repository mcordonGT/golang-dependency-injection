package vehicle

type EngineType enum {
	Gasoline = iota
	Diesel
	Electric
}

type MotorCycle struct {
	engine vehicleEngine
}
