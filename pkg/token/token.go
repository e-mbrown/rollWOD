package token

type TokenType string

type Token struct {
	Type TokenType
	Literal []byte
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	//IDENT & LITERALS
	IDENT = "IDENT"
	INT = "INT"
	FLOAT = "FLOAT"

	// KEYWORDS
	FUNCTION = "FN"
	DECLARE = "DCL"
	CREATE = "CREATE"
	
	//OP
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"

	// DELIMITER
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

