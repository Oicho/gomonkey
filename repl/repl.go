package repl

import (
	"bufio"
	"fmt"
	"io"

	"gomonkey/lexer"
	"gomonkey/parser"
)

func MainLoop(in io.Reader, out io.Writer) {
	reader := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, "$ ")
		if !reader.Scan() {
			return
		}
		line := reader.Text()

		if line == "quit" {
			fmt.Fprintf(out, "GoodBye!")
			break
		} else {
			l := lexer.New(line)
			parser := parser.New(l)
			program, _ := parser.ParseProgram()
			s := program.String()
			fmt.Fprintf(out, "%+v\n", s)
		}
	}
}
