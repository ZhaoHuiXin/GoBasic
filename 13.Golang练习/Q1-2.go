package main

import "fmt"

func main() {
	for i:=1;i<=100;i++{
		a := i%3
		b := i%5
		switch{
		case a==0 && b==0:
			fmt.Println("FizzBuzz")
		case a==0:
			fmt.Println("Fizz")
		case b==0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}