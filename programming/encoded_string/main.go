package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net"
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

	// first step read info
	var buf = make([]byte, startBufSize)
	n, err := conn.Read(buf)
	if err != nil {
		utils.Fatalf("read info %s", err)
	}

	encodeStr, err := utils.ParseEncodeString(utils.LastString(buf[:n]))
	if err != nil {
		utils.Fatalf("parse encode string %s", encodeStr)
	}

	bs64, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		utils.Fatalf("b64 decode %s", err)
	}

	if _, err = conn.Write(utils.AddLF(bs64)); err != nil {
		utils.Fatalf("write conn %s", err)
	}

	n, err = conn.Read(buf)
	if err != nil {
		utils.Fatalf("read info %s", err)
	}

	fmt.Printf("flag %s\n", buf[:n])
}
