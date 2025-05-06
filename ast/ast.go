package ast


// expressions return values, identifiers don't
// let <identifier> = <expression>

import (
	"intbook/token"
)

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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // {Type: LET, Literal: "let"}
	Name *Identifier 
	Value Expression
}

func (ls *LetStatement) statementNode() {} //make ls a statementNode
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token // {Type: IDENT, Literal: "variable name"}
	Value string	
}

func (i *Identifier) expressionNode() {} //make our identifier an expression, because when called it will return a value
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}
