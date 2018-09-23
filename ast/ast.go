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

func (ls *LetStatement) statementNode() {}

//TokenLiteral returns, you guessed it. A token literal.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//Identifier ...
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

//ReturnStatement ...
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

//TokenLiteral ...
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
