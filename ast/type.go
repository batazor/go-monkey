package ast

import "github.com/batazor/go-monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

type Identifier struct {
	Token token.Token // the token.EDENT token
	Value string
}
