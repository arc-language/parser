// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// BaseArcParserListener is a complete listener for a parse tree produced by ArcParser.
type BaseArcParserListener struct{}

var _ ArcParserListener = &BaseArcParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArcParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArcParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseArcParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseArcParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseArcParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseArcParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseArcParserListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseArcParserListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterExternDecl is called when production externDecl is entered.
func (s *BaseArcParserListener) EnterExternDecl(ctx *ExternDeclContext) {}

// ExitExternDecl is called when production externDecl is exited.
func (s *BaseArcParserListener) ExitExternDecl(ctx *ExternDeclContext) {}

// EnterExternFunctionDecl is called when production externFunctionDecl is entered.
func (s *BaseArcParserListener) EnterExternFunctionDecl(ctx *ExternFunctionDeclContext) {}

// ExitExternFunctionDecl is called when production externFunctionDecl is exited.
func (s *BaseArcParserListener) ExitExternFunctionDecl(ctx *ExternFunctionDeclContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseArcParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseArcParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterExternParameterList is called when production externParameterList is entered.
func (s *BaseArcParserListener) EnterExternParameterList(ctx *ExternParameterListContext) {}

// ExitExternParameterList is called when production externParameterList is exited.
func (s *BaseArcParserListener) ExitExternParameterList(ctx *ExternParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseArcParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseArcParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseArcParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseArcParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseArcParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseArcParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructFieldList is called when production structFieldList is entered.
func (s *BaseArcParserListener) EnterStructFieldList(ctx *StructFieldListContext) {}

// ExitStructFieldList is called when production structFieldList is exited.
func (s *BaseArcParserListener) ExitStructFieldList(ctx *StructFieldListContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseArcParserListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseArcParserListener) ExitStructField(ctx *StructFieldContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseArcParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseArcParserListener) ExitBlock(ctx *BlockContext) {}

// EnterVarDeclStmt is called when production VarDeclStmt is entered.
func (s *BaseArcParserListener) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production VarDeclStmt is exited.
func (s *BaseArcParserListener) ExitVarDeclStmt(ctx *VarDeclStmtContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseArcParserListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseArcParserListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterStoreStmt is called when production StoreStmt is entered.
func (s *BaseArcParserListener) EnterStoreStmt(ctx *StoreStmtContext) {}

// ExitStoreStmt is called when production StoreStmt is exited.
func (s *BaseArcParserListener) ExitStoreStmt(ctx *StoreStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseArcParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseArcParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterExprStmt is called when production ExprStmt is entered.
func (s *BaseArcParserListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production ExprStmt is exited.
func (s *BaseArcParserListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseArcParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseArcParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterPrimitiveTypeSpec is called when production PrimitiveTypeSpec is entered.
func (s *BaseArcParserListener) EnterPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) {}

// ExitPrimitiveTypeSpec is called when production PrimitiveTypeSpec is exited.
func (s *BaseArcParserListener) ExitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) {}

// EnterPointerType is called when production PointerType is entered.
func (s *BaseArcParserListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production PointerType is exited.
func (s *BaseArcParserListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterReferenceType is called when production ReferenceType is entered.
func (s *BaseArcParserListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production ReferenceType is exited.
func (s *BaseArcParserListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterVectorType is called when production VectorType is entered.
func (s *BaseArcParserListener) EnterVectorType(ctx *VectorTypeContext) {}

// ExitVectorType is called when production VectorType is exited.
func (s *BaseArcParserListener) ExitVectorType(ctx *VectorTypeContext) {}

// EnterMapType is called when production MapType is entered.
func (s *BaseArcParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production MapType is exited.
func (s *BaseArcParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterStructType is called when production StructType is entered.
func (s *BaseArcParserListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production StructType is exited.
func (s *BaseArcParserListener) ExitStructType(ctx *StructTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseArcParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseArcParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterDerefExpr is called when production DerefExpr is entered.
func (s *BaseArcParserListener) EnterDerefExpr(ctx *DerefExprContext) {}

// ExitDerefExpr is called when production DerefExpr is exited.
func (s *BaseArcParserListener) ExitDerefExpr(ctx *DerefExprContext) {}

// EnterComparisonExpr is called when production ComparisonExpr is entered.
func (s *BaseArcParserListener) EnterComparisonExpr(ctx *ComparisonExprContext) {}

// ExitComparisonExpr is called when production ComparisonExpr is exited.
func (s *BaseArcParserListener) ExitComparisonExpr(ctx *ComparisonExprContext) {}

// EnterLogicalNotExpr is called when production LogicalNotExpr is entered.
func (s *BaseArcParserListener) EnterLogicalNotExpr(ctx *LogicalNotExprContext) {}

// ExitLogicalNotExpr is called when production LogicalNotExpr is exited.
func (s *BaseArcParserListener) ExitLogicalNotExpr(ctx *LogicalNotExprContext) {}

// EnterLogicalAndExpr is called when production LogicalAndExpr is entered.
func (s *BaseArcParserListener) EnterLogicalAndExpr(ctx *LogicalAndExprContext) {}

// ExitLogicalAndExpr is called when production LogicalAndExpr is exited.
func (s *BaseArcParserListener) ExitLogicalAndExpr(ctx *LogicalAndExprContext) {}

// EnterLogicalOrExpr is called when production LogicalOrExpr is entered.
func (s *BaseArcParserListener) EnterLogicalOrExpr(ctx *LogicalOrExprContext) {}

// ExitLogicalOrExpr is called when production LogicalOrExpr is exited.
func (s *BaseArcParserListener) ExitLogicalOrExpr(ctx *LogicalOrExprContext) {}

// EnterEqualityExpr is called when production EqualityExpr is entered.
func (s *BaseArcParserListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production EqualityExpr is exited.
func (s *BaseArcParserListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterMulDivModExpr is called when production MulDivModExpr is entered.
func (s *BaseArcParserListener) EnterMulDivModExpr(ctx *MulDivModExprContext) {}

// ExitMulDivModExpr is called when production MulDivModExpr is exited.
func (s *BaseArcParserListener) ExitMulDivModExpr(ctx *MulDivModExprContext) {}

// EnterAddrOfExpr is called when production AddrOfExpr is entered.
func (s *BaseArcParserListener) EnterAddrOfExpr(ctx *AddrOfExprContext) {}

// ExitAddrOfExpr is called when production AddrOfExpr is exited.
func (s *BaseArcParserListener) ExitAddrOfExpr(ctx *AddrOfExprContext) {}

// EnterCastExpr is called when production CastExpr is entered.
func (s *BaseArcParserListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production CastExpr is exited.
func (s *BaseArcParserListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterPrimaryExpr is called when production PrimaryExpr is entered.
func (s *BaseArcParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production PrimaryExpr is exited.
func (s *BaseArcParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseArcParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseArcParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterFieldAccessExpr is called when production FieldAccessExpr is entered.
func (s *BaseArcParserListener) EnterFieldAccessExpr(ctx *FieldAccessExprContext) {}

// ExitFieldAccessExpr is called when production FieldAccessExpr is exited.
func (s *BaseArcParserListener) ExitFieldAccessExpr(ctx *FieldAccessExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseArcParserListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseArcParserListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterAllocaExpr is called when production AllocaExpr is entered.
func (s *BaseArcParserListener) EnterAllocaExpr(ctx *AllocaExprContext) {}

// ExitAllocaExpr is called when production AllocaExpr is exited.
func (s *BaseArcParserListener) ExitAllocaExpr(ctx *AllocaExprContext) {}

// EnterUnaryMinusExpr is called when production UnaryMinusExpr is entered.
func (s *BaseArcParserListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production UnaryMinusExpr is exited.
func (s *BaseArcParserListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseArcParserListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseArcParserListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseArcParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseArcParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseArcParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseArcParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterCharLiteral is called when production CharLiteral is entered.
func (s *BaseArcParserListener) EnterCharLiteral(ctx *CharLiteralContext) {}

// ExitCharLiteral is called when production CharLiteral is exited.
func (s *BaseArcParserListener) ExitCharLiteral(ctx *CharLiteralContext) {}

// EnterBoolLiteral is called when production BoolLiteral is entered.
func (s *BaseArcParserListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production BoolLiteral is exited.
func (s *BaseArcParserListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseArcParserListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseArcParserListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterStructLiteralExpr is called when production StructLiteralExpr is entered.
func (s *BaseArcParserListener) EnterStructLiteralExpr(ctx *StructLiteralExprContext) {}

// ExitStructLiteralExpr is called when production StructLiteralExpr is exited.
func (s *BaseArcParserListener) ExitStructLiteralExpr(ctx *StructLiteralExprContext) {}

// EnterVectorLiteralExpr is called when production VectorLiteralExpr is entered.
func (s *BaseArcParserListener) EnterVectorLiteralExpr(ctx *VectorLiteralExprContext) {}

// ExitVectorLiteralExpr is called when production VectorLiteralExpr is exited.
func (s *BaseArcParserListener) ExitVectorLiteralExpr(ctx *VectorLiteralExprContext) {}

// EnterMapLiteralExpr is called when production MapLiteralExpr is entered.
func (s *BaseArcParserListener) EnterMapLiteralExpr(ctx *MapLiteralExprContext) {}

// ExitMapLiteralExpr is called when production MapLiteralExpr is exited.
func (s *BaseArcParserListener) ExitMapLiteralExpr(ctx *MapLiteralExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseArcParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseArcParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BaseArcParserListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BaseArcParserListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterFieldInitList is called when production fieldInitList is entered.
func (s *BaseArcParserListener) EnterFieldInitList(ctx *FieldInitListContext) {}

// ExitFieldInitList is called when production fieldInitList is exited.
func (s *BaseArcParserListener) ExitFieldInitList(ctx *FieldInitListContext) {}

// EnterFieldInit is called when production fieldInit is entered.
func (s *BaseArcParserListener) EnterFieldInit(ctx *FieldInitContext) {}

// ExitFieldInit is called when production fieldInit is exited.
func (s *BaseArcParserListener) ExitFieldInit(ctx *FieldInitContext) {}

// EnterVectorLiteral is called when production vectorLiteral is entered.
func (s *BaseArcParserListener) EnterVectorLiteral(ctx *VectorLiteralContext) {}

// ExitVectorLiteral is called when production vectorLiteral is exited.
func (s *BaseArcParserListener) ExitVectorLiteral(ctx *VectorLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseArcParserListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseArcParserListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseArcParserListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseArcParserListener) ExitMapEntry(ctx *MapEntryContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseArcParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseArcParserListener) ExitArgumentList(ctx *ArgumentListContext) {}
