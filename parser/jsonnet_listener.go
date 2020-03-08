// Code generated from Jsonnet.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Jsonnet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonnetListener is a complete listener for a parse tree produced by JsonnetParser.
type JsonnetListener interface {
	antlr.ParseTreeListener

	// EnterJsonnet is called when entering the jsonnet production.
	EnterJsonnet(c *JsonnetContext)

	// EnterImport is called when entering the Import production.
	EnterImport(c *ImportContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterApply is called when entering the Apply production.
	EnterApply(c *ApplyContext)

	// EnterBinaryExpr is called when entering the BinaryExpr production.
	EnterBinaryExpr(c *BinaryExprContext)

	// EnterApplyBrace is called when entering the ApplyBrace production.
	EnterApplyBrace(c *ApplyBraceContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterArray is called when entering the Array production.
	EnterArray(c *ArrayContext)

	// EnterFunction is called when entering the Function production.
	EnterFunction(c *FunctionContext)

	// EnterErrorExpr is called when entering the ErrorExpr production.
	EnterErrorExpr(c *ErrorExprContext)

	// EnterInSuper is called when entering the InSuper production.
	EnterInSuper(c *InSuperContext)

	// EnterArrayComp is called when entering the ArrayComp production.
	EnterArrayComp(c *ArrayCompContext)

	// EnterAssert is called when entering the Assert production.
	EnterAssert(c *AssertContext)

	// EnterIndexSuperExpr is called when entering the IndexSuperExpr production.
	EnterIndexSuperExpr(c *IndexSuperExprContext)

	// EnterSlice is called when entering the Slice production.
	EnterSlice(c *SliceContext)

	// EnterValue is called when entering the Value production.
	EnterValue(c *ValueContext)

	// EnterIndexSuper is called when entering the IndexSuper production.
	EnterIndexSuper(c *IndexSuperContext)

	// EnterObject is called when entering the Object production.
	EnterObject(c *ObjectContext)

	// EnterIfThenElse is called when entering the IfThenElse production.
	EnterIfThenElse(c *IfThenElseContext)

	// EnterLocalBind is called when entering the LocalBind production.
	EnterLocalBind(c *LocalBindContext)

	// EnterMembers is called when entering the Members production.
	EnterMembers(c *MembersContext)

	// EnterObjectComp is called when entering the ObjectComp production.
	EnterObjectComp(c *ObjectCompContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterValueField is called when entering the ValueField production.
	EnterValueField(c *ValueFieldContext)

	// EnterFunctionField is called when entering the FunctionField production.
	EnterFunctionField(c *FunctionFieldContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterObjlocal is called when entering the objlocal production.
	EnterObjlocal(c *ObjlocalContext)

	// EnterForspec is called when entering the forspec production.
	EnterForspec(c *ForspecContext)

	// EnterIfspec is called when entering the ifspec production.
	EnterIfspec(c *IfspecContext)

	// EnterFieldname is called when entering the fieldname production.
	EnterFieldname(c *FieldnameContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterValueBind is called when entering the ValueBind production.
	EnterValueBind(c *ValueBindContext)

	// EnterFunctionBind is called when entering the FunctionBind production.
	EnterFunctionBind(c *FunctionBindContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// ExitJsonnet is called when exiting the jsonnet production.
	ExitJsonnet(c *JsonnetContext)

	// ExitImport is called when exiting the Import production.
	ExitImport(c *ImportContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitApply is called when exiting the Apply production.
	ExitApply(c *ApplyContext)

	// ExitBinaryExpr is called when exiting the BinaryExpr production.
	ExitBinaryExpr(c *BinaryExprContext)

	// ExitApplyBrace is called when exiting the ApplyBrace production.
	ExitApplyBrace(c *ApplyBraceContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitArray is called when exiting the Array production.
	ExitArray(c *ArrayContext)

	// ExitFunction is called when exiting the Function production.
	ExitFunction(c *FunctionContext)

	// ExitErrorExpr is called when exiting the ErrorExpr production.
	ExitErrorExpr(c *ErrorExprContext)

	// ExitInSuper is called when exiting the InSuper production.
	ExitInSuper(c *InSuperContext)

	// ExitArrayComp is called when exiting the ArrayComp production.
	ExitArrayComp(c *ArrayCompContext)

	// ExitAssert is called when exiting the Assert production.
	ExitAssert(c *AssertContext)

	// ExitIndexSuperExpr is called when exiting the IndexSuperExpr production.
	ExitIndexSuperExpr(c *IndexSuperExprContext)

	// ExitSlice is called when exiting the Slice production.
	ExitSlice(c *SliceContext)

	// ExitValue is called when exiting the Value production.
	ExitValue(c *ValueContext)

	// ExitIndexSuper is called when exiting the IndexSuper production.
	ExitIndexSuper(c *IndexSuperContext)

	// ExitObject is called when exiting the Object production.
	ExitObject(c *ObjectContext)

	// ExitIfThenElse is called when exiting the IfThenElse production.
	ExitIfThenElse(c *IfThenElseContext)

	// ExitLocalBind is called when exiting the LocalBind production.
	ExitLocalBind(c *LocalBindContext)

	// ExitMembers is called when exiting the Members production.
	ExitMembers(c *MembersContext)

	// ExitObjectComp is called when exiting the ObjectComp production.
	ExitObjectComp(c *ObjectCompContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitValueField is called when exiting the ValueField production.
	ExitValueField(c *ValueFieldContext)

	// ExitFunctionField is called when exiting the FunctionField production.
	ExitFunctionField(c *FunctionFieldContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitObjlocal is called when exiting the objlocal production.
	ExitObjlocal(c *ObjlocalContext)

	// ExitForspec is called when exiting the forspec production.
	ExitForspec(c *ForspecContext)

	// ExitIfspec is called when exiting the ifspec production.
	ExitIfspec(c *IfspecContext)

	// ExitFieldname is called when exiting the fieldname production.
	ExitFieldname(c *FieldnameContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitValueBind is called when exiting the ValueBind production.
	ExitValueBind(c *ValueBindContext)

	// ExitFunctionBind is called when exiting the FunctionBind production.
	ExitFunctionBind(c *FunctionBindContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)
}
