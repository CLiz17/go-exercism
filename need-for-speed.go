package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	if (car.battery - car.batteryDrain) >= 0 {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.distance + car.speed,
		}
	} else {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery,
			distance:     car.distance,
		}
	}
}

func CanFinish(car Car, track Track) bool {
	if (car.battery/car.batteryDrain)*car.speed >= track.distance {
		return true
	} else {
		return false
	}
}
