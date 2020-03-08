// Code generated from Jsonnet.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Jsonnet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJsonnetListener is a complete listener for a parse tree produced by JsonnetParser.
type BaseJsonnetListener struct{}

var _ JsonnetListener = &BaseJsonnetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonnetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonnetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonnetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonnetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJsonnet is called when production jsonnet is entered.
func (s *BaseJsonnetListener) EnterJsonnet(ctx *JsonnetContext) {}

// ExitJsonnet is called when production jsonnet is exited.
func (s *BaseJsonnetListener) ExitJsonnet(ctx *JsonnetContext) {}

// EnterImport is called when production Import is entered.
func (s *BaseJsonnetListener) EnterImport(ctx *ImportContext) {}

// ExitImport is called when production Import is exited.
func (s *BaseJsonnetListener) ExitImport(ctx *ImportContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseJsonnetListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseJsonnetListener) ExitParens(ctx *ParensContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseJsonnetListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseJsonnetListener) ExitVar(ctx *VarContext) {}

// EnterApply is called when production Apply is entered.
func (s *BaseJsonnetListener) EnterApply(ctx *ApplyContext) {}

// ExitApply is called when production Apply is exited.
func (s *BaseJsonnetListener) ExitApply(ctx *ApplyContext) {}

// EnterBinaryExpr is called when production BinaryExpr is entered.
func (s *BaseJsonnetListener) EnterBinaryExpr(ctx *BinaryExprContext) {}

// ExitBinaryExpr is called when production BinaryExpr is exited.
func (s *BaseJsonnetListener) ExitBinaryExpr(ctx *BinaryExprContext) {}

// EnterApplyBrace is called when production ApplyBrace is entered.
func (s *BaseJsonnetListener) EnterApplyBrace(ctx *ApplyBraceContext) {}

// ExitApplyBrace is called when production ApplyBrace is exited.
func (s *BaseJsonnetListener) ExitApplyBrace(ctx *ApplyBraceContext) {}

