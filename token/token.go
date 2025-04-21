package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"

  ASSIGN = "="
  PLUS = "+"
  COMMA = ","
  SEMICOLON = ";"
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

	MINUS = "-"
	BANG = "!"
	SLASH = "/"
	ASTHERISK = "*"

	GT = ">"
	LT = "<"

	//double
	EQ = "=="
	NOT_EQ = "!="

  // keywords
  FUNCTION = "FUNCTION" 
  LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	RETURN = "RETURN"
	IF = "IF"
	ELSE = "ELSE"
)

//defined token types (keywords)
var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
	"if": IF,
	"else": ELSE,	
}

func LookupIdent(ident string) TokenType {
  // if the given identifier is in the keyword list (is ok), return the given token
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT // if it's not, return type "IDENT"
}
