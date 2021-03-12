package main

import (
	"fmt"
	"math"
	"github.com/golang-learning/go-crash-course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7)) // 2
	fmt.Println(math.Ceil(2.7)) // 3
	fmt.Println(math.Sqrt(16)) // 4

	fmt.Println(strutil.Reverse("Hello"))
}
