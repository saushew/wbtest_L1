package main

import "fmt"

func main() {
	tempretures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -0.1, 5.5}

	tempGroups := make(map[int][]float64)

	for _, t := range tempretures {
		g := int(t/10)*10
		if _, ok := tempGroups[g]; ok {
			tempGroups[g] = append(tempGroups[g], t)
		} else {
			tempGroups[g] = []float64{t}
		}
	}

	fmt.Println(tempGroups)
}
