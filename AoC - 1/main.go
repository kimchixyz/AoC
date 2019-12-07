package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Testing")

	file, err := os.Open("input.log")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fuel float64
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			var fuelForFuel float64
			// sum += calcFuel(mass) -- PART 1

			// -- PART 2
			for mass > 0 {
				mass = calcFuel(mass)
				if mass > 0 {
					fuelForFuel += mass
				}
			}

			fuel += fuelForFuel

		}

	}
	fmt.Println(fuel)
}

func calcFuel(mass float64) float64 {

	var result float64

	result = math.Floor(mass/3 - 2)

	return result
}
