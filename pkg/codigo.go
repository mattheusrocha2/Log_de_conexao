package pkg

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func Menu() {
	for {
		fmt.Println("")
		fmt.Println(" -------------- Menu --------------")
		fmt.Println(" -------- Escolha uma Opção ----- ")
		fmt.Println("1 - Verificar Site ") //digita site, verifica se está online ou offline e cria um arquivo de log.
		fmt.Println("2 - Ler Log")         //Ler o arquivo de log.
		fmt.Println("0 - Sair ")           //sair do sistema.
		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Opção ", opcao, " Selecionada!")
			site()
		case 2:
			fmt.Println("Opção ", opcao, " Selecionada!")
		case 0:
			fmt.Println("Opção ", opcao, " Selecionada. Saindo do Sistema ...")
			os.Exit(0)
		default:
			fmt.Println("Opção Invalida!")
		}
		fmt.Println("")
	}
}

func site() {
	var https = "https://"
	var site string
	fmt.Print("Site: ")
	fmt.Scan(&site)
	siteDigitado := (https + site)
	fmt.Println(siteDigitado)

	response, erro := http.Get(siteDigitado)
	if erro != nil {
		fmt.Println("Ocorreu um erro: ", erro)
		MenuScape()
	}

	if response.StatusCode == 200 {
		fmt.Println("Site: " + siteDigitado + " Online!")
		//RETORNA O IP DO HOST
		ip, erro := net.LookupIP(site)
		if erro != nil {
			fmt.Print("Ocorreu um erro: ", erro)
		}
		fmt.Print("IP: ", ip)
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", siteDigitado, "Offline!")
		fmt.Println("Status Code: ", response.StatusCode)
		//RETORNA O IP DO HOST
		ip, erro := net.LookupIP(site)
		if erro != nil {
			fmt.Print("Ocorreu um erro: ", erro)
		}
		fmt.Print("IP: ", ip)
		registraLog(site, false)
	}

}

// MENU DE ESCAPE CASO OCORRA UM ERRO NO PRIMEIRO IF DA FUNÇÃO SITE
func MenuScape() {
	for {
		fmt.Println("")
		fmt.Println(" -------------- Menu --------------")
		fmt.Println(" -------- Escolha uma Opção ----- ")
		fmt.Println("1 - Verificar Site ") //digita site, verifica se está online ou offline e cria um arquivo de log.
		fmt.Println("2 - Ler Log")         //Ler o arquivo de log.
		fmt.Println("0 - Sair ")           //sair do sistema.
		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Opção ", opcao, " Selecionada!")
			site()
		case 2:
			fmt.Println("Opção ", opcao, " Selecionada!")
		case 0:
			fmt.Println("Opção ", opcao, " Selecionada. Saindo do Sistema ...")
			os.Exit(0)
		default:
			fmt.Println("Opção Invalida!")
		}
		fmt.Println("")
	}
}

// FUNÇÃO PARA CRIAR O ARQUIVO DE LOG E REGISTRAR OS LOGS
func registraLog(site string, status bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //FUNÇÃO QUE ABRE O ARQUIVO, CRIA QUANDO NÃO EXISTE E INSERI OS REGISTROS DENTRO.
	if erro != nil {
		fmt.Println("Ocorreu um erro: ", erro)
	}
	//MANDANDO ESCREVER O SITE COMO O STATUS TRUE OR FALSE - função strconv (converte o valor bool em string)
	//strconv.Formatbool(status)
	ipHost, _ := net.LookupIP(site)
	if status {
		var str = "Status Online"
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + str + " - IP Host: " + fmt.Sprint(ipHost, "\n"))

	} else {
		var str = "Status Offline"
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + str + " - IP Host: " + fmt.Sprint(ipHost, "\n"))
	}

}
