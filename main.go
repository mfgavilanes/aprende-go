package main

import "fmt"

//import "github.com/rs/zerolog/log"

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
}
