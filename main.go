package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
	WaitPtr := flag.Bool("w", false, "Wait for a response after each message")

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

		if *BytePtr {
			payload := strings.Split(strLine, " ")

			var data []byte

			for _, v := range payload {
				dst := make([]byte, hex.DecodedLen(len(v)))
				n, err := hex.Decode(dst, []byte(v))
				if err != nil {
					log.Fatal(err)
				}
				data = append(data, dst[:n]...)
			}

			conn.Write(data)
		} else {
			fmt.Println(conn)
			fmt.Fprintf(conn, strLine)
		}

		if *WaitPtr {
			fmt.Println(ioutil.ReadAll(conn))
		}

		if hasMoreInLine || err == io.EOF {
			break
		}

	}

}
