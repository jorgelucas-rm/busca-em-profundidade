package main

import "fmt"

type Lasanha struct {   //O nome é lasanha porque é uma # (Tenho um professor que chama assim)
    casas []Casa        //Antes eu chamava de cubo, mas não faz sntido porque cubo tem 6 lados
}

func (l Lasanha) String() string {
    return fmt.Sprintf("%v", l.casas)
}
