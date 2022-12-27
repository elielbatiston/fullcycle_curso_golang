package main

import "github.com/batistondeoliveira/fullcycle_curso_golang/7-Pakaging/03/math"

// go work init ./math ./sistema
// go mod tidy -e => ignora pacotes nao existentes

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
}
