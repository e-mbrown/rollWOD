package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type CreateStmt struct {
	Token token.Token
	Stmt *ObjStmt
}

func(cs *CreateStmt) stmtNode(){}
func(cs *CreateStmt) TokenLiteral() []byte {return cs.Token.Literal}
func(cs *CreateStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(cs.TokenLiteral()))
	out.WriteString(cs.Stmt.String())


	return out.String()
}

// Todo: Expand
// type ObjTyp int

// const (
// 	User ObjTyp = iota
// 	Campaign
// 	Character
// )

// var ObjName = map[ObjTyp]string {
// 	User: "user",
// 	Campaign: "campaign",
// 	Character: "character",
// }

type ObjStmt struct {
	StmtType []byte
}

func(os *ObjStmt) String()string { return string(os.StmtType) }