package types

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Pos struct {
	Line   int
	Column int
}

func NewPos(token antlr.Token) Pos {
	return Pos{
		Line:   token.GetLine(),
		Column: token.GetColumn(),
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("(line:%d, col:%d)", p.Line, p.Column)
}

func (p Pos) IsValid() bool {
	return p.Line > 0
}

func (p Pos) Less(op Pos) bool {
	if p.Line == op.Line {
		return p.Column < op.Column
	}
	return p.Line < op.Line
}

func (p Pos) LessEq(op Pos) bool {
	return !op.Less(p)
}

type Use struct {
	StartPos Pos
	File     string
}

func (u Use) String() string {
	return fmt.Sprintf("(%s:%d:%d)", u.File, u.StartPos.Line, u.StartPos.Column)
}

const (
	ValueDeclaration    = "Value"
	FunctionDeclaration = "Function"
	ParamDeclaration    = "Param"
	ImportDeclaration   = "Import"
)

type Declaration struct {
	Name     string
	StartPos Pos
	File     string
	Parent   *Scope
	Inner    *Scope
	Type     string
	Uses     []Use
}

func (dcl *Declaration) String() string {
	return fmt.Sprintf("%s{name:%s, pos: %v, uses: %v}", dcl.Type, dcl.Name, dcl.StartPos, dcl.Uses)
}
