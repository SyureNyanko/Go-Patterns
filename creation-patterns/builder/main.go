package main

/* package car */
import (
	"fmt"
)

/******* builder pattern ********/
type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

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

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

/******* builder pattern ********/

/* implement */
type car struct {
	color  Color
	wheels Wheels
	speed  float64
}

func (c *car) Drive() {
	fmt.Println("Driving...")
}

func (c *car) Stop() {
	fmt.Println("Stopping...")
}

func NewBuilder() *builder {
	return &builder{}
}

type builder struct {
	color  Color
	wheels Wheels
	speed  float64
}

func (b *builder) Wheels(w Wheels) *builder {
	fmt.Printf("Wheels %s is build \n", w)
	return b
}

func (b *builder) TopSpeed(s Speed) *builder {
	fmt.Printf("Speed %f is build \n", s)
	return b
}

func (b *builder) Paint(c Color) *builder {
	fmt.Printf("Paint %s is build \n", c)
	return b
}

func (b *builder) Build() *car {
	return &car{
		color:  b.color,
		wheels: b.wheels,
		speed:  b.speed,
	}
}

/* use case */
func main() {
	assembly := NewBuilder().Paint(RedColor)
	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	/* red sports car */
	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()

}
