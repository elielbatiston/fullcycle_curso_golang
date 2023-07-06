package main

import "github.com/batistondeoliveira/fullcycle_curso_golang/7-Pakaging/03/math"

// go mod edit -replace github.com/batistondeoliveira/fullcycle_curso_golang/7-Pakaging/03/math=../math

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
}
