package main

import (
	"fmt"
)

func main() {

	// const message string = "Go go go"
	// var message string = "Go go go"
	// message  := "Go go go"

	// var message string
	// message = "Go go go !!!!"

	// =========== нулевые значения ========

	// даст пустую строку
	// var message string

	// даст 0
	// var message int or float32/64

	// ============== типы данных ===========

	// var message float32
	// var message = []byte("asd")
	// var c rune = 'f'

	// var b bool

	// множественное присвоение:
	// d, e, f := 1, 2, 3
	// d, _, f = 3, 5, 7
	// e, f = f, e
	// d, e = 3, 6

	// fmt.Println(reflect.TypeOf(message))
	// fmt.Println(message, b, c, d, e, f)

	message := sayHello("John", 45)
	printMessage(message)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You are %d", name, age)
}
