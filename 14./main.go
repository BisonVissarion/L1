package main

import (
	"fmt"
)

func getType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	var x int = 42
	var y string = "Hello"
	var z bool = true
	var ch chan interface{} = make(chan interface{})

	fmt.Printf("Тип переменной x: %s\n", getType(x))
	fmt.Printf("Тип переменной y: %s\n", getType(y))
	fmt.Printf("Тип переменной z: %s\n", getType(z))
	fmt.Printf("Тип переменной ch: %s\n", getType(ch))
}
