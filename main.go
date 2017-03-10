package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func printToConn(conn net.Conn, message string, interpreted bool) {
	if interpreted {
		dst := make([]byte, hex.DecodedLen(len(message)))
		n, err := hex.Decode(dst, []byte(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dst[:n])
		conn.Write(dst[:n])
	} else {
		fmt.Println(conn)
		fmt.Fprintf(conn, message)
	}
}

func main() {
	BytePtr := flag.Bool("b", false, "Send as interpreted bytes")
	SpacePtr := flag.Bool("s", false, "Split after each space")

	flag.Parse()

	buff := bufio.NewReader(os.Stdin)

	if len(flag.Args()) < 1 {
		fmt.Println("You need to specify an address")
		os.Exit(0)
	}
	conn, err := net.Dial("tcp", flag.Args()[0])

	if err != nil {
		log.Fatal(err)
	}

	for {
		line, hasMoreInLine, err := buff.ReadLine()

		strLine := string(line)

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if *SpacePtr || *BytePtr {
			payload := strings.Split(strLine, " ")
			for _, v := range payload {
				printToConn(conn, v, *BytePtr)
			}
		} else {
			printToConn(conn, strLine, *BytePtr)
		}

		if hasMoreInLine || err == io.EOF {
			break
		}

	}

}
