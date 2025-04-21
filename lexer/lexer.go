package lexer

import (
  "intbook/token"
)

type Lexer struct {
  input string
  position int // last char position
  readPosition int // actual char position
  ch byte // last char value
}

//create a new lexer, set it the input (plain text) and first char
func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
} 

// read actual char and setup next reading position
func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) { 
    l.ch = 0 //if finished input, char value is 0 (nul) 
  } else {
    l.ch = l.input[l.readPosition] // set actual char value 
  }

  l.position = l.readPosition // last char becomes actual char 
  l.readPosition += 1 // actual char is the next one
} 

//reads the actual value of l.ch, prepares the next char, and returns a token accordingly
func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  l.skipWhitespace()

  switch l.ch {
    case '=':
			if l.peekChar() == '=' {
				firstCh := l.ch
				l.readChar() //move on to the next char 
				tok = token.Token{Type: token.EQ, Literal: string(firstCh) + string(l.ch)}
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
    case ';':
      tok = newToken(token.SEMICOLON, l.ch)
    case '(':
      tok = newToken(token.LPAREN, l.ch)
    case ')':
      tok = newToken(token.RPAREN, l.ch)
    case '{':
      tok = newToken(token.LBRACE, l.ch)
    case '}':
      tok = newToken(token.RBRACE, l.ch)
    case ',':
      tok = newToken(token.COMMA, l.ch)
    case '+':
      tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' {
				firstCh := l.ch
				l.readChar()
				tok = token.Token{Type: token.NOT_EQ, Literal: string(firstCh) + string(l.ch)} 
			} else {
				tok = newToken(token.BANG, l.ch)
			}
		case '*':
			tok = newToken(token.ASTHERISK, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
    case 0:
      tok.Literal = ""
      tok.Type = "EOF"
    default: //if it's none of those, check if it's a letter
      if isLetter(l.ch) {
        tok.Literal = l.readIdentifier() //if it is, read the whole identifier and set it as the literal of the token
        tok.Type = token.LookupIdent(tok.Literal) //check if it is a keyword

        return tok
      } else if isDigit(l.ch) {
        tok.Type = token.INT
        tok.Literal = l.readNumber()
        return tok
      } else {
        tok = newToken(token.ILLEGAL, l.ch) //if it's not, it's an illegal char
      }
  }

  l.readChar()
  return tok
 
}
  
//initialize tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

//got a letter, look for an identifier
func (l *Lexer) readIdentifier() string {
  initial := l.position //save the initial letter's position of the identifier

  for isLetter(l.ch) {
    l.readChar() //for every letter, read the char, if not a letter it stops
  } 

  return l.input[initial:l.position]
}

//same as with identifiers
func (l *Lexer) readNumber() string {
  initial := l.position
  for isDigit(l.ch) {
    l.readChar()
  }

  return l.input[initial:l.position]
}

//this only function expresses all the symbols that will be able to be in all identifiers 
func isLetter(ch byte) bool { //if it's between a to z or it's a _ return true
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' 
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar() //go to the next char
  }
}

//checks for the next token and returns it, it enables two-character tokens
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 
	} else {
		return l.input[l.readPosition]
	}
}
