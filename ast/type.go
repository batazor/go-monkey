package ast

import "github.com/batazor/go-monkey/token"

type Node interface {
	TokenLiteral() string
	String() string
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
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // the prefix token, e.g. !
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean struct {
	Token token.Token
	Value bool
}

type IfExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      token.Token // the `{` token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

type CallExpression struct {
	Token     token.Token // The `(` token
	Function  Expression  // Identifier or FunctionLitereal
	Arguments []Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}
