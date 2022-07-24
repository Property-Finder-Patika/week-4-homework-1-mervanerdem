/*Add error handling to the example calculator.go of ch07 we studied at the end of
  last class. And also make sure that the example works file even though math function
  name entered by the user is lower case of name registered by the calculator. i.e.
  make the calculator match math function names case insensitive.
*/
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []MathFunction //calling interface from other file which name is math.go
}

func (c *Calculator) addMathFunction(m MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	functions()
	calculator()
	startCalculator()
}

func functions() {
	sin := Sin{"Sinus"}
	fmt.Printf("%v\n", sin)
	sin30 := sin.Calculate(math.Pi / 3)
	fmt.Printf("Sinus of 30 degree is %f\n", sin30)

	cos := Cos{"Cosinus"}
	fmt.Printf("%v\n", sin)
	cos30 := cos.Calculate(math.Pi / 2)
	fmt.Printf("Cosinus of 30 degree is %f\n", cos30)

	log := Log{"Log of base e"}
	fmt.Printf("%v\n", log)
	logE := log.Calculate(2.71828)
	fmt.Printf("Log of Euler constant is %f\n", logE)

	var mf1 MathFunction = sin
	fmt.Printf("%v\n", mf1)

	mf1 = cos
	fmt.Printf("%v\n", mf1)

	mf1 = log
	fmt.Printf("%v\n", mf1)
}

func calculator() {
	myCalculator := Calculator{}

	myCalculator.addMathFunction(Sin{"Sinus"})
	myCalculator.addMathFunction(Cos{"Cosines"})
	myCalculator.addMathFunction(Log{"Log"})

	fmt.Println(myCalculator.doCalculation("Sinus", math.Pi/3))
	fmt.Println(myCalculator.doCalculation("Cosines", math.Pi/2))
	fmt.Println(myCalculator.doCalculation("Log", math.E))
}

func startCalculator() {
	myCalculator := Calculator{}
	myCalculator.addMathFunction(Sin{"Sine"})
	myCalculator.addMathFunction(Cos{"Cosine"})
	myCalculator.addMathFunction(Log{"Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if fName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")
}
