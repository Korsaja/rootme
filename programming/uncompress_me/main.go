package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"korsaj.io/rootme/programming/utils"
)

const startBufSize = 4096

var dialTimeout = 5 * time.Second

func main() {
	address := flag.String("a", "", "server address")
	flag.Parse()

	if *address == "" {
		utils.Fatalf("empty address")
	}

	conn, err := net.DialTimeout("tcp", *address, dialTimeout)
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
		}
		if err != nil {
			utils.Fatalf("parse encode string %s", err)
		}

		bs64, err := base64.StdEncoding.DecodeString(encodeStr)
		if err != nil {
			utils.Fatalf("dec bs64 %s", err)
		}
		r, err := zlib.NewReader(bytes.NewReader(bs64))
		if err != nil {
			utils.Fatalf("zlib reader %s", err)
		}

		var bb = new(bytes.Buffer)
		if _, err = io.Copy(bb, r); err != nil {
			utils.Fatalf("copy zlib %s", err)
		}

		_ = r.Close()
		if _, err = conn.Write(utils.AddLF(bb.Bytes())); err != nil {
			utils.Fatalf("write conn %s", err)
		}
	}

	n, err := conn.Read(buf)
	if err != nil {
		utils.Fatalf("read info %s", err)
	}

	fmt.Printf("flag %s\n", buf[:n])

}
