package main

import (
	"fmt"

	"main/model"
)

func DepthFirstSearch(startGrid model.Grid, goalGrid model.Grid) {
	stack := []model.Grid{startGrid}
	visited := map[string]bool{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current.Key()] {
			continue
		}

		visited[current.Key()] = true

		if current.Equals(goalGrid) {
			fmt.Printf("Simulações realizadas: %d \n\n", len(visited))
			current.Tree()
			return
		}

		children := current.Children()
		for _, grid := range children {
			if !visited[grid.Key()] {
				stack = append(stack, grid)
			}
		}

	}
	fmt.Println("Resultado não encontrado.")
	fmt.Printf("Simulações realizadas: %d", len(visited))
}
