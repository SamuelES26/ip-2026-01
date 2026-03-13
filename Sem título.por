programa {
inclua biblioteca Matematica --> mat

  funcao inicio() {
    real publico, pop, ger, arq, cad
    real v2, v3, v4, v5, div, mult
    inteiro v1

    escreva("Informe os valores: \n")
    leia(v1, v2, v3, v4, v5)
    publico = v1 
    pop = ((v2 /100) *v1) *1
    ger = ((v3 /100) *v1) *5
    arq = ((v4 /100) *v1) *10
    cad = ((v5 /100) *v1) *20

    mult = pop + ger + arq + cad

    mult = mat.arredondar(mult, 3)
  
    escreva("A RENDA DO JOGO FOI: ", mult)

   









  }
}
