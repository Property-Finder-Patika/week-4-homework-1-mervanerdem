/*
Create an abstraction to make conversions among different temperature systems such
as Celsius, Fahrenheit and Kelvin. Create a calculator that is loosely coupled to
that abstraction i.e. program to an interface, not an implementation principles is
observed. Make sure that the program can interact with the user so that user can
achieve some conversions using the program.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Using is 1f,1c,1k... \nExp: 1c Output will be 0C = 32F = 273K")
	fmt.Println("For exit press enter...")
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}

}
