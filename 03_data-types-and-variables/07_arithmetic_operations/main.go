package main

import (
	"fmt"
)

func main(){
    num1 := 14
    num2 := 23

    sum := num1 + num2
    subtracted := num1 - num2
    multiplied := num1 * num2
    divided := 56 / 8
    modulus := 15 % 2

    fmt.Println(sum, subtracted, multiplied, divided, modulus)

    //The value is not yet changed
    fmt.Println(num1, num2)

    //These operations change the values stored inside the variables
    num1++
    num2--
    fmt.Println(num1, num2)
}