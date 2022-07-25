package main

import "fmt"

func main() {
	numberList := []int{}
	for i := 0; i <= 10; i++ {
		numberList = append(numberList, i)
	}

	for i := range numberList {
		determineAndPrint(i)
	}
}

func determineAndPrint(i int) {
	var result string = ""
	if i%2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}
	fmt.Printf("%v is %v.\n", i, result)
}
