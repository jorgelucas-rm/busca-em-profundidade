package main

import (
	"fmt"

	"main/model"
)

func main() {
	startGrid := model.Grid{
		Position: []int{
			1, 2, 3,
			4, 5, 0, // Quadrado "Vazio"
			7, 8, 6,
		},
	}

	goalGrid := model.Grid{
		Position: []int{
			1, 2, 3,
			4, 5, 6,
			0, 7, 8, // Quadrado "Vazio"
		},
	}

	fmt.Println("Início da busca:\n")
	fmt.Println("Alvo:")
	goalGrid.Print()

	DepthFirstSearch(startGrid, goalGrid)
}
