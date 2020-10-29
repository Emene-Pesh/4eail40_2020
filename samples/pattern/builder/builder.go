package main

import "fmt"

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

// Car represents a car.
type Car interface {
	Drive()
	Stop()
}

type FamilyCar struct {
	paint    Color
	wheels   Wheels
	topSpeed float64
}

func (f FamilyCar) Drive() {
	fmt.Printf("Drive at %f km/h\n", f.topSpeed)
}

func (f FamilyCar) Stop() {
	fmt.Println("")
}

type SportCar struct {
	paint    Color
	wheels   Wheels
	topSpeed float64
}

func (s SportCar) Drive() {
	fmt.Printf("Drive at %f km/h\n", s.topSpeed)
}

func (s SportCar) Stop() {
	fmt.Println("")
}

// ---- Builder implementation

// FamilyCarBuilder is a concrete implementation for a CarBuilder.
type FamilyCarBuilder struct {
	paint    Color
	topSpeed float64
}

// NewFamilyCarBuilder returns a CarBuilder object.
func NewFamilyCarBuilder() *FamilyCarBuilder {
	return new(FamilyCarBuilder)
}

func (cb *FamilyCarBuilder) Color(c Color) *FamilyCarBuilder {
	cb.paint = c
	return cb
}
func (cb *FamilyCarBuilder) TopSpeed(s float64) *FamilyCarBuilder {
	cb.topSpeed = s
	return cb
}

func (cb *FamilyCarBuilder) Build() *FamilyCar {
	return &FamilyCar{cb.paint, SteelWheels, cb.topSpeed}
}

// SportCarBuilder is a concrete implementation for a CarBuilder.
type SportCarBuilder struct {
	paint    Color
	wheels   Wheels
	topSpeed float64
}

func NewSportCarBuilder() *SportCarBuilder {
	return new(SportCarBuilder)
}

func (cb *SportCarBuilder) Color(c Color) *SportCarBuilder {
	cb.paint = c
	return cb
}

func (cb *SportCarBuilder) Wheels(w Wheels) *SportCarBuilder {
	cb.wheels = w
	return cb
}

func (cb *SportCarBuilder) TopSpeed(s float64) *SportCarBuilder {
	cb.topSpeed = s
	return cb
}

func (cb *SportCarBuilder) Build() *SportCar {
	return &SportCar{cb.paint, cb.wheels, cb.topSpeed}
}

func main() {

	familyCar := NewFamilyCarBuilder().
		Color(BlueColor).
		TopSpeed(150).
		Build() // The car object is only constructed at this point.
	familyCar.Drive()

	sportsCar := NewSportCarBuilder().
		Color(RedColor).
		Wheels(SteelWheels).
		TopSpeed(300).
		Build()
	sportsCar.Drive()
}
