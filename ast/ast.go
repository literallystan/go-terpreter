package ast

import "github.com/literallystan/go-terpreter/token"

//Node ...
type Node interface {
	TokenLiteral() string
}

//Statement ...
type Statement interface {
	Node
	statementNode()
}

//Expression ...
type Expression interface {
	Node
	expressionNode()
}

//Program ...
type Program struct {
	Statements []Statement
}

//TokenLiteral - return the statements TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}
