package jedlik

import "fmt"

// Drive drives the car once.
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}


// DisplayDistance returns the distance driven as a string.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the battery level as a string.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if the car can finish the given track distance.
func (c Car) CanFinish(trackDistance int) bool {
	maxDrives := c.battery / c.batteryDrain
	maxDistance := maxDrives * c.speed

	return maxDistance >= trackDistance
}

