package model

import (
	"fmt"
	"slices"
	"strings"
)

type Grid struct {
	Dad      *Grid
	Position []int
}

func (g Grid) String() string {
	return fmt.Sprintf("%v", g.Position)
}

func (g Grid) Print() {
	fmt.Println(g.Position[0], g.Position[1], g.Position[2])
	fmt.Println(g.Position[3], g.Position[4], g.Position[5])
	fmt.Println(g.Position[6], g.Position[7], g.Position[8])
	fmt.Print("\n")
}

func (g *Grid) Switch(pos1, pos2 int) {
	g.Position[pos1], g.Position[pos2] = g.Position[pos2], g.Position[pos1]
}

func (g Grid) Equals(grid Grid) bool {
	for i := range g.Position {
		if g.Position[i] != grid.Position[i] {
			return false
		}
	}
	return true
}

func (g Grid) In(list []Grid) bool {
	return slices.ContainsFunc(list, g.Equals)
}

func (g Grid) Key() string {
	var sb strings.Builder
	for _, v := range g.Position {
		sb.WriteByte(byte('0' + v))
	}
	return sb.String()
}

func (g Grid) Tree() {
	var caminho []Grid
	current := g

	for {
		caminho = append(caminho, current)
		if current.Dad == nil {
			break
		}
		current = *current.Dad
	}

	slices.Reverse(caminho)

	fmt.Println("Jogadas realizadas:\n")
	for _, grid := range caminho {
		grid.Print()
	}
}

func (g *Grid) Children() []Grid {
	var children []Grid

	var posNil int
	for i, value := range g.Position {
		if value == 0 {
			posNil = i
			break
		}
	}

	directions := []struct {
		x, y int
	}{
		{0, 1},  // Direita
		{-1, 0}, // Cima
		{0, -1}, // Esquerda
		{1, 0},  // Baixo
	}

	for _, dir := range directions {
		x := posNil/3 + dir.x
		y := posNil%3 + dir.y

		if x >= 0 && x <= 2 && y >= 0 && y <= 2 {
			novoPos := x*3 + y
			newGrid := Grid{Dad: g, Position: make([]int, len(g.Position))}
			copy(newGrid.Position, g.Position)
			newGrid.Switch(posNil, novoPos)

			children = append(children, newGrid)
		}
	}
	return children
}
