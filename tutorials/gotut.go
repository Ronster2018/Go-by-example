package main

import ( // All imports need to be in double quotes or in ()
	"fmt"
	"math"
	"math/rand" // unlike python, you use / instead of . for nested imports
	"time"
)

func test() {
	fmt.Println("Welcome to Go!")
}

func random() {
	rand.Seed(time.Now().UnixNano()) // Changing the random int  generators seed number. Without that, the random number will always be the same.
	fmt.Println("A random number from 1 to 100: ", rand.Intn(100))
}

func squareRoot() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

// You need to define the types of the paramaters and the data type of what's expected to reutrn if the functions is to return something.
// If both of the paramaters are the same type, you can get away with defining the type at the end
// Ex: func add(x, y float64)
func add(x float64, y float64) float64 {
	return x + y
}

// If you are reutrning several items, you need to specify how many you're reutnring and also the tupe
func greeting(a, b string) (string, string) {
	return a, b
}

// The main function is considered the main entrypoint that will always run.
func main() {
	test()
	random()
	squareRoot()

	var num1 float64 = 5.6
	var num2 float64 = 9.3

	//var num3, num4 float64 = 3.4, 2.3

	// num5, num6 := 3.6, 6.7 // We dont have to define the type. It can be assumed at runtime but it cannot change.

	const x int = 5 //  a variable that does not change

	w1, w2 := "Hey", "There"

	fmt.Println(greeting(w1, w2))

	fmt.Println(add(num1, num2))
}
