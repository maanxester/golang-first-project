
package main

import "fmt"
import "reflect"

func main() {

	//var nome = "Matheus"
	// Posso passar tanto o "var" para declarar uma variavel, ou passar := que também será uma variavel
	// Igual a de baixo

	 nome := "Matheus"


	 // Ficar atento quando for declarar um float, pois ele pode ser tanto float32 quanto float64
	 // Por padrão, claro não seja declarado, o Go irá automaticamente atribuir o float64
	 versao :=  1.2
	 fmt.Println("Hello World,", nome)
	 fmt.Println("Este programa está na versão", versao)


	 // função recleft.TypeOf servindo para informar qual tipo da variavel

	 fmt.Println("O tipo da minha varivavel versao é", reflect.TypeOf(versao))

	 fmt.Println("1 - Iniciar Monitoramento")
	 fmt.Println("2 - Exibir Logs")
	 fmt.Println("0 - Sair")

	 var comando int

	 //fmt.Scanf("%d", comando) ## Scanf é necessário especificar o tipo da minha variavel
	 fmt.Scan(&comando)

	 fmt.Println("O endereço de memoria da varial comando é", &comando)

	 fmt.Println("A escolha foi:", comando)

}
