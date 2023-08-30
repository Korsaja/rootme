package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"korsaj.io/rootme/programming/utils"
)

const startBufSize = 4096

func main() {
	address := flag.String("a", "", "server address")
	flag.Parse()

	if *address == "" {
		utils.Fatalf("empty address")
	}

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		utils.Fatalf("connection dial %s", err)
	}

	// first step read info
	var buf = make([]byte, startBufSize)
	n, err := conn.Read(buf)
	if err != nil {
		utils.Fatalf("read info %s", err)
	}

	f1, f2, err := parseNumber(utils.LastString(buf[:n]))
	if err != nil {
		utils.Fatalf("parse number %s", err)
	}

	res := fmt.Sprintf("%.2f", math.Sqrt(f1)*f2)

	if _, err = conn.Write(utils.AddLF([]byte(res))); err != nil {
		utils.Fatalf("write conn %s", err)
	}

	n, err = conn.Read(buf)
	if err != nil {
		utils.Fatalf("read flag %s", err)
	}

	fmt.Printf("flag: %s\n", buf[:n])
}

func parseNumber(s string, err error) (float64, float64, error) {
	if err != nil {
		return 0, 0, fmt.Errorf("input string %w", err)
	}
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return 0, 0, fmt.Errorf("not data to parse")
	}
	// find number ex 2
	var (
		first = true
		i, j  float64
	)
	for _, word := range words {
		n, err := strconv.ParseFloat(word, 64)
		if err == nil {
			if first {
				i = n
				first = false
			} else {
				j = n
				break
			}
		}
	}
	return i, j, nil
}
