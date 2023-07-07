package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*- Utilizando goroutines, crie um programa incrementador:
  - Tenha uma variável com o valor da contagem
  - Crie um monte de goroutines, onde cada uma deve:
      - Ler o valor do contador
      - Salvar este valor
      - Fazer yield da thread com runtime.Gosched()
      - Incrementar o valor salvo
      - Copiar o novo valor para a variável inicial
  - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
  - Demonstre que há uma condição de corrida utilizando a flag -race

  --CRIANDO PROPOSITALMENTE UMA CONDIÇÃO DE CORRIDA!

*/

var wg sync.WaitGroup
var contador int = 0

func goroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println("Eu sou a goroutine número:", j, "E a contagem está em:", contador)
			wg.Done()
		}()
	}
}

func main() {
	goroutines(10)
	wg.Wait()

}
