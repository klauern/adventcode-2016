package main

import (
	"fmt"
	"strconv"
	"strings"
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
	sides := strings.Split(s, " t")
	sd := make([]float64, 0)
	for _, v := range sides {
		fl, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(fmt.Sprintf("%v is not a float", sides[0]))
		}
		sd = append(sd, fl)
	}
	return sd
}
