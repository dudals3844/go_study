package main

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20) //f calculator = add에 있는 익명함수러 정의된다.
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}
