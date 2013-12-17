package transformer

import (
	"bufio"
	"io"
)

type Transformer interface {
	ProcessLine(line string) error
	WriteTransformed(out io.Writer) error
}

func NewRunner(in io.Reader, out io.Writer, t Transformer) Runner {
	return Runner{in, out, t}
}

type Runner struct {
	in  io.Reader
	out io.Writer
	t   Transformer
}

func (r Runner) Transform() error {
	var err error
	var line string
	bIn := bufio.NewReader(r.in)

	for err == nil {
		if line, err = bIn.ReadString('\n'); err == nil {
			err = r.t.ProcessLine(line)
		}
	}

	if err != io.EOF {
		return err
	}

	return r.t.WriteTransformed(r.out)
}
