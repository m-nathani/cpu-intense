package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		res := fib(number-1) + fib(number-2)
		res = res + 1 - 1
		return res
	}
}

func fibloop(flength int) {
	for {
		fib(flength)
	}
}

func main() {

	args := os.Args
	flength := 100000
	ngoroutines := 10
	if len(args) > 1 {
		i, _ := strconv.Atoi(args[1])
		flength = i
		i, _ = strconv.Atoi(args[2])
		ngoroutines = i
	}

	fmt.Printf("Fibonaci series for length: %d\n", flength)
	fmt.Printf("Number of goroutines: %d\n", ngoroutines)

	for i := 0; i < ngoroutines; i++ {
		go fibloop(flength)
	}

	fibloop(flength)
	fmt.Println("done")
}
