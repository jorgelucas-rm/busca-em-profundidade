package main

import "fmt"

func troca(lasanha *Lasanha, pos1, pos2 int) {
    lasanha.casas[pos1], lasanha.casas[pos2] = lasanha.casas[pos2], lasanha.casas[pos1]
}

func imprimirLasanha(c Lasanha) {
    fmt.Println(c.casas[0].valor, c.casas[1].valor, c.casas[2].valor)
    fmt.Println(c.casas[3].valor, c.casas[4].valor, c.casas[5].valor)
    fmt.Println(c.casas[6].valor, c.casas[7].valor, c.casas[8].valor)
    fmt.Print("\n")
}

func lasanhaPresente(lasanha Lasanha, lista []Lasanha) bool {
    for _, item := range lista {
        if lasanhaIgual(lasanha, item) {
            return true
        }
    }
    return false
}

func gerarFilhos(lasanha Lasanha, explorados []Lasanha, listaEspera []Lasanha) []Lasanha {
    var novosFilhos []Lasanha

    var posNulo int
    for i, casa := range lasanha.casas {
        if casa.valor == 0 {
            posNulo = i
            break
        }
    }

    movimentos := []struct {
        dx, dy int
    }{
        {0, 1},  // baixo
        {1, 0},  // direita
        {0, -1}, // cima
        {-1, 0}, // esquerda
    }

    for _, mov := range movimentos {
        novoX := posNulo/3 + mov.dx
        novoY := posNulo%3 + mov.dy

        if novoX >= 0 && novoX <= 2 && novoY >= 0 && novoY <= 2 {
            novoPos := novoX*3 + novoY
            novaLasanha := Lasanha{casas: make([]Casa, len(lasanha.casas))}
            copy(novaLasanha.casas, lasanha.casas)
            troca(&novaLasanha, posNulo, novoPos)

            if !lasanhaPresente(novaLasanha, explorados) && !lasanhaPresente(novaLasanha, listaEspera) {
                novosFilhos = append(novosFilhos, novaLasanha)
            }
        }
    }

    fmt.Println("Filhos:")
    for _, filho := range novosFilhos {
        imprimirLasanha(filho)
    }
    return novosFilhos
}

func lasanhaIgual(l1,l2 Lasanha) bool {
    for i := range l1.casas {
        if l1.casas[i].valor !=l2.casas[i].valor {
            return false
        }
    }
    return true
}

func BuscaEmProfundidade(inicial, target Lasanha) (int, bool) {
    listaEspera := []Lasanha{inicial}
    explorados := []Lasanha{}
    comparacoes := 0

    for len(listaEspera) > 0 {
        comparacoes++
        atual := listaEspera[0]
        listaEspera = listaEspera[1:]

        fmt.Println("Lasanha atual:")
        imprimirLasanha(atual)

        if lasanhaIgual(atual, target) {
            return comparacoes, true
        }

        novosFilhos := gerarFilhos(atual, explorados, listaEspera)
        listaEspera = append(novosFilhos, listaEspera...)

        explorados = append(explorados, atual)
    }

    return comparacoes, false
}