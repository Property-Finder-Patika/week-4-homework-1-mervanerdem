package main

import (
	"math"
	"math/rand"
	"time"
)

type MathFunction interface {
	Calculate(arg float64) float64
	GetName() string
}

type Sin struct {
	Name string
}

func (s Sin) Calculate(arg float64) float64 {
	return math.Sin(arg)
}

func (s Sin) GetName() string {
	return s.Name
}

type Cos struct {
	Name string
}

func (c Cos) Calculate(arg float64) float64 {
	return math.Cos(arg)
}

func (c Cos) GetName() string {
	return c.Name
}

type Log struct {
	Name string
}

func (l Log) Calculate(arg float64) float64 {
	return math.Log(arg)
}

func (l Log) GetName() string {
	return l.Name
}

func MathFunctionFactory() MathFunction {
	var mf MathFunction
	rand.Seed(time.Now().UnixNano())
	i := rand.Int31n(3)

	switch i {
	case 0:
		mf = Sin{"Sinus"}
	case 1:
		mf = Cos{"Cosines"}
	case 2:
		mf = Log{"Log"}
	}
	return mf
}
