package parser

import "github.com/batazor/go-monkey/token"
import "github.com/batazor/go-monkey/lexer"

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}
