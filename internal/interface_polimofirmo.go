package internal

import (
	"fmt"
	"math"
)

type Figura interface {
	CalcularArea() float64
	CalcularPerimetro() float64
}

type Cuadrado struct {
	Nombre string
	Lado   float64
}

func (c Cuadrado) CalcularArea() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) CalcularPerimetro() float64 {
	return c.Lado * 4
}

type Rectangulo struct {
	Nombre    string
	LadoLargo float64
	LadoCorto float64
}

func (r Rectangulo) CalcularArea() float64 {
	return r.LadoLargo * r.LadoCorto
}

func (r Rectangulo) CalcularPerimetro() float64 {
	return (r.LadoLargo + r.LadoCorto) * 2
}

type Circulo struct {
	Nombre string
	Radio  float64
}

func (c Circulo) CalcularArea() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) CalcularPerimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func NombreFigura(i interface{}) {
	fmt.Printf("Figura: %+v\n", i)
}

func Draw(f Figura) {
	fmt.Printf("Area: %f, Perimetro: %f\n", f.CalcularArea(), f.CalcularPerimetro())
}
