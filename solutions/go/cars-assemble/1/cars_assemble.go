package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//panic("CalculateWorkingCarsPerHour not implemented")
   return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//panic("CalculateWorkingCarsPerMinute not implemented")
    perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
    cost := 0
	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10
	cost = (groupsOfTen * 95000) + (remainingCars * 10000)
	return uint(cost)
}
