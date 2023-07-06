package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.20, "João": 2000.3, "Maria": 300.0}

	println(Soma(m))
	println(Soma(m2))

	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}

	println(Soma(m3))
}
