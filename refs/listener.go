package refs

import (
	"fmt"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"lsif-jsonnet/parser"
	"lsif-jsonnet/types"
)

type Listener struct {
	pathResolver *PathResolver
	*parser.BaseJsonnetListener
	file         string
	dcls         []*types.Declaration
	imports      []*Listener
	currentScope *types.Scope
	fileScope    *types.Scope
	errs         []error
}

func ParseFile(path string, pathResolver *PathResolver) (*Listener, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ParseData(data, path, pathResolver)
}

func ParseData(data []byte, path string, pathResolver *PathResolver) (*Listener, error) {
	is := antlr.NewInputStream(string(data))

	lexer := parser.NewJsonnetLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJsonnetParser(stream)

	ast := p.Jsonnet()

	startPos, endPos := types.NewPos(ast.GetStart()), types.NewPos(ast.GetStop())

	ll := NewListener(path, startPos, endPos, pathResolver)

	// Finally walk the tree
	antlr.ParseTreeWalkerDefault.Walk(ll, ast)

	return ll, nil
}

func NewListener(file string, pos, end types.Pos, pathResolver *PathResolver) *Listener {
	scope := types.NewScope(types.Universe, pos, end)
	ll := &Listener{
		currentScope: scope,
		fileScope:    scope,
		file:         file,
		pathResolver: pathResolver,
	}
	return ll
}

func (ll *Listener) File() string {
	return ll.file
}

func (ll *Listener) Declarations() []*types.Declaration {
	return ll.dcls
}

func (ll *Listener) Imports() []*Listener {
	return ll.imports
}

func (ll *Listener) EnterFunctionBind(ctx *parser.FunctionBindContext) {
	id := ctx.ID()
	startPos, endPos := types.NewPos(ctx.GetStart()), types.NewPos(ctx.GetStop())

	obj := ll.addDeclaration(id.GetSymbol(), types.FunctionDeclaration)

	scope := types.NewScope(ll.currentScope, startPos, endPos)
	ll.currentScope = scope
	obj.Inner = scope
}

func (ll *Listener) ExitFunctionBind(*parser.FunctionBindContext) {
	ll.currentScope = ll.currentScope.Parent()
}

func (ll *Listener) EnterFunctionField(ctx *parser.FunctionFieldContext) {
	fieldName := ctx.Fieldname()

	switch v := fieldName.(type) {
	case *parser.FieldnameContext:
		startPos, endPos := types.NewPos(ctx.GetStart()), types.NewPos(ctx.GetStop())

		obj := ll.addDeclaration(v.ID().GetSymbol(), types.FunctionDeclaration)

		scope := types.NewScope(ll.currentScope, startPos, endPos)
		ll.currentScope = scope
		obj.Inner = scope
	default:
	}
}

// ExitFunctionField is called when production FunctionField is exited.
func (ll *Listener) ExitFunctionField(*parser.FunctionFieldContext) {
	ll.currentScope = ll.currentScope.Parent()
}

func stripQuotes(str string) string {
	n := len(str)
	return str[1 : n-1]
}

func (ll *Listener) EnterValueBind(ctx *parser.ValueBindContext) {
	id := ctx.ID()
	startPos, endPos := types.NewPos(ctx.GetStart()), types.NewPos(ctx.GetStop())

	obj := ll.addDeclaration(id.GetSymbol(), types.ValueDeclaration)

	scope := types.NewScope(ll.currentScope, startPos, endPos)
	ll.currentScope = scope

	rhs := ctx.Expr()
	switch v := rhs.(type) {
	case *parser.ImportContext:
		obj.Type = types.ImportDeclaration
		path := stripQuotes(v.STRING().GetText())

		resolvedPath, err := ll.pathResolver.Resolve(path)
		if err != nil {
			obj.Inner = scope
			ll.errs = append(ll.errs, fmt.Errorf("failed to resolve import %s: %w", path, err))
		} else {
			ill, err := ParseFile(resolvedPath, ll.pathResolver)
			if err != nil {
				obj.Inner = scope
				ll.errs = append(ll.errs, fmt.Errorf("failed to parse %s: %w", resolvedPath, err))
			} else {
				ll.imports = append(ll.imports, ill)
				obj.Inner = ill.fileScope
			}
		}
	default:
		obj.Inner = scope
	}
}

func (ll *Listener) addDeclaration(token antlr.Token, declType string) *types.Declaration {
	obj := &types.Declaration{
		Name:     token.GetText(),
		File:     ll.file,
		StartPos: types.NewPos(token),
		Parent:   ll.currentScope,
		Type:     declType,
	}

	ll.dcls = append(ll.dcls, obj)

	ll.currentScope.Insert(obj)
	return obj
}

func (ll *Listener) ExitValueBind(*parser.ValueBindContext) {
	ll.currentScope = ll.currentScope.Parent()
}

func (ll *Listener) EnterVar(ctx *parser.VarContext) {
	id := ctx.ID()
	pos := types.NewPos(id.GetSymbol())
	_, decl := ll.currentScope.LookupParent(id.GetText(), pos)

	if decl != nil {
		decl.Uses = append(decl.Uses, types.Use{StartPos: pos, File: ll.file})
	}
}

func (ll *Listener) EnterParams(ctx *parser.ParamsContext) {
	for _, dcl := range ctx.GetPos() {
		ll.addDeclaration(dcl, types.ParamDeclaration)
	}
	for _, dcl := range ctx.GetNames() {
		ll.addDeclaration(dcl, types.ParamDeclaration)
	}
}

func (ll *Listener) EnterFunction(ctx *parser.FunctionContext) {
	startPos, endPos := types.NewPos(ctx.GetStart()), types.NewPos(ctx.GetStop())

	scope := types.NewScope(ll.currentScope, startPos, endPos)
	ll.currentScope = scope
}

func (ll *Listener) ExitFunction(*parser.FunctionContext) {
	ll.currentScope = ll.currentScope.Parent()
}

func (ll *Listener) EnterIndex(ctx *parser.IndexContext) {
	id := ctx.ID()
	pos := types.NewPos(id.GetSymbol())
	expr := ctx.Expr()

	switch v := expr.(type) {
	case *parser.VarContext:
		tt := v.ID()
		_, decl := ll.currentScope.LookupParent(tt.GetText(), types.NewPos(tt.GetSymbol()))
		if decl != nil && decl.Inner != nil {
			fieldDecl := decl.Inner.Lookup(id.GetText())

			if fieldDecl != nil {
				fieldDecl.Uses = append(fieldDecl.Uses, types.Use{StartPos: pos, File: ll.file})
			}
		}
	default:
	}
}
