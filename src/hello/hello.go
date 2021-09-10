
package main

import (
	"fmt"
	"os"
)

func main() {

	//var nome = "Matheus"
	// Posso passar tanto o "var" para declarar uma variavel, ou passar := que também será uma variavel
	// Igual a de baixo

	 //nome := "Matheus"


	 // Ficar atento quando for declarar um float, pois ele pode ser tanto float32 quanto float64
	 // Por padrão, claro não seja declarado, o Go irá automaticamente atribuir o float64
	 //versao :=  1.2
	 //fmt.Println("Hello World,", nome)
	 //fmt.Println("Este programa está na versão", versao)

	 // função recleft.TypeOf servindo para informar qual tipo da variavel

	//fmt.Println("O tipo da minha varivavel versao é", reflect.TypeOf(versao))


	//fmt.Println("1 - Iniciar Monitoramento")
	// fmt.Println("2 - Exibir Logs")
	// fmt.Println("0 - Sair")

	//var comando int

	//fmt.Scanf("%d", comando) ## Scanf é necessário especificar o tipo da minha variavel
	//fmt.Scan(&comando) // Caso o tipo da variavel não seja o declarado, por padrão ele executa a primeira opção

	//fmt.Println("O endereço de memoria da varial comando é", &comando)

	//fmt.Println("A escolha foi:", comando)

	// Condiçoes if, else e switch

	// if comando == 1 {
	//	 fmt.Println("Monitorando...")
	// }else if comando == 2 {
	//	 fmt.Println("Exibindo Logs")
	//}else if comando == 0 {
	//	 fmt.Println("Saindo...")
	// }else {
	//	 fmt.Println("Erro")
	// }

	//-------------

	//var comando int = leComando()


	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Erro")
		os.Exit(-1) // -1 representa algo inesperado e informa ao S.O
	}
}

func exibeIntroducao() {
	nome := "Matheus"
	versao :=  1.2
	fmt.Println("Hello World,", nome)
	fmt.Println("Este programa está na versão", versao)
}

// Caso o objetivo seja que a função retorne algo, é necessário especificar o tipo do que será retornado

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço de memoria da varial comando é", &comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}
