package main

import (
	"errors"
	"fmt"
	"time"
)

func init() {
	fmt.Println("Initializing package")
}

func main3() {
	res := add(5, 6)
	fmt.Println(res)
	natija, err := divideAB(12, 4)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(natija)
	// 3. Anonymous functions
	func() {
		fmt.Println("Hello, anonymous!")
	}()
	fmt.Println(sum(1, 2, 2))
	p := Person{Name: "Akramjon"}
	fmt.Println(p.Greet())

	fmt.Print(time.Now())
}

// 1. Regular function
func add(a, b int) int {
	return a + b
}

// 2. Functions with Multiple return values
func divideAB(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

// 4. Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// 5. Methods (Functions with Receivers)
type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello " + p.Name
}