// EnterIndex is called when production Index is entered.
func (s *BaseJsonnetListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production Index is exited.
func (s *BaseJsonnetListener) ExitIndex(ctx *IndexContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BaseJsonnetListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BaseJsonnetListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterIndexExpr is called when production IndexExpr is entered.
func (s *BaseJsonnetListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production IndexExpr is exited.
func (s *BaseJsonnetListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterArray is called when production Array is entered.
func (s *BaseJsonnetListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production Array is exited.
func (s *BaseJsonnetListener) ExitArray(ctx *ArrayContext) {}

// EnterFunction is called when production Function is entered.
func (s *BaseJsonnetListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production Function is exited.
func (s *BaseJsonnetListener) ExitFunction(ctx *FunctionContext) {}

// EnterErrorExpr is called when production ErrorExpr is entered.
func (s *BaseJsonnetListener) EnterErrorExpr(ctx *ErrorExprContext) {}

// ExitErrorExpr is called when production ErrorExpr is exited.
func (s *BaseJsonnetListener) ExitErrorExpr(ctx *ErrorExprContext) {}

// EnterInSuper is called when production InSuper is entered.
func (s *BaseJsonnetListener) EnterInSuper(ctx *InSuperContext) {}

// ExitInSuper is called when production InSuper is exited.
func (s *BaseJsonnetListener) ExitInSuper(ctx *InSuperContext) {}

// EnterArrayComp is called when production ArrayComp is entered.
func (s *BaseJsonnetListener) EnterArrayComp(ctx *ArrayCompContext) {}

// ExitArrayComp is called when production ArrayComp is exited.
func (s *BaseJsonnetListener) ExitArrayComp(ctx *ArrayCompContext) {}

// EnterAssert is called when production Assert is entered.
func (s *BaseJsonnetListener) EnterAssert(ctx *AssertContext) {}

// ExitAssert is called when production Assert is exited.
func (s *BaseJsonnetListener) ExitAssert(ctx *AssertContext) {}

// EnterIndexSuperExpr is called when production IndexSuperExpr is entered.
func (s *BaseJsonnetListener) EnterIndexSuperExpr(ctx *IndexSuperExprContext) {}

// ExitIndexSuperExpr is called when production IndexSuperExpr is exited.
func (s *BaseJsonnetListener) ExitIndexSuperExpr(ctx *IndexSuperExprContext) {}

// EnterSlice is called when production Slice is entered.
func (s *BaseJsonnetListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production Slice is exited.
func (s *BaseJsonnetListener) ExitSlice(ctx *SliceContext) {}

// EnterValue is called when production Value is entered.
func (s *BaseJsonnetListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production Value is exited.
func (s *BaseJsonnetListener) ExitValue(ctx *ValueContext) {}

// EnterIndexSuper is called when production IndexSuper is entered.
func (s *BaseJsonnetListener) EnterIndexSuper(ctx *IndexSuperContext) {}

// ExitIndexSuper is called when production IndexSuper is exited.
func (s *BaseJsonnetListener) ExitIndexSuper(ctx *IndexSuperContext) {}

// EnterObject is called when production Object is entered.
func (s *BaseJsonnetListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production Object is exited.
func (s *BaseJsonnetListener) ExitObject(ctx *ObjectContext) {}

// EnterIfThenElse is called when production IfThenElse is entered.
func (s *BaseJsonnetListener) EnterIfThenElse(ctx *IfThenElseContext) {}

// ExitIfThenElse is called when production IfThenElse is exited.
func (s *BaseJsonnetListener) ExitIfThenElse(ctx *IfThenElseContext) {}

// EnterLocalBind is called when production LocalBind is entered.
func (s *BaseJsonnetListener) EnterLocalBind(ctx *LocalBindContext) {}

// ExitLocalBind is called when production LocalBind is exited.
func (s *BaseJsonnetListener) ExitLocalBind(ctx *LocalBindContext) {}

// EnterMembers is called when production Members is entered.
func (s *BaseJsonnetListener) EnterMembers(ctx *MembersContext) {}

// ExitMembers is called when production Members is exited.
func (s *BaseJsonnetListener) ExitMembers(ctx *MembersContext) {}

// EnterObjectComp is called when production ObjectComp is entered.
func (s *BaseJsonnetListener) EnterObjectComp(ctx *ObjectCompContext) {}

// ExitObjectComp is called when production ObjectComp is exited.
func (s *BaseJsonnetListener) ExitObjectComp(ctx *ObjectCompContext) {}

// EnterMember is called when production member is entered.
func (s *BaseJsonnetListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseJsonnetListener) ExitMember(ctx *MemberContext) {}

// EnterValueField is called when production ValueField is entered.
func (s *BaseJsonnetListener) EnterValueField(ctx *ValueFieldContext) {}

// ExitValueField is called when production ValueField is exited.
func (s *BaseJsonnetListener) ExitValueField(ctx *ValueFieldContext) {}

// EnterFunctionField is called when production FunctionField is entered.
func (s *BaseJsonnetListener) EnterFunctionField(ctx *FunctionFieldContext) {}

// ExitFunctionField is called when production FunctionField is exited.
func (s *BaseJsonnetListener) ExitFunctionField(ctx *FunctionFieldContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseJsonnetListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseJsonnetListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterObjlocal is called when production objlocal is entered.
func (s *BaseJsonnetListener) EnterObjlocal(ctx *ObjlocalContext) {}

// ExitObjlocal is called when production objlocal is exited.
func (s *BaseJsonnetListener) ExitObjlocal(ctx *ObjlocalContext) {}

// EnterForspec is called when production forspec is entered.
func (s *BaseJsonnetListener) EnterForspec(ctx *ForspecContext) {}

// ExitForspec is called when production forspec is exited.
func (s *BaseJsonnetListener) ExitForspec(ctx *ForspecContext) {}

// EnterIfspec is called when production ifspec is entered.
func (s *BaseJsonnetListener) EnterIfspec(ctx *IfspecContext) {}

// ExitIfspec is called when production ifspec is exited.
func (s *BaseJsonnetListener) ExitIfspec(ctx *IfspecContext) {}

// EnterFieldname is called when production fieldname is entered.
func (s *BaseJsonnetListener) EnterFieldname(ctx *FieldnameContext) {}

// ExitFieldname is called when production fieldname is exited.
func (s *BaseJsonnetListener) ExitFieldname(ctx *FieldnameContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseJsonnetListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseJsonnetListener) ExitAssertion(ctx *AssertionContext) {}

// EnterValueBind is called when production ValueBind is entered.
func (s *BaseJsonnetListener) EnterValueBind(ctx *ValueBindContext) {}

// ExitValueBind is called when production ValueBind is exited.
func (s *BaseJsonnetListener) ExitValueBind(ctx *ValueBindContext) {}

// EnterFunctionBind is called when production FunctionBind is entered.
func (s *BaseJsonnetListener) EnterFunctionBind(ctx *FunctionBindContext) {}

// ExitFunctionBind is called when production FunctionBind is exited.
func (s *BaseJsonnetListener) ExitFunctionBind(ctx *FunctionBindContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseJsonnetListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseJsonnetListener) ExitArgs(ctx *ArgsContext) {}

// EnterParams is called when production params is entered.
func (s *BaseJsonnetListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseJsonnetListener) ExitParams(ctx *ParamsContext) {}
