package main

func main() {
	a := 10
	b := 5
	c := 30

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > c || a > b {
		println("a > C || a > b")
	}
}
