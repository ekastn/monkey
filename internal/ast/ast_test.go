package ast

import (
	"testing"

	"github.com/eka-septian/monkey/internal/token"
)

func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "foo"},
                    Value: "foo",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "bar"},
                    Value: "bar",
                },
            },
        },
    }

    if program.Sring() != "let foo = bar;" {
        t.Errorf("program.String() wrong, got=%q", program.Sring())
    }
}
