package token

type TokenType string

const (
	EOF = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	// Key words
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNC"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	DIVIDE   = "/"
	MULTIPLY = "*"
	BANG     = "!"
	ASSIGN   = "="

	// Comparator
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// Delimiters
	COMMA      = ","
	SEMICOLLON = ";"
	LPAREN     = "("
	RPAREN     = "("
	LBRACE     = "{"
	RBRACE     = "}"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"if":     IF,
	"else":   ELSE,
	"func":   FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
