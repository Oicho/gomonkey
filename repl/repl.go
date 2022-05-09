package repl

import (
	"bufio"
	"fmt"
	"io"

	"gomonkey/lexer"
	"gomonkey/token"
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
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Fprintf(out, "%+v\n", tok)
			}
		}
	}
}
