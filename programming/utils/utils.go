package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func AddLF(b []byte) []byte {
	if len(b) > 0 && b[len(b)-1] != '\n' {
		return append(b, '\n')
	}
	return b
}

func Fatalf(format string, a ...any) {
	fmt.Printf("Error: %s\n", fmt.Sprintf(format, a...))
	os.Exit(1)
}

func LastString(b []byte) (string, error) {
	rb := bytes.NewReader(b)
	sc := bufio.NewScanner(rb)

	content := make([]string, 0)
	for sc.Scan() {
		content = append(content, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("scan err: %w", err)
	}

	if len(content) == 0 {
		return "", fmt.Errorf("no content")
	}

	return content[len(content)-1], nil
}

func ParseEncodeString(s string, err error) (string, error) {
	if err != nil {
		return "", fmt.Errorf("input string %w", err)
	}

	i := strings.IndexByte(s, byte('\''))
	j := strings.IndexByte(s, byte('.'))
	if i == -1 || j == -1 {
		return s, ErrNotFound
	}
	return s[i+1 : j-1], nil
}
