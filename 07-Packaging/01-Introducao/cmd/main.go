package main

import (
	"fmt"

	"github.com/batistondeoliveira/fullcycle_curso_golang/7-Pakaging/01/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	//fmt.Println("Hello, World!")
}
