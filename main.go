/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-10
 * This file is to display and sum 3 constants
 */

package main

import "fmt"

func main() {
//initialize numbers
const NUMBER1 int = 10
const NUMBER2 int = -20
const NUMBER3 int = 85

// INPUT - none

// PROCESS
// calculates the sum
answer := NUMBER1 + NUMBER2 + NUMBER3

// OUTPUT
// display the result
fmt.Println("The sum of", NUMBER1, "&", NUMBER2, "&", NUMBER3, "is:", answer)

fmt.Println("\nDone.")
}
