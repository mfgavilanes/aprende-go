package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red("¡Hola espacio de trabajo de color ROJO!"))
}
