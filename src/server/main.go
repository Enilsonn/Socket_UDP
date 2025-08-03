package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// definindo o endereço IP do Servidor
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}
	// IP -> 0.0.0.0 executa todas as interfaces disponíveis: internet, ethernet, localhost...

	// abrindo o socket UDP
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("O Servidor não pôde ser iniciado", err)
		return
	}
	defer conn.Close()

	fmt.Println("Servidor UDP está sendo executado na porta 8080")

	// declarando buffer de bytes do servidor
	buffer := make([]byte, 2048)

	// loop infinito de execeção do servidor
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer) // o buffer por si só já é um ponteiro
		if err != nil {
			fmt.Println("Aconteceu um arro ao tentar ler:", err)
			return
		}

		//messagem recebida (slice de bytes) -> string -> stringUpperCase -> (slice de bytes)
		bufferUpper := strings.ToUpper(string(buffer[:n]))
		response := []byte(bufferUpper)

		conn.WriteToUDP(response, remoteAddr)
	}
}
