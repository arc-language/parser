// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// ArcParserListener is a complete listener for a parse tree produced by ArcParser.
type ArcParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructFieldList is called when entering the structFieldList production.
	EnterStructFieldList(c *StructFieldListContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterVarDeclStmt is called when entering the VarDeclStmt production.
	EnterVarDeclStmt(c *VarDeclStmtContext)

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterStoreStmt is called when entering the StoreStmt production.
	EnterStoreStmt(c *StoreStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterExprStmt is called when entering the ExprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterPrimitiveTypeSpec is called when entering the PrimitiveTypeSpec production.
	EnterPrimitiveTypeSpec(c *PrimitiveTypeSpecContext)

	// EnterPointerType is called when entering the PointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterReferenceType is called when entering the ReferenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterVectorType is called when entering the VectorType production.
	EnterVectorType(c *VectorTypeContext)

	// EnterMapType is called when entering the MapType production.
	EnterMapType(c *MapTypeContext)

	// EnterStructType is called when entering the StructType production.
	EnterStructType(c *StructTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterDerefExpr is called when entering the DerefExpr production.
	EnterDerefExpr(c *DerefExprContext)

	// EnterComparisonExpr is called when entering the ComparisonExpr production.
	EnterComparisonExpr(c *ComparisonExprContext)

	// EnterLogicalNotExpr is called when entering the LogicalNotExpr production.
	EnterLogicalNotExpr(c *LogicalNotExprContext)

	// EnterLogicalAndExpr is called when entering the LogicalAndExpr production.
	EnterLogicalAndExpr(c *LogicalAndExprContext)

	// EnterLogicalOrExpr is called when entering the LogicalOrExpr production.
	EnterLogicalOrExpr(c *LogicalOrExprContext)

	// EnterEqualityExpr is called when entering the EqualityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterMulDivModExpr is called when entering the MulDivModExpr production.
	EnterMulDivModExpr(c *MulDivModExprContext)

	// EnterAddrOfExpr is called when entering the AddrOfExpr production.
	EnterAddrOfExpr(c *AddrOfExprContext)

	// EnterCastExpr is called when entering the CastExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterFieldAccessExpr is called when entering the FieldAccessExpr production.
	EnterFieldAccessExpr(c *FieldAccessExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterAllocaExpr is called when entering the AllocaExpr production.
	EnterAllocaExpr(c *AllocaExprContext)

	// EnterUnaryMinusExpr is called when entering the UnaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterCharLiteral is called when entering the CharLiteral production.
	EnterCharLiteral(c *CharLiteralContext)

	// EnterBoolLiteral is called when entering the BoolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterStructLiteralExpr is called when entering the StructLiteralExpr production.
	EnterStructLiteralExpr(c *StructLiteralExprContext)

	// EnterVectorLiteralExpr is called when entering the VectorLiteralExpr production.
	EnterVectorLiteralExpr(c *VectorLiteralExprContext)

	// EnterMapLiteralExpr is called when entering the MapLiteralExpr production.
	EnterMapLiteralExpr(c *MapLiteralExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterFieldInitList is called when entering the fieldInitList production.
	EnterFieldInitList(c *FieldInitListContext)

	// EnterFieldInit is called when entering the fieldInit production.
	EnterFieldInit(c *FieldInitContext)

	// EnterVectorLiteral is called when entering the vectorLiteral production.
	EnterVectorLiteral(c *VectorLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructFieldList is called when exiting the structFieldList production.
	ExitStructFieldList(c *StructFieldListContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitVarDeclStmt is called when exiting the VarDeclStmt production.
	ExitVarDeclStmt(c *VarDeclStmtContext)

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitStoreStmt is called when exiting the StoreStmt production.
	ExitStoreStmt(c *StoreStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitExprStmt is called when exiting the ExprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitPrimitiveTypeSpec is called when exiting the PrimitiveTypeSpec production.
	ExitPrimitiveTypeSpec(c *PrimitiveTypeSpecContext)

	// ExitPointerType is called when exiting the PointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitReferenceType is called when exiting the ReferenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitVectorType is called when exiting the VectorType production.
	ExitVectorType(c *VectorTypeContext)

	// ExitMapType is called when exiting the MapType production.
	ExitMapType(c *MapTypeContext)

	// ExitStructType is called when exiting the StructType production.
	ExitStructType(c *StructTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitDerefExpr is called when exiting the DerefExpr production.
	ExitDerefExpr(c *DerefExprContext)

	// ExitComparisonExpr is called when exiting the ComparisonExpr production.
	ExitComparisonExpr(c *ComparisonExprContext)

	// ExitLogicalNotExpr is called when exiting the LogicalNotExpr production.
	ExitLogicalNotExpr(c *LogicalNotExprContext)

	// ExitLogicalAndExpr is called when exiting the LogicalAndExpr production.
	ExitLogicalAndExpr(c *LogicalAndExprContext)

	// ExitLogicalOrExpr is called when exiting the LogicalOrExpr production.
	ExitLogicalOrExpr(c *LogicalOrExprContext)

	// ExitEqualityExpr is called when exiting the EqualityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitMulDivModExpr is called when exiting the MulDivModExpr production.
	ExitMulDivModExpr(c *MulDivModExprContext)

	// ExitAddrOfExpr is called when exiting the AddrOfExpr production.
	ExitAddrOfExpr(c *AddrOfExprContext)

	// ExitCastExpr is called when exiting the CastExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitFieldAccessExpr is called when exiting the FieldAccessExpr production.
	ExitFieldAccessExpr(c *FieldAccessExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitAllocaExpr is called when exiting the AllocaExpr production.
	ExitAllocaExpr(c *AllocaExprContext)

	// ExitUnaryMinusExpr is called when exiting the UnaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitCharLiteral is called when exiting the CharLiteral production.
	ExitCharLiteral(c *CharLiteralContext)

	// ExitBoolLiteral is called when exiting the BoolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitStructLiteralExpr is called when exiting the StructLiteralExpr production.
	ExitStructLiteralExpr(c *StructLiteralExprContext)

	// ExitVectorLiteralExpr is called when exiting the VectorLiteralExpr production.
	ExitVectorLiteralExpr(c *VectorLiteralExprContext)

	// ExitMapLiteralExpr is called when exiting the MapLiteralExpr production.
	ExitMapLiteralExpr(c *MapLiteralExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitFieldInitList is called when exiting the fieldInitList production.
	ExitFieldInitList(c *FieldInitListContext)

	// ExitFieldInit is called when exiting the fieldInit production.
	ExitFieldInit(c *FieldInitContext)

	// ExitVectorLiteral is called when exiting the vectorLiteral production.
	ExitVectorLiteral(c *VectorLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)
}
