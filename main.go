package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked returns you dont have to return
func lengthAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// return 1
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// return 1

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if with a twist in go

// func canIDrink(age int) bool {
// 	if koreanAge := age + 1; koreanAge < 18 {
// 		// if age < 18 {
// 		return false
// 	}
// 	return true
// 	// } else {
// 	// 	return true
// }

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age > 18:
		return true
	case age == 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(12))
	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)
	defer fmt.Println("I'm done imagine")
	repeatMe("Avicenna", "Ibn Rush'd", "Hanbal", "Al-Ghazali")
	totalLength, upperName := lenAndUpper("Jafa")
	fmt.Println(totalLength, upperName)
	length, uppercase := lengthAndUpper("Jaafar")
	fmt.Println(length, uppercase)
	name := "Fifa"
	fmt.Println(name)
	fmt.Println(multiply(20, 1))
}
