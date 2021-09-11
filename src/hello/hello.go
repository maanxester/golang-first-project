
package main

import (
	"fmt"
	"net/http"
	"os"
	//"reflect"
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



	// Quando a função retorna dois valores, é necessário passar duas variaveis
	// Exemplo: caso a função retorne a idade e o nome, vai ser necessário escrever:
	// nome, idade := devolveNomeEIdade()
	//nome := devolveNome()
	//fmt.Println(nome)


	slice()
	//array()
	exibeIntroducao()


	// Enquanto for True, esse for irá executar sempre
	for {
		exibeMenu()


		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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
}


// Caso eu queira retornar dois valores, é necessário eu colocar ao lado da função, o tipo dos dois valores
// Exemplo: func devolveNomeEIdade() (string, int) { code...

// Caso tenha uma função que retorne dois valores e eu queria apenas um único valor, posso fazer o seguinte:
// chamar a minha função que retorna dois valores, e quando eu declarar minha variavel, eu passo na seguinte forma:
// _, idade := NOMEDAFUNÇÃO
// devo passar este underline na variavel que não quero pegar.
//func devolveNome() string {
//	nome := "Matheus"
//	return nome
//}



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
	//fmt.Println("O endereço de memoria da varial comando é", &comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

// Array OBS: Não muito usado
// Dessa forma de array, é necessário passar a quantidade de arrays que serão listados, por exemplo como foi feito:
// var nomes [4] string
// Tamanho fixo
func array() {
	var nomes [4] string
		nomes[0] = "Jorge"
		nomes[1] = "Maria"
		nomes[3] = "Matheus"

		fmt.Println(nomes)
}


// Slice
// No Slice, não é necessário informar quantos itens serão listados na variavel
// Tamanho dinamico
// Porém, com capacidade dobrada
func slice() {
	nomes := [] string{"Matheus", "Jorge", "Lucas"}
	// Append
	nomes = append(nomes, "Maria")


	 // For sem range
	for i := 0 ; 1 < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	// For com o range
	for i, nome  := range nomes {
		fmt.Println("Na posição", i, "tem o nome:", nome)
	}

}

// função len(variavel) para consultar tamanho
// função cap(variavel) para consultar capacidade




 // Função que retorna o Get de uma página descobrindo se ela está on ou off
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.google.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println(site, "está no ar.")
	}else {
		fmt.Println(site, "não retornou o status de sucesso. Status Code:", resp.StatusCode)
	}
}
