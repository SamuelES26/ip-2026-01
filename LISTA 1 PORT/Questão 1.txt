programa 
{             inclua biblioteca Matematica --> mat
  funcao inicio() 
  {         
    real n1, n2, n3, media, produto

  escreva("Digite suas notas: \n")
  leia(n1, n2, n3)

  

  media = (n1 + n2 + n3) /3
  media = mat.arredondar(media, 2)
  


  escreva("Média: ", media, "\n") 
se(media > 6.0)escreva("APROVADO")
  se(media < 6.0)escreva("REPROVADO")
 
  }
}
