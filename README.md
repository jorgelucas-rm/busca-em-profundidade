# Busca em Profundidade (8-Puzzle Solver)

Implementação em Go do algoritmo de busca em profundidade (DFS) aplicado ao clássico quebra-cabeça de 8 peças (8-puzzle): um tabuleiro 3x3 com peças numeradas de 1 a 8 e um espaço vazio, onde o objetivo é alcançar um arranjo alvo movendo o espaço vazio.

## Como rodar

```bash
go run .
```

O estado inicial e o estado alvo são definidos em `main.go`. Edite os literais `startGrid` e `goalGrid` para testar outros arranjos.

> **Nota:** nem todo par (estado inicial, estado alvo) é alcançável. O 8-puzzle tem paridade: a partir de um estado inicial só é possível alcançar metade das permutações possíveis do tabuleiro. Se o alvo escolhido tiver paridade diferente da do estado inicial, a busca percorre todo o espaço de estados (até ~181.440 estados) e termina com "Resultado não encontrado."

## Estrutura

- `main.go` — define o estado inicial e o alvo, dispara a busca.
- `dfs.go` — implementação do algoritmo de busca em profundidade (`DepthFirstSearch`).
- `model/grid.go` — tipo `Grid`, representando um arranjo do tabuleiro como `Position []int` (o índice da slice codifica a posição: linha = `i/3`, coluna = `i%3`); métodos para comparar estados (`Equals`), gerar estados vizinhos (`Children`), gerar uma chave única por estado (`Key`) e reconstruir/imprimir o caminho até a solução (`Tree`, `Print`).
- `algoritmo.pdf` — material de referência sobre o algoritmo.

## Algoritmo

A busca usa uma pilha (LIFO) e um conjunto de estados visitados (`map[string]bool`, chaveado por `Grid.Key()`) para não reprocessar o mesmo arranjo duas vezes. Cada estado guarda um ponteiro (`Dad`) para o estado que o originou, permitindo reconstruir a sequência de jogadas até a solução quando ela é encontrada.
