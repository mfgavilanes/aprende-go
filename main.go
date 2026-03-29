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

type Usuario struct {
	Name string
}

type mobile struct {
	brand string
}

type laptop struct {
	cpu string
}

type toaster struct {
	amount int
}

type kettle struct {
	quantity string
}

type socket struct{}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

func (socket) Plug(device mobile, power int) {
	device.Draw(power)
}

type television struct {
	brand string
}

type aireAcondicionado struct {
	mode string
}

type altavoz struct {
	power int
}

type controlRemoto struct{}

func (t television) Encender(power int) {
	fmt.Printf("%T -> marca: %s, potencia: %d\n", t, t.brand, power)
}

func (controlRemoto) Usar(device television, power int) {
	device.Encender(power)
}

func main() {

	s := "ñ"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))

	//log.Info().Msg("Hola")
	var z string = "Hola"
	fmt.Println(z, z[0])

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

	var a2 = [4]int{20, 15, 5, 25}
	var s2 = a2[1:3]
	// Salida: Array: [20 15 5 25], Longitud: 4, Capacidad: 4
	fmt.Printf("Array: %v, Longitud: %d, Capacidad: %d\n", a2, len(a2), cap(a2))

	// Output: Slice [15 5], Length: 2, Capacity: 3
	fmt.Printf("Slice: %v, Longitud: %d, Capacidad: %d", s2, len(s2), cap(s2))

	var m = map[string]Usuario{
		"a": {"Pedro"},
		"b": {"Sebastián"},
	}

	m["c"] = Usuario{"Manuel"}

	fmt.Println(m)
	c, ok := m["c"]
	fmt.Println("Key c:", c, ok)

	d, ok := m["d"]
	fmt.Println("Key d:", d, ok)

	m2 := mobile{"Apple"}
	//l2 := laptop{"Intel i9"}

	s3 := socket{}

	s3.Plug(m2, 10)
	//s3.Plug(l2, 50) // Error: cannot use l as mobile value in argument
	tv := television{"Samsung"}

	rc := controlRemoto{}
	rc.Usar(tv, 10)
}
