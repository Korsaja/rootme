package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"korsaj.io/rootme/programming/utils"
)

const startBufSize = 4096

func main() {
	address := flag.String("a", "", "server address")
	flag.Parse()

	if *address == "" {
		utils.Fatalf("empty address")
	}

	conn, err := net.DialTimeout("tcp", *address, time.Second*5)
	if err != nil {
		utils.Fatalf("connection %s", err)
	}

	var buf = make([]byte, startBufSize)
	for {

		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.Fatalf("read info %s", err)
		}

		encodeStr, err := utils.ParseEncodeString(utils.LastString(buf[:n]))
		if errors.Is(err, utils.ErrNotFound) {
			log.Printf("check input: %s", encodeStr)
			os.Exit(0)
		} else if err != nil {
			utils.Fatalf("parse encode string %s", err)
		}

		rot13 := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return 'A' + (r-'A'+13)%26
			case r >= 'a' && r <= 'z':
				return 'a' + (r-'a'+13)%26
			}
			return r
		}

		decode := strings.Map(rot13, encodeStr)

		if _, err = conn.Write(utils.AddLF([]byte(decode))); err != nil {
			utils.Fatalf("write conn %s", err)
		}
	}

	n, err := conn.Read(buf)
	if err != nil {
		utils.Fatalf("read info %s", err)
	}

	fmt.Printf("flag %s\n", buf[:n])

}
