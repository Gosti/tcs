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

func main() {
	BytePtr := flag.Bool("b", false, "Send as interpreted bytes")

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

		r := make([]byte, 1024)

		n, err := conn.Read(r)
		fmt.Println("Reply =>", r[0:n])

		if hasMoreInLine || err == io.EOF {
			break
		}

	}

}
