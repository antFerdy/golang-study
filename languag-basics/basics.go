package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init was called")
}

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

	// sliceDemo()

	// sliceAndArrayDemo()

	// sliceWithMake()

	// structDemo()

	// structWithNewDemo()

	// mapDemo()

	// initPackageDemo()

	// interfaceDemo()

	// embeddedInterface()

	// goroutineDemo()

	channelDemo()
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

func sliceAndArrayDemo() {
	coffeeShops := [5]string{"Starbucks", "Discovery coffee", "Espresso", "Costa coffee"}

	coffeeSlice := coffeeShops[0:4]
	fmt.Println(coffeeSlice)

	coffeeSlice[2] = "Espresso coffee shop"
	fmt.Println(coffeeShops)
}

func sliceWithMake() {
	sliceWithMke := make([]string, 5, 100)

	fmt.Println(sliceWithMke)
	fmt.Println(len(sliceWithMke))
}

type Person struct {
	name string
	age  int
}

func structDemo() {

	var p1 Person = Person{"Poli", 19}
	var p2 Person = Person{"Ember", 19}
	var p3 Person = Person{"Roy", 19}
	var p4 Person = Person{name: "Heli", age: 19}

	fmt.Println(p1, p2, p3, p4)
	fmt.Println("First person name", p1.name)
	modifyStruct(p1)
	fmt.Println("First person name after modify", p1.name)

	fmt.Println(&p1)
	autoDereference(&p1)
}

func modifyStruct(person Person) {
	person.name = "POLI"
}

func autoDereference(personPointer *Person) {
	fmt.Println(personPointer)
	fmt.Println(personPointer.name)
}

func structWithNewDemo() {
	type SyncedBuffer struct {
		lock   sync.Mutex
		buffer bytes.Buffer
	}

	p := new(SyncedBuffer)
	fmt.Println("Created with new", *p)

	var v SyncedBuffer
	fmt.Println("Just declared", v)
}

func mapDemo() {
	var timeZoneMap = map[string]int{
		"UTC": -6,
		"EST": -12,
		"PST": -8,
		"GMT": 0,
	}

	fmt.Println(timeZoneMap["UTC"])
	fmt.Println(timeZoneMap["INEXISTENT_KEY"]) //expect 0

	//check zero value present in map
	seconds, ok := fmt.Println(timeZoneMap["GMT"]) //ok = true expected

	fmt.Println(seconds, ok)

	offset("EEE", timeZoneMap)

	//check key exitst with absent value
	if _, present := timeZoneMap["GMT"]; present {
		fmt.Println("Value exists")
	}
}

func offset(tz string, timeZone map[string]int) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

func initPackageDemo() {
	//result := init.Hello()
	//fmt.Println("init result", result)
}

type Killer interface {
	kill(prayName string) bool
}

type Hitman struct {
	name   string
	weapon string
}

func (hitman Hitman) kill(prayName string) bool {
	fmt.Printf("Hitman %s killed pray %s by %s \n", hitman.name, prayName, hitman.weapon)
	return true
}

func interfaceDemo() {

	hitman1 := Hitman{"Hitman1", "Knife"}

	hitman1.kill("John Doe")

	killSomeone(hitman1)
}

func killSomeone(k Killer) {
	fmt.Println(k.kill("Someone"))
}

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func embeddedInterface() {
	c := Cube{3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
	fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
	fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}

func goroutineDemo() {
	go boring("go boring!")
	boring("boring!")
}

func boring(msg string) {

	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

	}

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channelDemo() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
