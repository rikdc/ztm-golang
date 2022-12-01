//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Car string
type Motorcycle string
type Truck string

func (v Car) String() string {
	return fmt.Sprintf("Car: %v", string(v))
}

func (d Car) PickLift() Lift {
	return StandardLift
}

func (v Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(v))
}

func (d Motorcycle) PickLift() Lift {
	return SmallLift
}

func (v Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(v))
}

func (d Truck) PickLift() Lift {
	return LargeLift
}

func directVehicle(p LiftPicker) {
	captions := make(map[Lift]string, 2)
	captions[SmallLift] = "Small"
	captions[StandardLift] = "Standard"
	captions[LargeLift] = "Large"

	fmt.Printf("send %v to %s lift\n", p, captions[p.PickLift()])
}

func main() {
	c := Car("Honda Civic")
	t := Truck("Ford F150")
	m := Motorcycle("A Motorcycle")

	directVehicle(c)
	directVehicle(t)
	directVehicle(m)
}
