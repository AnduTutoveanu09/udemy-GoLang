package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	problem123()
	problem4()
	problem5()
	problem6()
	problem7()
	problem8()
}

func problem123(){
	message := "Hello World!"
	name := "Andu Tutoveanu"
	age := 23
	fmt.Println(message)
	fmt.Println(name)
	fmt.Println(age)
}

func problem4(){
	for i :=0; i <= 10; i++{
		fmt.Printf("%d ", i)
	}
}

func problem5(){
	var a float64 = 2
	var b float64 = 11
	res := math.Pow(a, b)
	fmt.Println("\n", res)

}

func problem6(){
	for j :=1; j <= 1000; j++{
		fmt.Println(j)
	}
}

func problem7(){
	randomNumber := rand.IntN(10)
	fmt.Println(randomNumber)
}

func problem8(){
	fmt.Println(time.Now().Format(time.RFC850))
}