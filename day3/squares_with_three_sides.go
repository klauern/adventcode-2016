package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

func LargestSide(sides []float64) float64 {
	largest := -1.0
	for _, v := range sides {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func canTriangle(x, y, z float64) bool {
	return x < y+z
}

func IsTriangle(sides []float64) bool {
	largest := LargestSide(sides)
	if sides[0] == largest {
		return canTriangle(sides[0], sides[1], sides[2])
	}
	if sides[1] == largest {
		return canTriangle(sides[1], sides[0], sides[2])
	}
	if sides[2] == largest {
		return canTriangle(sides[2], sides[0], sides[1])
	}
	return false
}

func NewSide(s string) []float64 {
	fmt.Printf("Parsing line %v\n", s)
	sides := strings.Fields(s)
	sd := make([]float64, 0)
	fmt.Printf("Length of ary %v\n", len(sides))
	for _, v := range sides {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(fmt.Sprintf("%v is not an integer", v))
		}
		sd = append(sd, n)
	}
	return sd
}

func main() {
	input := helpers.MustLoadFile("input.txt")
	input = strings.TrimSpace(input)
	input = strings.Trim(input, " \r\n\t")
	inputs := strings.Split(input, "\n")
	var possibles int
	for _, v := range inputs {
		fmt.Printf("Length of %v is %v\n", v, len(v))
		s := NewSide(v)
		if IsTriangle(s) {
			possibles++
		}
	}
	fmt.Printf("There are %v possible triangles", possibles)
}
