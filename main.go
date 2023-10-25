package main

import (
	"fmt"
	"math"
)

func main() {
	// ax*x+bx+c=0

	var a, b, c float64
	fmt.Println("Son kiriting")
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		fmt.Println("tenglamani ildizi yoq")
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("Tenglama 1 taildiz ega: x = %.2f\n", x)
	} else {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Tenglama 2 ta ildizga ega:\n")
		fmt.Printf("x1 = %.2f\n", x1)
		fmt.Printf("x2 = %.2f\n", x2)
	}

}
