package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
    }
	car.battery = 100;
	car.speed = speed;
    car.batteryDrain = batteryDrain
    

    return car
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{    
    }
	track.distance = distance    

    return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery - car.batteryDrain >= 0 {
	car.battery = car.battery - car.batteryDrain;
    car.distance = car.distance + car.speed
} else if car.battery - car.batteryDrain < 0{
        car.distance = car.distance ;
}
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

    var decision bool
    
    result1 := track.distance / car.speed
    result2 := car.battery - (car.batteryDrain * result1)

    if result2 < 0 {
        decision = false
    }else{
        decision = true
    }
    return decision
}
