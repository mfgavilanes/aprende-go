package main

import "fmt"

// import "github.com/rs/zerolog/log"
type Producto struct {
	Nombre            string
	Precio, Descuento float64
	Stock             int
	Disponible        bool
}

func main() {
	//log.Info().Msg("Hola")
	fmt.Println("hola")

	a := 4

	var p *int = &a

	fmt.Println("dirección:", p)
	fmt.Println("valor:", *p)

	p = &a

	fmt.Println("antes", a)
	fmt.Println("dirección:", p)

	*p = 8
	fmt.Println("después:", a)

	var prod1 Producto

	fmt.Println("Producto 1:", prod1)

	var prod2 Producto

	fmt.Println("Producto 2:", prod2)

	var prod3 = Producto{
		Nombre:     "Laptop Dell",
		Precio:     1299.99,
		Stock:      15,
		Disponible: true,
	}

	fmt.Println("Producto 3:", prod3)
}
