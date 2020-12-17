package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cube struct {
	x int
	y int
	z int
}

func readActiveCubes(file string) map[cube]bool {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	activeCubes := make(map[cube]bool)
	for x, line := range lines {
		for y, state := range line {
			if state == '#' {
				activeCubes[cube{x, y, 0}] = true
			}
		}
	}
	return activeCubes
}

func countNeighbors(activeCubes map[cube]bool, currentCube cube) int {
	counter := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				// Skip itself
				if i == 0 && j == 0 && k == 0 {
					continue
				}

				searchCube := cube{currentCube.x + i, currentCube.y + j, currentCube.z + k}
				if activeCubes[searchCube] {
					counter++
				}
			}
		}
	}
	return counter
}

func cycle(activeCubes map[cube]bool, cycles int) map[cube]bool {
	var output map[cube]bool
	for cycleCount := 0; cycleCount < cycles; cycleCount++ {
		output = copyMap(activeCubes)
		// Handle active cubes
		for c := range activeCubes {
			activeNeighbors := countNeighbors(activeCubes, c)
			if activeNeighbors <= 1 || activeNeighbors >= 4 {
				delete(output, c)
			}
		}
		// Handle inactive cubes
		for c := range activeCubes {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					for k := -1; k <= 1; k++ {
						searchCube := cube{c.x + i, c.y + j, c.z + k}
						// The current cube is active so skip it
						if activeCubes[searchCube] {
							continue
						}
						activeNeighbors := countNeighbors(activeCubes, searchCube)
						if activeNeighbors == 3 {
							output[searchCube] = true
						}
					}
				}
			}
		}
		activeCubes = output
	}
	return output
}

func copyMap(original map[cube]bool) map[cube]bool {
	copy := make(map[cube]bool)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func main() {
	activeCubes := readActiveCubes("input.txt")
	activeCubes = cycle(activeCubes, 6)
	fmt.Println(len(activeCubes))
}
