package main

const a = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "Eliel"
	e float64 = 1.2
)

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	a := "X"
	println(a)
}
