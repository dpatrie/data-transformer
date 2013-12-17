package main

import (
	"fmt"
	"github.com/dpatrie/data-transformer/transformer"
	"io"
	"os"
	"strings"
)

func main() {
	r := transformer.NewRunner(os.Stdin, os.Stdout, &WordCount{})
	if err := r.Transform(); err != nil {
		fmt.Println(err)
	}
}

type WordCount struct {
	words map[string]int
}

func (s *WordCount) ProcessLine(line string) error {
	if s.words == nil {
		s.words = make(map[string]int)
	}
	key := strings.Trim(line, "\n")
	s.words[key]++

	return nil
}

func (s *WordCount) WriteTransformed(w io.Writer) error {
	for k, v := range s.words {
		if _, err := fmt.Fprintf(w, "%s\t%d\n", k, v); err != nil {
			return err
		}
	}
	return nil
}
