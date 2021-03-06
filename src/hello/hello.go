
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	//"io/ioutil"
	"net/http"
	"os"
	"time"
	//"reflect"
	"bufio"
)

const monitoramentoSite = 5 // Constante são "variáveis" imutáveis.
const delay = 5



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


	//slice()
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
			imprimeLogs()
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
//func slice() {
//	nomes := [] string{"Matheus", "Jorge", "Lucas"}
//	// Append
//	nomes = append(nomes, "Maria")
//
//
//	 // For sem range
//	for i := 0 ; 1 < len(nomes); i++ {
//		fmt.Println(nomes[i])
//	}
//
//	// For com o range
//	for i, nome  := range nomes {
//		fmt.Println("Na posição", i, "tem o nome:", nome)
//	}
//
//}

// função len(variavel) para consultar tamanho
// função cap(variavel) para consultar capacidade




 // Função que retorna o Get de uma página descobrindo se ela está on ou off
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://www.google.com.br", "https://www.facebook.com/5324432",
	//	"https://www.alura.com.br"}

	sites := leSitesArquivo()

	for i := 0; i < monitoramentoSite ; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second) //
	}
}

func testaSite(site string) { // <- Dessa forma se passa um dado que será recebido na chamada da função
	resp, err := http.Get(site)

	if err != nil {
		log.Fatal("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "está no ar.")
		registraLog(site, true)
	}else {
		fmt.Println(site, "não retornou o status de sucesso. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}


func leSitesArquivo() []string {

	var sites []string

	//arquivo, err := os.Open("sites.txt") // Ficar atento, sempre é necessário tratar os erros
	//arquivo, err := ioutil.ReadFile("sites.txt") // Retorna um array de bytes, sendo necessário convertê-lo

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		log.Fatal("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {											// \n representa uma quebra de linha dentro do arquivo .txt
		linha, err := leitor.ReadString('\n') // Transformará o conteudo armazenado no endereço de memória em uma string.
		linha = strings.TrimSpace(linha) // Quando já tenho uma variável declarada, e quero chamar ela, não é necessário
										// declarar com o := novamente.
		sites = append(sites, linha)

		if err == io.EOF {
			fmt.Println("O leitor chegou no fim da linha")
			break
		}
	}
	//fmt.Println(string(arquivo)) // Realizando a conversão do array de bytes

	arquivo.Close() // Boa prática, pois, foi utilizado o os.Open, e é necessário fecha-lo para ser mais amigavel com o OS.


	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2016 15:04:06") + " - " + site + "- online: " +
		strconv.FormatBool(status) + "\n") // Biblioteca strconv que basicamente faz a conversão para qualquer tipo

	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		log.Fatal("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}