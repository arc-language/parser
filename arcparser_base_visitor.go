// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

type BaseArcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArcParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternDecl(ctx *ExternDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternParameterList(ctx *ExternParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructFieldList(ctx *StructFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStoreStmt(ctx *StoreStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVectorType(ctx *VectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitDerefExpr(ctx *DerefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLogicalNotExpr(ctx *LogicalNotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMulDivModExpr(ctx *MulDivModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAddrOfExpr(ctx *AddrOfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCastExpr(ctx *CastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFieldAccessExpr(ctx *FieldAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAllocaExpr(ctx *AllocaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCharLiteral(ctx *CharLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructLiteralExpr(ctx *StructLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVectorLiteralExpr(ctx *VectorLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapLiteralExpr(ctx *MapLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFieldInitList(ctx *FieldInitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVectorLiteral(ctx *VectorLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapEntry(ctx *MapEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}
