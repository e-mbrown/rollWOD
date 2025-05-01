package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal []byte
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//IDENT & LITERALS
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	// KEYWORDS
	FUNCTION = "FUNC"
	DECLARE  = "DCL"
	CREATE   = "CREATE"
	IF = "IF"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
	RETURN = "RETURN"	

	//OP
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH = "/"
	LCARET = "<"
	RCARET = ">"
	BANG = "!"
	EQ = "=="
	N_EQ = "!="

	// DELIMITER
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

var keywords = map[string]TokenType{
	"func":  FUNCTION,
	"dcl": DECLARE,
	"if" : IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LookupIdent(lit []byte) TokenType {
	if tok, ok := keywords[string(lit)]; ok {
		return tok
	}

	return IDENT
}
