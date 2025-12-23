package jedlik

import "fmt"

func (car *Car) Drive() {
	if car.battery-car.batteryDrain < 0 {
		return
	}
	car.battery = car.battery - car.batteryDrain
	car.distance = car.distance + car.speed
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	maxDistance := car.battery * car.speed / car.batteryDrain
	return maxDistance >= trackDistance

}
