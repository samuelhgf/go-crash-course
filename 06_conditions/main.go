package main

import "fmt"

func main() {
	x := 15
	y := 15

	if x <= y {
		fmt.Printf("%d is less then or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d\n", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color id red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}
}
