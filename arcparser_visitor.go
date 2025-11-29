// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ArcParser.
type ArcParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ArcParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ArcParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ArcParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externDecl.
	VisitExternDecl(ctx *ExternDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externFunctionDecl.
	VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#externParameterList.
	VisitExternParameterList(ctx *ExternParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by ArcParser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#structFieldList.
	VisitStructFieldList(ctx *StructFieldListContext) interface{}

	// Visit a parse tree produced by ArcParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ArcParser#VarDeclStmt.
	VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#StoreStmt.
	VisitStoreStmt(ctx *StoreStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#ExprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ArcParser#PrimitiveTypeSpec.
	VisitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) interface{}

	// Visit a parse tree produced by ArcParser#PointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#ReferenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#VectorType.
	VisitVectorType(ctx *VectorTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#MapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#StructType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#DerefExpr.
	VisitDerefExpr(ctx *DerefExprContext) interface{}

	// Visit a parse tree produced by ArcParser#ComparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by ArcParser#LogicalNotExpr.
	VisitLogicalNotExpr(ctx *LogicalNotExprContext) interface{}

	// Visit a parse tree produced by ArcParser#LogicalAndExpr.
	VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{}

	// Visit a parse tree produced by ArcParser#LogicalOrExpr.
	VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{}

	// Visit a parse tree produced by ArcParser#EqualityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by ArcParser#MulDivModExpr.
	VisitMulDivModExpr(ctx *MulDivModExprContext) interface{}

	// Visit a parse tree produced by ArcParser#AddrOfExpr.
	VisitAddrOfExpr(ctx *AddrOfExprContext) interface{}

	// Visit a parse tree produced by ArcParser#CastExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by ArcParser#PrimaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by ArcParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by ArcParser#FieldAccessExpr.
	VisitFieldAccessExpr(ctx *FieldAccessExprContext) interface{}

	// Visit a parse tree produced by ArcParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by ArcParser#AllocaExpr.
	VisitAllocaExpr(ctx *AllocaExprContext) interface{}

	// Visit a parse tree produced by ArcParser#UnaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by ArcParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#CharLiteral.
	VisitCharLiteral(ctx *CharLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#BoolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#IdentifierExpr.
	VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

	// Visit a parse tree produced by ArcParser#StructLiteralExpr.
	VisitStructLiteralExpr(ctx *StructLiteralExprContext) interface{}

	// Visit a parse tree produced by ArcParser#VectorLiteralExpr.
	VisitVectorLiteralExpr(ctx *VectorLiteralExprContext) interface{}

	// Visit a parse tree produced by ArcParser#MapLiteralExpr.
	VisitMapLiteralExpr(ctx *MapLiteralExprContext) interface{}

	// Visit a parse tree produced by ArcParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ArcParser#structLiteral.
	VisitStructLiteral(ctx *StructLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#fieldInitList.
	VisitFieldInitList(ctx *FieldInitListContext) interface{}

	// Visit a parse tree produced by ArcParser#fieldInit.
	VisitFieldInit(ctx *FieldInitContext) interface{}

	// Visit a parse tree produced by ArcParser#vectorLiteral.
	VisitVectorLiteral(ctx *VectorLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#mapEntry.
	VisitMapEntry(ctx *MapEntryContext) interface{}

	// Visit a parse tree produced by ArcParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}
}
