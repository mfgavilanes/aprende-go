package main

import (
	"encoding/json"
	"fmt"
)

// import "github.com/rs/zerolog/log"
type Producto struct {
	Nombre            string
	Precio, Descuento float64
	Stock             int
	Disponible        bool
}

type Producto2 struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"unidades_en_stock"`
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

	var p1 Producto

	fmt.Println("Producto 1:", p1)

	var p2 = Producto{
		Nombre:     "Laptop Dell",
		Precio:     1299.99,
		Stock:      15,
		Disponible: true,
	}

	fmt.Println("Producto 2:", p2)

	var p3 = Producto{
		Nombre: "Teclado Mecánico",
		Precio: 89.99,
	}

	fmt.Println("Producto 3:", p3)

	var p4 = Producto{"Ratón Óptico", 25.50, 50, 0, true}

	fmt.Println("Producto 4:", p4)

	var b = struct {
		Nombre string
	}{"Golang"}

	fmt.Println("Anónima:", b)

	prod := Producto2{Nombre: "Laptop", Precio: 1299.99, Stock: 15}

	jsonBytes, _ := json.Marshal(prod)
	fmt.Println(string(jsonBytes)) // ← Usa jsonBytes
}
