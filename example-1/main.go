package main

import (
	"github.com/mcordonGT/golang-dependency-injection/example-1/engine"
	"github.com/mcordonGT/golang-dependency-injection/example-1/vehicle"
)

func main() {

	// Create a new DieselEngine with fuel capacity of 100
	dieselEngine := engine.NewDieselEngine(100)

	// Create a new Truc with the diesel engine
	dieselTruck := vehicle.NewTruck(dieselEngine)
	dieselTruck.Start()
	dieselTruck.Drive()
	dieselTruck.Stop()

	// Create a new ElectricEngine with battery capacity of 100
	electricEngine = engine.NewElectricEngine(100)

	// Create a new Truc with the electric engine
	electricTruck := vehicle.NewTruck(electricEngine) // Compilation error!!!
	electricTruck.Start()
	electricTruck.Drive()
	electricTruck.Stop()

	// Create a new ElectricEngine with battery capacity of 100
	electricEngine := engine.NewElectricEngine(100)

	// Create a new Car with the electric engine
	electricCar := vehicle.NewCar(electricEngine)
	electricCar.Start()
	electricCar.Drive()
	electricCar.Stop()

	// Create a new GasolineEngine with fuel capacity of 50
	gasolineEngine := engine.NewGasolineEngine(50)

	// Create a new Car with the gasoline engine
	gasolineCar := vehicle.NewCar(gasolineEngine)
	gasolineCar.Start()
	gasolineCar.Drive()
	gasolineCar.Stop()

	// Create a new MotorCycle with Gasoline engine and fuel capacity of 20
	gasolineMotorCycle := vehicle.NewMotorCycle(vehicle.Gasoline, 20)
	gasolineMotorCycle.Start()
	gasolineMotorCycle.Drive()
	gasolineMotorCycle.Stop()

	// Create a new MotorCycle with Electric engine and battery capacity of 50
	electricMotorCycle := vehicle.NewMotorCycle(vehicle.Electric, 50)
	electricMotorCycle.Start()
	electricMotorCycle.Drive()
	electricMotorCycle.Stop()
}
