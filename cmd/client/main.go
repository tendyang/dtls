package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/pions/dtls/internal/ice"
	"github.com/pions/dtls/pkg/dtls"
)

const bufSize = 8192

func main() {
	a, _ := ice.Listen("127.0.0.1:5555", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 4444})

	dtlsConn, err := dtls.Dial(a /* localCertificate */, nil /* localPrivateKey */, nil)
	check(err)
	defer dtlsConn.Close()

	go func() {
		b := make([]byte, bufSize)
		for {
			n, err := dtlsConn.Read(b)
			check(err)
			fmt.Printf("Got message: %s\n", string(b[:n]))
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
