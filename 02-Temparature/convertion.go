package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Measurement interface {
	String() string
}

type Temperature struct {
	CelciusT    float64
	FahrenheitT float64
	KelvinT     float64
}

func FromCelcius(c float64) Temperature {
	return Temperature{c, (c * 9 / 5) + 32, c + 273.15}
}

func FromFarenheit(f float64) Temperature {
	return Temperature{(f + 32) * 9 / 5, f, (f + 459.67) * 5 / 9}
}

func FromKelvin(k float64) Temperature {
	return Temperature{k - 273.15, k*9/5 - 459.67, k}
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.3gC = %.3gF = %.3gK ", t.Celcius(), t.Farenheit(), t.Kelvin())
}

func (t Temperature) Celcius() float64 {
	return float64(t.CelciusT)
}

func (t Temperature) Farenheit() float64 {
	return float64(t.FahrenheitT)
}

func (t Temperature) Kelvin() float64 {
	return float64(t.KelvinT)
}

func newMeasurement(f float64, unit string) (Measurement, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "c":
		return FromCelcius(f), nil
	case "f":
		return FromFarenheit(f), nil
	case "k":
		return FromKelvin(f), nil
	default:
		return nil, fmt.Errorf("Unexpected unit %v", unit)
	}

}

func printMeasurement(s string) {
	re := regexp.MustCompile(`(-?\d+(?:\.\d+)?)([A-Za-z]+)`)
	match := re.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("Expecting <number><unit>, got %q", s)
	}
	f, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatalf("%v isn't a number.", match[1])
	}
	if match[2] == "" {
		log.Fatalf("No unit specified.")
	}
	unit := match[2]
	m, err := newMeasurement(f, unit)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(m)
}
