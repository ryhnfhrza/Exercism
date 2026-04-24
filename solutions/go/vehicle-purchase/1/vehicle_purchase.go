package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true;
    }else{
        return false;
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var carOption string

    if option1 < option2 {
        carOption = option1;
    }else{
        carOption = option2;
    }
	return carOption + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var carPrice float64
	if age < 3 {
        carPrice = (originalPrice * 80) / 100;
    }else if age >= 3 && age<10{
        carPrice = (originalPrice * 70) / 100;
    }else{
        carPrice = (originalPrice * 50) / 100;
    }

    return carPrice
}
