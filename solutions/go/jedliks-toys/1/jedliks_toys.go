package jedlik

import "fmt"

// TODO: define the 'Drive()' method

// Drive method facilitates the driving of the car
// it reduces the battery percentage and increases the distance travelled.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method

// DisplayDistance provides the string for the Distance LED to display
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
// DisplayBattery provides the string for the Battery LED to display
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

func (car Car) CanFinish(trackDistance int) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
