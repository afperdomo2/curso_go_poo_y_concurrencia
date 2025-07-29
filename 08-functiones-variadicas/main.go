package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (int, int, int) {
	return x * 2, x * 3, x * 4
}

func getValuesWithReturnNames(x int) (double int, triple int, quadruple int) {
	double = x * 2
	triple = x * 3
	quadruple = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(4, 5, 6, 7, 8))
	fmt.Println(sum(9, 10, 11, 12, 13, 14, 15))

	printNames("Alice", "Bob", "Charlie")

	val1, val2, val3 := getValues(5)
	fmt.Println("Los valores son:", val1, val2, val3)

	val1, val2, val3 = getValuesWithReturnNames(5)
	fmt.Println("Los valores con nombres son:", val1, val2, val3)
}
