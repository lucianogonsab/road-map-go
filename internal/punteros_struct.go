package internal

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func IntercambiarValores(a, b *int) {
	*a, *b = *b, *a
}

func IntercambiarNombre(p *Persona) {
	p1 := *p
	p.Nombre, p.Apellido = p1.Apellido, p1.Nombre
}
