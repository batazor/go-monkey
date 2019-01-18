package parser

import (
	"github.com/batazor/go-monkey/ast"
	"github.com/batazor/go-monkey/token"
)
import "github.com/batazor/go-monkey/lexer"

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}
