package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("não conheço comando")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	nome := "igor"
	versao := 1.1

	fmt.Println("ola senhor, ", nome)
	fmt.Println("entre programa esta na versao: ", versao)

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println(comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- sair do programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("deu erro", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "esta com problema", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLog() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
