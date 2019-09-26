package main

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}

	println(sum)

	//infinite loop

	// 	for {
	// 		println("Infinite loop")
	// 	}

	//for-each syntax
	names := []string{"young", "min", "choi"}

	for index, name := range names {
		println(index, name)
	}

	////break, continue

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue //밑에 코드를 뛰어넘고 다시 27행으로 간다.
		}
		a++
		if a > 10 {
			break
		}
	}
	print(a)

}
