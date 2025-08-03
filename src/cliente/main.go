package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// antes de tudo deve-se atribuir uma porta para o cliente escutar
	// nesse caso, vamos deixar o SO decidir uma porta qualquer
	localAddr := net.UDPAddr{
		Port: 0,                      // 0 representa para o SO qualquer porta disponível
		IP:   net.ParseIP("0.0.0.0"), // dizemos ao SO que o cliente pode faze qualquer interface de conexão
	}

	conn, err := net.ListenUDP("udp", &localAddr)
	if err != nil {
		fmt.Println("Aconteceu um erro ao criar o cliente UDP:", err)
		return
	}
	defer conn.Close()

	// definindo o endereço do servidor
	serverAddr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"), // nesse caso vamos usar o localhost ouvindo nessa porta
	}

	reader := bufio.NewReader(os.Stdin)         // abrindo escutador
	CmdTextBytes, err := reader.ReadBytes('\n') // nao converte pasa string, já retorna o slice de bytes
	if err != nil {
		fmt.Println("Erro na leitura da mensagem", err)
		return
	}

	// enviando mensagem ao servidor
	if _, err := conn.WriteToUDP(CmdTextBytes, &serverAddr); err != nil {
		fmt.Println("Aconteceu um erro ao tentar enviar a mensagem", err)
		return
	}
	fmt.Println("Mensagem enviada ao servidor")

	// esperando resposta do servidor
	buffer := make([]byte, 2048)
	n, addr, err := conn.ReadFromUDP(buffer) // passar &buffer seria redundante como ja disse
	if err != nil {
		fmt.Println("Erro ao ler resposta do servidor", err)
		return
	}
	fmt.Printf("A resposta de %s foi: %s\n", addr, string(buffer[:n]))
}
