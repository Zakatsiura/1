package main

import (
	"errors"
	"fmt"
	// "log"
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

	// message := sayHello("John", 45)
	// printMessage(message)

	// mess, _, err := enterTheClub(15)
	// if err != nil {
	// 	// log.Fatal(err)
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(mess)

	// fmt.Println(enterTheClub(85))

	fmt.Println(prediction("mon"))
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You are %d", name, age)
}

func enterTheClub(age int) (string, bool, error) {
	if age >= 18 && age < 45 {
		response := "Yes"
		return response, true, nil
	}
	if age >= 45 && age < 65 {
		response := "Yes but be carefull"
		return response, true, nil
	}
	if age > 65 {
		response := "Not for you dude"
		return response, false, errors.New("You are too old")
	}
	return "No", false, errors.New("You are too young")
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "mon":
		return "Good M", nil
	case "tue":
		return "Good T", nil
	case "wen":
		return "Good W", nil
	case "th":
		return "Good Th", nil
	case "fr":
		return "Good F", nil
	case "st":
		return "Good St", nil
	case "sd":
		return "Good Sd", nil
	default:
		return "Hello", errors.New("wrong day")
	}
}
  