// Golang program to illustrate the usage of
// fmt.Scanln() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
	"math"
)

// Calling main
func main() {

	// Declaring some variables
	var num_count int

	fmt.Println("How many number do you want to loop ?")
	fmt.Scanln(&num_count)

	for i := 0; i < num_count; i++ {
		num_check := i + 1
		//	fmt.Println(num_check)

		res_5 := math.Sqrt(float64(num_check))
		//	fmt.Println("square -", res_5)
		res_5Check := math.Floor(float64(res_5))
		//	fmt.Println("check -", res_5Check)

		//  fmt.Printf("\nResult 5: %.1f", res_5)
		res_6 := math.Cbrt(float64(num_check))
		//	fmt.Println("Cube -", res_6)
		res_6Check := math.Floor(float64(res_6))
		//	fmt.Println("check -", res_6Check)
		if num_check != 1 {
			switch {
			case (res_5 == res_5Check) && (res_6 == res_6Check):
				fmt.Println(num_check, "is Square Cube")
			//	checkPoint1 := 1
			case res_5 == res_5Check:
				fmt.Println(num_check, "is Square")
			//	checkPoint1 := 1
			case res_6 == res_6Check:
				fmt.Println(num_check, "is Cube")
			//	checkPoint2 := 1
			default:
				fmt.Println(num_check, "is Ordinary number")
			}
		} else {
			fmt.Println(num_check, "is Ordinary number")

		}

	}

}
