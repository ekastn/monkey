package repl

import (
	"io"
	"strings"

	"github.com/ekastn/monkey/internal/evaluator"
	"github.com/ekastn/monkey/internal/lexer"
	"github.com/ekastn/monkey/internal/object"
	"github.com/ekastn/monkey/internal/parser"

	"github.com/reeflective/readline"
)

func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()

	rl := readline.NewShell()
	rl.Prompt.Primary(func() string { return ">> " })
	rl.AcceptMultiline = multiline

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserError(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserError(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "  "+msg+"\n")
	}
}

func multiline(line []rune) (accept bool) {
	if !strings.HasSuffix(string(line), ";") {
		return false
	}

	return true
}
