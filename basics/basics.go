package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {

	// variablesDemo()

	// forLoopDemo()

	// ifStatementDemo()

	// switchDemo()

	// result := demoNamedFunc()
	// fmt.Println(result)

	// name, age := demoMultipleReturnValues()
	// fmt.Println(name, age)

	// pointerDemo()

	// errorDemo()

	// arrayDemo()

	sliceDemo()

}

func variablesDemo() {
	var c, python, java bool
	var x, y, z int = 0, 0, 0
	var i int
	j := 1.0
	fmt.Println(i, c, python, java)
	fmt.Println(x, y, z)
	fmt.Printf("Type of j: %T, value of j: %v \n", j, j)
	fmt.Println("Default val of i", i)
}

func forLoopDemo() {

	fmt.Println("Basic loop")
	for i := 100; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("Loop without pre, and post statements")
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("Square root calculation with for loop")
	fmt.Println(Sqrt(25))

}

func Sqrt(x float64) float64 {
	fmt.Println("Find square root for x: ", x)
	times := 0
	guess := 1.0
	for z := 1.0; times <= 10; times++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("Value of z: %g \n", z)
		guess = z
	}

	return guess
}

func ifStatementDemo() error {
	if err := os.Chmod("log.txt", 0777); err != nil {
		log.Fatal(err)
	}

	return nil
}

func switchDemo() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func demoNamedFunc() (s string) {
	s = "Test named funct"
	return
}

func demoMultipleReturnValues() (name string, age int) {
	name = "Harry"
	age = 11
	return
}

func pointerDemo() {
	name := "Someone"
	pName := &name
	fmt.Println(name, pName)

	//dereference
	fmt.Println(*pName)

	changeVarValueV1(name)
	fmt.Println(name, pName)

	changeVarValueV2(&name)
	fmt.Println(name, pName)

	//empty pointer returns nil
	var emptyPointer *string
	fmt.Println(emptyPointer)

}

func changeVarValueV1(variable string) {
	variable = "Changed" //just changes var in local scope
}

func changeVarValueV2(variable *string) {
	*variable = "Changed" //changes value by dereference
}

func errorDemo() {
	change, err := buyWhine(11, 100)
	if err != nil {
		fmt.Println("Error while buying whine, ", err.Error())
		return
	}

	fmt.Println("Whine has been bought, change left", change)
}

func buyWhine(age, amount int) (change int, err error) {

	const whinePrice = 100

	if age < 18 {
		return amount, errors.New("age is lower than 18")
	}

	if amount < whinePrice {
		return amount, errors.New("not enough money")
	}

	change = amount - whinePrice
	err = nil

	return
}

func arrayDemo() {

	coffeeShops := [3]string{"costa", "starbucks", "segafredo"}

	for index, value := range coffeeShops {
		fmt.Printf("Index: %d, Value: %s \n", index, value)
	}

	fmt.Println("c way traversing")
	for i := 0; i < len(coffeeShops); i++ {
		fmt.Println(coffeeShops[i])
	}
}

func sliceDemo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	slicedPrimes := []int{}

	for idx := range primes {
		slicedPrimes = append(slicedPrimes, primes[idx])
	}

	fmt.Println(slicedPrimes)
}
