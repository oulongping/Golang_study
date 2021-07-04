package main

import "fmt"
import "rsc.io/quote"

import "example.com/greetings"

func main() {

	fmt.Println("hello world!")
	fmt.Println(quote.Go())

	// Get a greeting message and print it.
	// 自定义模块的函数调用
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
