package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Function that works with every type which implements interface io.Reader
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') {
				out[string(b)]++
			}

		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}

}

func main() {

	// Since function is defined for interfaces we can use it on strings
	s := "The quick brown fox jumped over the lazy dogs"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counts)

	// Since function is defined for interfaces we can also use it on files
	f, err := os.Open("main.go")
	defer func() { f.Close() }()
	if err != nil {
		fmt.Println(err)
		return
	}
	counts, err = countLetters(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counts)

}
