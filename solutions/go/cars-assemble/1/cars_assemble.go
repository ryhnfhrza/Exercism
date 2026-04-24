package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    
	return float64(productionRate) * successRate /  100;
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionInHour := float64(productionRate) * successRate / 100
	productionInMinute := int(productionInHour / 60)

    return productionInMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupTen := int(carsCount / 10)
	priceGroupTen := groupTen * 95000
    
    anotherGroup := int (carsCount % 10)
    priceAnotherGroup := anotherGroup * 10000

    return uint(priceGroupTen) + uint(priceAnotherGroup)
}
