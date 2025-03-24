package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	var file *os.File
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		l, err := lexer.NewLexer(file)
		if err != nil {
			log.Fatal("Issue making the lexer")
		}

		l.TakeInput(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

		fmt.Println(scanner.Text())

	}
}