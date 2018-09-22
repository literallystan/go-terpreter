package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/literallystan/go-terpreter/lexer"
	"github.com/literallystan/go-terpreter/token"
)

const PROMPT = ">> "

//Start - Start the repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
