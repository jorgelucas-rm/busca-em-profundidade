package main

import "fmt"

func main() {
    inicialCube := Lasanha{
        casas: []Casa{
            {coordenadaX: 0, coordenadaY: 0, valor: 1},
            {coordenadaX: 1, coordenadaY: 0, valor: 2},
            {coordenadaX: 2, coordenadaY: 0, valor: 3},
            {coordenadaX: 0, coordenadaY: 1, valor: 4},
            {coordenadaX: 1, coordenadaY: 1, valor: 5},
            {coordenadaX: 2, coordenadaY: 1, valor: 0},// Valor nulo
            {coordenadaX: 0, coordenadaY: 2, valor: 7},
            {coordenadaX: 1, coordenadaY: 2, valor: 8},
            {coordenadaX: 2, coordenadaY: 2, valor: 6},
        },
    }

    cuboTarget := Lasanha{
        casas: []Casa{
            {coordenadaX: 0, coordenadaY: 0, valor: 1},
            {coordenadaX: 1, coordenadaY: 0, valor: 2},
            {coordenadaX: 2, coordenadaY: 0, valor: 3},
            {coordenadaX: 0, coordenadaY: 1, valor: 4},
            {coordenadaX: 1, coordenadaY: 1, valor: 5},
            {coordenadaX: 2, coordenadaY: 1, valor: 6},
            {coordenadaX: 0, coordenadaY: 2, valor: 0},// Valor nulo
            {coordenadaX: 1, coordenadaY: 2, valor: 7},
            {coordenadaX: 2, coordenadaY: 2, valor: 8}, 
        },
    }

    fmt.Println("Início da busca:")
	fmt.Print("\n")
    fmt.Println("Cubo alvo:")
	imprimirLasanha(cuboTarget)

    movimentos, encontrado := BuscaEmProfundidade(inicialCube, cuboTarget)

    if encontrado {
        fmt.Printf("Cubo alvo encontrado em %d jogadas.", movimentos)
    } else {
        fmt.Println("Cubo alvo não encontrado.")
    }
}
