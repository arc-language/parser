

// THIS the parser package 
package "github.com/arc-language/parser"


CONSTANTS

const (
	ArcLexerLET             = 1
	ArcLexerCONST           = 2
	ArcLexerFUNC            = 3
	ArcLexerEXTERN          = 4
	ArcLexerSTRUCT          = 5
	ArcLexerRETURN          = 6
	ArcLexerIF              = 7
	ArcLexerELSE            = 8
	ArcLexerALLOCA          = 9
	ArcLexerCAST            = 10
	ArcLexerINT32           = 11
	ArcLexerINT64           = 12
	ArcLexerFLOAT32         = 13
	ArcLexerFLOAT64         = 14
	ArcLexerBOOL            = 15
	ArcLexerSTRING          = 16
	ArcLexerBYTE            = 17
	ArcLexerCHAR            = 18
	ArcLexerVOID            = 19
	ArcLexerVECTOR          = 20
	ArcLexerMAP             = 21
	ArcLexerTRUE            = 22
	ArcLexerFALSE           = 23
	ArcLexerPLUS            = 24
	ArcLexerMINUS           = 25
	ArcLexerSTAR            = 26
	ArcLexerSLASH           = 27
	ArcLexerPERCENT         = 28
	ArcLexerEQ              = 29
	ArcLexerNE              = 30
	ArcLexerLT              = 31
	ArcLexerLE              = 32
	ArcLexerGT              = 33
	ArcLexerGE              = 34
	ArcLexerAND             = 35
	ArcLexerOR              = 36
	ArcLexerNOT             = 37
	ArcLexerAMP             = 38
	ArcLexerAT              = 39
	ArcLexerASSIGN          = 40
	ArcLexerARROW           = 41
	ArcLexerELLIPSIS        = 42
	ArcLexerLPAREN          = 43
	ArcLexerRPAREN          = 44
	ArcLexerLBRACE          = 45
	ArcLexerRBRACE          = 46
	ArcLexerLBRACK          = 47
	ArcLexerRBRACK          = 48
	ArcLexerCOMMA           = 49
	ArcLexerCOLON           = 50
	ArcLexerSEMICOLON       = 51
	ArcLexerDOT             = 52
	ArcLexerIDENTIFIER      = 53
	ArcLexerINTEGER_LITERAL = 54
	ArcLexerFLOAT_LITERAL   = 55
	ArcLexerSTRING_LITERAL  = 56
	ArcLexerCHAR_LITERAL    = 57
	ArcLexerWS              = 58
	ArcLexerLINE_COMMENT    = 59
	ArcLexerBLOCK_COMMENT   = 60
)
    ArcLexer tokens.

const (
	ArcParserEOF             = antlr.TokenEOF
	ArcParserLET             = 1
	ArcParserCONST           = 2
	ArcParserFUNC            = 3
	ArcParserEXTERN          = 4
	ArcParserSTRUCT          = 5
	ArcParserRETURN          = 6
	ArcParserIF              = 7
	ArcParserELSE            = 8
	ArcParserALLOCA          = 9
	ArcParserCAST            = 10
	ArcParserINT32           = 11
	ArcParserINT64           = 12
	ArcParserFLOAT32         = 13
	ArcParserFLOAT64         = 14
	ArcParserBOOL            = 15
	ArcParserSTRING          = 16
	ArcParserBYTE            = 17
	ArcParserCHAR            = 18
	ArcParserVOID            = 19
	ArcParserVECTOR          = 20
	ArcParserMAP             = 21
	ArcParserTRUE            = 22
	ArcParserFALSE           = 23
	ArcParserPLUS            = 24
	ArcParserMINUS           = 25
	ArcParserSTAR            = 26
	ArcParserSLASH           = 27
	ArcParserPERCENT         = 28
	ArcParserEQ              = 29
	ArcParserNE              = 30
	ArcParserLT              = 31
	ArcParserLE              = 32
	ArcParserGT              = 33
	ArcParserGE              = 34
	ArcParserAND             = 35
	ArcParserOR              = 36
	ArcParserNOT             = 37
	ArcParserAMP             = 38
	ArcParserAT              = 39
	ArcParserASSIGN          = 40
	ArcParserARROW           = 41
	ArcParserELLIPSIS        = 42
	ArcParserLPAREN          = 43
	ArcParserRPAREN          = 44
	ArcParserLBRACE          = 45
	ArcParserRBRACE          = 46
	ArcParserLBRACK          = 47
	ArcParserRBRACK          = 48
	ArcParserCOMMA           = 49
	ArcParserCOLON           = 50
	ArcParserSEMICOLON       = 51
	ArcParserDOT             = 52
	ArcParserIDENTIFIER      = 53
	ArcParserINTEGER_LITERAL = 54
	ArcParserFLOAT_LITERAL   = 55
	ArcParserSTRING_LITERAL  = 56
	ArcParserCHAR_LITERAL    = 57
	ArcParserWS              = 58
	ArcParserLINE_COMMENT    = 59
	ArcParserBLOCK_COMMENT   = 60
)
    ArcParser tokens.

const (
	ArcParserRULE_program             = 0
	ArcParserRULE_declaration         = 1
	ArcParserRULE_variableDecl        = 2
	ArcParserRULE_functionDecl        = 3
	ArcParserRULE_externDecl          = 4
	ArcParserRULE_externFunctionDecl  = 5
	ArcParserRULE_parameterList       = 6
	ArcParserRULE_externParameterList = 7
	ArcParserRULE_parameter           = 8
	ArcParserRULE_returnType          = 9
	ArcParserRULE_structDecl          = 10
	ArcParserRULE_structFieldList     = 11
	ArcParserRULE_structField         = 12
	ArcParserRULE_block               = 13
	ArcParserRULE_statement           = 14
	ArcParserRULE_ifStatement         = 15
	ArcParserRULE_type                = 16
	ArcParserRULE_primitiveType       = 17
	ArcParserRULE_expression          = 18
	ArcParserRULE_primary             = 19
	ArcParserRULE_structLiteral       = 20
	ArcParserRULE_fieldInitList       = 21
	ArcParserRULE_fieldInit           = 22
	ArcParserRULE_vectorLiteral       = 23
	ArcParserRULE_mapLiteral          = 24
	ArcParserRULE_mapEntry            = 25
	ArcParserRULE_argumentList        = 26
)
    ArcParser rules.


VARIABLES

var ArcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}
var ArcParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

FUNCTIONS

func ArcLexerInit()
    ArcLexerInit initializes any static state used to implement ArcLexer. By
    default the static state used to implement the lexer is lazily initialized
    during the first call to NewArcLexer(). You can call this function if you
    wish to initialize the static state ahead of time.

func ArcParserInit()
    ArcParserInit initializes any static state used to implement ArcParser. By
    default the static state used to implement the parser is lazily initialized
    during the first call to NewArcParser(). You can call this function if you
    wish to initialize the static state ahead of time.

func InitEmptyArgumentListContext(p *ArgumentListContext)
func InitEmptyBlockContext(p *BlockContext)
func InitEmptyDeclarationContext(p *DeclarationContext)
func InitEmptyExpressionContext(p *ExpressionContext)
func InitEmptyExternDeclContext(p *ExternDeclContext)
func InitEmptyExternFunctionDeclContext(p *ExternFunctionDeclContext)
func InitEmptyExternParameterListContext(p *ExternParameterListContext)
func InitEmptyFieldInitContext(p *FieldInitContext)
func InitEmptyFieldInitListContext(p *FieldInitListContext)
func InitEmptyFunctionDeclContext(p *FunctionDeclContext)
func InitEmptyIfStatementContext(p *IfStatementContext)
func InitEmptyMapEntryContext(p *MapEntryContext)
func InitEmptyMapLiteralContext(p *MapLiteralContext)
func InitEmptyParameterContext(p *ParameterContext)
func InitEmptyParameterListContext(p *ParameterListContext)
func InitEmptyPrimaryContext(p *PrimaryContext)
func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext)
func InitEmptyProgramContext(p *ProgramContext)
func InitEmptyReturnTypeContext(p *ReturnTypeContext)
func InitEmptyStatementContext(p *StatementContext)
func InitEmptyStructDeclContext(p *StructDeclContext)
func InitEmptyStructFieldContext(p *StructFieldContext)
func InitEmptyStructFieldListContext(p *StructFieldListContext)
func InitEmptyStructLiteralContext(p *StructLiteralContext)
func InitEmptyTypeContext(p *TypeContext)
func InitEmptyVariableDeclContext(p *VariableDeclContext)
func InitEmptyVectorLiteralContext(p *VectorLiteralContext)

TYPES

type AddSubExprContext struct {
	ExpressionContext
	// Has unexported fields.
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AddSubExprContext) AllExpression() []IExpressionContext

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AddSubExprContext) Expression(i int) IExpressionContext

func (s *AddSubExprContext) GetOp() antlr.Token

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext

func (s *AddSubExprContext) MINUS() antlr.TerminalNode

func (s *AddSubExprContext) PLUS() antlr.TerminalNode

func (s *AddSubExprContext) SetOp(v antlr.Token)

type AddrOfExprContext struct {
	ExpressionContext
}

func NewAddrOfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddrOfExprContext

func (s *AddrOfExprContext) AMP() antlr.TerminalNode

func (s *AddrOfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AddrOfExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AddrOfExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AddrOfExprContext) Expression() IExpressionContext

func (s *AddrOfExprContext) GetRuleContext() antlr.RuleContext

type AllocaExprContext struct {
	ExpressionContext
}

func NewAllocaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocaExprContext

func (s *AllocaExprContext) ALLOCA() antlr.TerminalNode

func (s *AllocaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AllocaExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AllocaExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AllocaExprContext) GetRuleContext() antlr.RuleContext

func (s *AllocaExprContext) LPAREN() antlr.TerminalNode

func (s *AllocaExprContext) RPAREN() antlr.TerminalNode

func (s *AllocaExprContext) Type_() ITypeContext

type ArcLexer struct {
	*antlr.BaseLexer

	// Has unexported fields.
}

func NewArcLexer(input antlr.CharStream) *ArcLexer
    NewArcLexer produces a new lexer instance for the optional input
    antlr.CharStream.

type ArcParser struct {
	*antlr.BaseParser
}

func NewArcParser(input antlr.TokenStream) *ArcParser
    NewArcParser produces a new parser instance for the optional input
    antlr.TokenStream.

func (p *ArcParser) ArgumentList() (localctx IArgumentListContext)

func (p *ArcParser) Block() (localctx IBlockContext)

func (p *ArcParser) Declaration() (localctx IDeclarationContext)

func (p *ArcParser) Expression() (localctx IExpressionContext)

func (p *ArcParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool

func (p *ArcParser) ExternDecl() (localctx IExternDeclContext)

func (p *ArcParser) ExternFunctionDecl() (localctx IExternFunctionDeclContext)

func (p *ArcParser) ExternParameterList() (localctx IExternParameterListContext)

func (p *ArcParser) FieldInit() (localctx IFieldInitContext)

func (p *ArcParser) FieldInitList() (localctx IFieldInitListContext)

func (p *ArcParser) FunctionDecl() (localctx IFunctionDeclContext)

func (p *ArcParser) IfStatement() (localctx IIfStatementContext)

func (p *ArcParser) MapEntry() (localctx IMapEntryContext)

func (p *ArcParser) MapLiteral() (localctx IMapLiteralContext)

func (p *ArcParser) Parameter() (localctx IParameterContext)

func (p *ArcParser) ParameterList() (localctx IParameterListContext)

func (p *ArcParser) Primary() (localctx IPrimaryContext)

func (p *ArcParser) PrimitiveType() (localctx IPrimitiveTypeContext)

func (p *ArcParser) Program() (localctx IProgramContext)

func (p *ArcParser) ReturnType() (localctx IReturnTypeContext)

func (p *ArcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool

func (p *ArcParser) Statement() (localctx IStatementContext)

func (p *ArcParser) StructDecl() (localctx IStructDeclContext)

func (p *ArcParser) StructField() (localctx IStructFieldContext)

func (p *ArcParser) StructFieldList() (localctx IStructFieldListContext)

func (p *ArcParser) StructLiteral() (localctx IStructLiteralContext)

func (p *ArcParser) Type_() (localctx ITypeContext)

func (p *ArcParser) VariableDecl() (localctx IVariableDeclContext)

func (p *ArcParser) VectorLiteral() (localctx IVectorLiteralContext)

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

	// EnterExternDecl is called when entering the externDecl production.
	EnterExternDecl(c *ExternDeclContext)

	// EnterExternFunctionDecl is called when entering the externFunctionDecl production.
	EnterExternFunctionDecl(c *ExternFunctionDeclContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterExternParameterList is called when entering the externParameterList production.
	EnterExternParameterList(c *ExternParameterListContext)

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

	// ExitExternDecl is called when exiting the externDecl production.
	ExitExternDecl(c *ExternDeclContext)

	// ExitExternFunctionDecl is called when exiting the externFunctionDecl production.
	ExitExternFunctionDecl(c *ExternFunctionDeclContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitExternParameterList is called when exiting the externParameterList production.
	ExitExternParameterList(c *ExternParameterListContext)

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
    ArcParserListener is a complete listener for a parse tree produced by
    ArcParser.

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
    A complete Visitor for a parse tree produced by ArcParser.

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext

func NewEmptyArgumentListContext() *ArgumentListContext

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode

func (s *ArgumentListContext) AllExpression() []IExpressionContext

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ArgumentListContext) Expression(i int) IExpressionContext

func (s *ArgumentListContext) GetParser() antlr.Parser

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext

func (*ArgumentListContext) IsArgumentListContext()

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type AssignStmtContext struct {
	StatementContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext

func (s *AssignStmtContext) ASSIGN() antlr.TerminalNode

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AssignStmtContext) Expression() IExpressionContext

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext

func (s *AssignStmtContext) IDENTIFIER() antlr.TerminalNode

type BaseArcParserListener struct{}
    BaseArcParserListener is a complete listener for a parse tree produced by
    ArcParser.

func (s *BaseArcParserListener) EnterAddSubExpr(ctx *AddSubExprContext)
    EnterAddSubExpr is called when production AddSubExpr is entered.

func (s *BaseArcParserListener) EnterAddrOfExpr(ctx *AddrOfExprContext)
    EnterAddrOfExpr is called when production AddrOfExpr is entered.

func (s *BaseArcParserListener) EnterAllocaExpr(ctx *AllocaExprContext)
    EnterAllocaExpr is called when production AllocaExpr is entered.

func (s *BaseArcParserListener) EnterArgumentList(ctx *ArgumentListContext)
    EnterArgumentList is called when production argumentList is entered.

func (s *BaseArcParserListener) EnterAssignStmt(ctx *AssignStmtContext)
    EnterAssignStmt is called when production AssignStmt is entered.

func (s *BaseArcParserListener) EnterBlock(ctx *BlockContext)
    EnterBlock is called when production block is entered.

func (s *BaseArcParserListener) EnterBoolLiteral(ctx *BoolLiteralContext)
    EnterBoolLiteral is called when production BoolLiteral is entered.

func (s *BaseArcParserListener) EnterCallExpr(ctx *CallExprContext)
    EnterCallExpr is called when production CallExpr is entered.

func (s *BaseArcParserListener) EnterCastExpr(ctx *CastExprContext)
    EnterCastExpr is called when production CastExpr is entered.

func (s *BaseArcParserListener) EnterCharLiteral(ctx *CharLiteralContext)
    EnterCharLiteral is called when production CharLiteral is entered.

func (s *BaseArcParserListener) EnterComparisonExpr(ctx *ComparisonExprContext)
    EnterComparisonExpr is called when production ComparisonExpr is entered.

func (s *BaseArcParserListener) EnterDeclaration(ctx *DeclarationContext)
    EnterDeclaration is called when production declaration is entered.

func (s *BaseArcParserListener) EnterDerefExpr(ctx *DerefExprContext)
    EnterDerefExpr is called when production DerefExpr is entered.

func (s *BaseArcParserListener) EnterEqualityExpr(ctx *EqualityExprContext)
    EnterEqualityExpr is called when production EqualityExpr is entered.

func (s *BaseArcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext)
    EnterEveryRule is called when any rule is entered.

func (s *BaseArcParserListener) EnterExprStmt(ctx *ExprStmtContext)
    EnterExprStmt is called when production ExprStmt is entered.

func (s *BaseArcParserListener) EnterExternDecl(ctx *ExternDeclContext)
    EnterExternDecl is called when production externDecl is entered.

func (s *BaseArcParserListener) EnterExternFunctionDecl(ctx *ExternFunctionDeclContext)
    EnterExternFunctionDecl is called when production externFunctionDecl is
    entered.

func (s *BaseArcParserListener) EnterExternParameterList(ctx *ExternParameterListContext)
    EnterExternParameterList is called when production externParameterList is
    entered.

func (s *BaseArcParserListener) EnterFieldAccessExpr(ctx *FieldAccessExprContext)
    EnterFieldAccessExpr is called when production FieldAccessExpr is entered.

func (s *BaseArcParserListener) EnterFieldInit(ctx *FieldInitContext)
    EnterFieldInit is called when production fieldInit is entered.

func (s *BaseArcParserListener) EnterFieldInitList(ctx *FieldInitListContext)
    EnterFieldInitList is called when production fieldInitList is entered.

func (s *BaseArcParserListener) EnterFloatLiteral(ctx *FloatLiteralContext)
    EnterFloatLiteral is called when production FloatLiteral is entered.

func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext)
    EnterFunctionDecl is called when production functionDecl is entered.

func (s *BaseArcParserListener) EnterIdentifierExpr(ctx *IdentifierExprContext)
    EnterIdentifierExpr is called when production IdentifierExpr is entered.

func (s *BaseArcParserListener) EnterIfStatement(ctx *IfStatementContext)
    EnterIfStatement is called when production ifStatement is entered.

func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext)
    EnterIfStmt is called when production IfStmt is entered.

func (s *BaseArcParserListener) EnterIntLiteral(ctx *IntLiteralContext)
    EnterIntLiteral is called when production IntLiteral is entered.

func (s *BaseArcParserListener) EnterLogicalAndExpr(ctx *LogicalAndExprContext)
    EnterLogicalAndExpr is called when production LogicalAndExpr is entered.

func (s *BaseArcParserListener) EnterLogicalNotExpr(ctx *LogicalNotExprContext)
    EnterLogicalNotExpr is called when production LogicalNotExpr is entered.

func (s *BaseArcParserListener) EnterLogicalOrExpr(ctx *LogicalOrExprContext)
    EnterLogicalOrExpr is called when production LogicalOrExpr is entered.

func (s *BaseArcParserListener) EnterMapEntry(ctx *MapEntryContext)
    EnterMapEntry is called when production mapEntry is entered.

func (s *BaseArcParserListener) EnterMapLiteral(ctx *MapLiteralContext)
    EnterMapLiteral is called when production mapLiteral is entered.

func (s *BaseArcParserListener) EnterMapLiteralExpr(ctx *MapLiteralExprContext)
    EnterMapLiteralExpr is called when production MapLiteralExpr is entered.

func (s *BaseArcParserListener) EnterMapType(ctx *MapTypeContext)
    EnterMapType is called when production MapType is entered.

func (s *BaseArcParserListener) EnterMulDivModExpr(ctx *MulDivModExprContext)
    EnterMulDivModExpr is called when production MulDivModExpr is entered.

func (s *BaseArcParserListener) EnterParameter(ctx *ParameterContext)
    EnterParameter is called when production parameter is entered.

func (s *BaseArcParserListener) EnterParameterList(ctx *ParameterListContext)
    EnterParameterList is called when production parameterList is entered.

func (s *BaseArcParserListener) EnterParenExpr(ctx *ParenExprContext)
    EnterParenExpr is called when production ParenExpr is entered.

func (s *BaseArcParserListener) EnterPointerType(ctx *PointerTypeContext)
    EnterPointerType is called when production PointerType is entered.

func (s *BaseArcParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext)
    EnterPrimaryExpr is called when production PrimaryExpr is entered.

func (s *BaseArcParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext)
    EnterPrimitiveType is called when production primitiveType is entered.

func (s *BaseArcParserListener) EnterPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext)
    EnterPrimitiveTypeSpec is called when production PrimitiveTypeSpec is
    entered.

func (s *BaseArcParserListener) EnterProgram(ctx *ProgramContext)
    EnterProgram is called when production program is entered.

func (s *BaseArcParserListener) EnterReferenceType(ctx *ReferenceTypeContext)
    EnterReferenceType is called when production ReferenceType is entered.

func (s *BaseArcParserListener) EnterReturnStmt(ctx *ReturnStmtContext)
    EnterReturnStmt is called when production ReturnStmt is entered.

func (s *BaseArcParserListener) EnterReturnType(ctx *ReturnTypeContext)
    EnterReturnType is called when production returnType is entered.

func (s *BaseArcParserListener) EnterStoreStmt(ctx *StoreStmtContext)
    EnterStoreStmt is called when production StoreStmt is entered.

func (s *BaseArcParserListener) EnterStringLiteral(ctx *StringLiteralContext)
    EnterStringLiteral is called when production StringLiteral is entered.

func (s *BaseArcParserListener) EnterStructDecl(ctx *StructDeclContext)
    EnterStructDecl is called when production structDecl is entered.

func (s *BaseArcParserListener) EnterStructField(ctx *StructFieldContext)
    EnterStructField is called when production structField is entered.

func (s *BaseArcParserListener) EnterStructFieldList(ctx *StructFieldListContext)
    EnterStructFieldList is called when production structFieldList is entered.

func (s *BaseArcParserListener) EnterStructLiteral(ctx *StructLiteralContext)
    EnterStructLiteral is called when production structLiteral is entered.

func (s *BaseArcParserListener) EnterStructLiteralExpr(ctx *StructLiteralExprContext)
    EnterStructLiteralExpr is called when production StructLiteralExpr is
    entered.

func (s *BaseArcParserListener) EnterStructType(ctx *StructTypeContext)
    EnterStructType is called when production StructType is entered.

func (s *BaseArcParserListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext)
    EnterUnaryMinusExpr is called when production UnaryMinusExpr is entered.

func (s *BaseArcParserListener) EnterVarDeclStmt(ctx *VarDeclStmtContext)
    EnterVarDeclStmt is called when production VarDeclStmt is entered.

func (s *BaseArcParserListener) EnterVariableDecl(ctx *VariableDeclContext)
    EnterVariableDecl is called when production variableDecl is entered.

func (s *BaseArcParserListener) EnterVectorLiteral(ctx *VectorLiteralContext)
    EnterVectorLiteral is called when production vectorLiteral is entered.

func (s *BaseArcParserListener) EnterVectorLiteralExpr(ctx *VectorLiteralExprContext)
    EnterVectorLiteralExpr is called when production VectorLiteralExpr is
    entered.

func (s *BaseArcParserListener) EnterVectorType(ctx *VectorTypeContext)
    EnterVectorType is called when production VectorType is entered.

func (s *BaseArcParserListener) ExitAddSubExpr(ctx *AddSubExprContext)
    ExitAddSubExpr is called when production AddSubExpr is exited.

func (s *BaseArcParserListener) ExitAddrOfExpr(ctx *AddrOfExprContext)
    ExitAddrOfExpr is called when production AddrOfExpr is exited.

func (s *BaseArcParserListener) ExitAllocaExpr(ctx *AllocaExprContext)
    ExitAllocaExpr is called when production AllocaExpr is exited.

func (s *BaseArcParserListener) ExitArgumentList(ctx *ArgumentListContext)
    ExitArgumentList is called when production argumentList is exited.

func (s *BaseArcParserListener) ExitAssignStmt(ctx *AssignStmtContext)
    ExitAssignStmt is called when production AssignStmt is exited.

func (s *BaseArcParserListener) ExitBlock(ctx *BlockContext)
    ExitBlock is called when production block is exited.

func (s *BaseArcParserListener) ExitBoolLiteral(ctx *BoolLiteralContext)
    ExitBoolLiteral is called when production BoolLiteral is exited.

func (s *BaseArcParserListener) ExitCallExpr(ctx *CallExprContext)
    ExitCallExpr is called when production CallExpr is exited.

func (s *BaseArcParserListener) ExitCastExpr(ctx *CastExprContext)
    ExitCastExpr is called when production CastExpr is exited.

func (s *BaseArcParserListener) ExitCharLiteral(ctx *CharLiteralContext)
    ExitCharLiteral is called when production CharLiteral is exited.

func (s *BaseArcParserListener) ExitComparisonExpr(ctx *ComparisonExprContext)
    ExitComparisonExpr is called when production ComparisonExpr is exited.

func (s *BaseArcParserListener) ExitDeclaration(ctx *DeclarationContext)
    ExitDeclaration is called when production declaration is exited.

func (s *BaseArcParserListener) ExitDerefExpr(ctx *DerefExprContext)
    ExitDerefExpr is called when production DerefExpr is exited.

func (s *BaseArcParserListener) ExitEqualityExpr(ctx *EqualityExprContext)
    ExitEqualityExpr is called when production EqualityExpr is exited.

func (s *BaseArcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext)
    ExitEveryRule is called when any rule is exited.

func (s *BaseArcParserListener) ExitExprStmt(ctx *ExprStmtContext)
    ExitExprStmt is called when production ExprStmt is exited.

func (s *BaseArcParserListener) ExitExternDecl(ctx *ExternDeclContext)
    ExitExternDecl is called when production externDecl is exited.

func (s *BaseArcParserListener) ExitExternFunctionDecl(ctx *ExternFunctionDeclContext)
    ExitExternFunctionDecl is called when production externFunctionDecl is
    exited.

func (s *BaseArcParserListener) ExitExternParameterList(ctx *ExternParameterListContext)
    ExitExternParameterList is called when production externParameterList is
    exited.

func (s *BaseArcParserListener) ExitFieldAccessExpr(ctx *FieldAccessExprContext)
    ExitFieldAccessExpr is called when production FieldAccessExpr is exited.

func (s *BaseArcParserListener) ExitFieldInit(ctx *FieldInitContext)
    ExitFieldInit is called when production fieldInit is exited.

func (s *BaseArcParserListener) ExitFieldInitList(ctx *FieldInitListContext)
    ExitFieldInitList is called when production fieldInitList is exited.

func (s *BaseArcParserListener) ExitFloatLiteral(ctx *FloatLiteralContext)
    ExitFloatLiteral is called when production FloatLiteral is exited.

func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext)
    ExitFunctionDecl is called when production functionDecl is exited.

func (s *BaseArcParserListener) ExitIdentifierExpr(ctx *IdentifierExprContext)
    ExitIdentifierExpr is called when production IdentifierExpr is exited.

func (s *BaseArcParserListener) ExitIfStatement(ctx *IfStatementContext)
    ExitIfStatement is called when production ifStatement is exited.

func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext)
    ExitIfStmt is called when production IfStmt is exited.

func (s *BaseArcParserListener) ExitIntLiteral(ctx *IntLiteralContext)
    ExitIntLiteral is called when production IntLiteral is exited.

func (s *BaseArcParserListener) ExitLogicalAndExpr(ctx *LogicalAndExprContext)
    ExitLogicalAndExpr is called when production LogicalAndExpr is exited.

func (s *BaseArcParserListener) ExitLogicalNotExpr(ctx *LogicalNotExprContext)
    ExitLogicalNotExpr is called when production LogicalNotExpr is exited.

func (s *BaseArcParserListener) ExitLogicalOrExpr(ctx *LogicalOrExprContext)
    ExitLogicalOrExpr is called when production LogicalOrExpr is exited.

func (s *BaseArcParserListener) ExitMapEntry(ctx *MapEntryContext)
    ExitMapEntry is called when production mapEntry is exited.

func (s *BaseArcParserListener) ExitMapLiteral(ctx *MapLiteralContext)
    ExitMapLiteral is called when production mapLiteral is exited.

func (s *BaseArcParserListener) ExitMapLiteralExpr(ctx *MapLiteralExprContext)
    ExitMapLiteralExpr is called when production MapLiteralExpr is exited.

func (s *BaseArcParserListener) ExitMapType(ctx *MapTypeContext)
    ExitMapType is called when production MapType is exited.

func (s *BaseArcParserListener) ExitMulDivModExpr(ctx *MulDivModExprContext)
    ExitMulDivModExpr is called when production MulDivModExpr is exited.

func (s *BaseArcParserListener) ExitParameter(ctx *ParameterContext)
    ExitParameter is called when production parameter is exited.

func (s *BaseArcParserListener) ExitParameterList(ctx *ParameterListContext)
    ExitParameterList is called when production parameterList is exited.

func (s *BaseArcParserListener) ExitParenExpr(ctx *ParenExprContext)
    ExitParenExpr is called when production ParenExpr is exited.

func (s *BaseArcParserListener) ExitPointerType(ctx *PointerTypeContext)
    ExitPointerType is called when production PointerType is exited.

func (s *BaseArcParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext)
    ExitPrimaryExpr is called when production PrimaryExpr is exited.

func (s *BaseArcParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext)
    ExitPrimitiveType is called when production primitiveType is exited.

func (s *BaseArcParserListener) ExitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext)
    ExitPrimitiveTypeSpec is called when production PrimitiveTypeSpec is exited.

func (s *BaseArcParserListener) ExitProgram(ctx *ProgramContext)
    ExitProgram is called when production program is exited.

func (s *BaseArcParserListener) ExitReferenceType(ctx *ReferenceTypeContext)
    ExitReferenceType is called when production ReferenceType is exited.

func (s *BaseArcParserListener) ExitReturnStmt(ctx *ReturnStmtContext)
    ExitReturnStmt is called when production ReturnStmt is exited.

func (s *BaseArcParserListener) ExitReturnType(ctx *ReturnTypeContext)
    ExitReturnType is called when production returnType is exited.

func (s *BaseArcParserListener) ExitStoreStmt(ctx *StoreStmtContext)
    ExitStoreStmt is called when production StoreStmt is exited.

func (s *BaseArcParserListener) ExitStringLiteral(ctx *StringLiteralContext)
    ExitStringLiteral is called when production StringLiteral is exited.

func (s *BaseArcParserListener) ExitStructDecl(ctx *StructDeclContext)
    ExitStructDecl is called when production structDecl is exited.

func (s *BaseArcParserListener) ExitStructField(ctx *StructFieldContext)
    ExitStructField is called when production structField is exited.

func (s *BaseArcParserListener) ExitStructFieldList(ctx *StructFieldListContext)
    ExitStructFieldList is called when production structFieldList is exited.

func (s *BaseArcParserListener) ExitStructLiteral(ctx *StructLiteralContext)
    ExitStructLiteral is called when production structLiteral is exited.

func (s *BaseArcParserListener) ExitStructLiteralExpr(ctx *StructLiteralExprContext)
    ExitStructLiteralExpr is called when production StructLiteralExpr is exited.

func (s *BaseArcParserListener) ExitStructType(ctx *StructTypeContext)
    ExitStructType is called when production StructType is exited.

func (s *BaseArcParserListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext)
    ExitUnaryMinusExpr is called when production UnaryMinusExpr is exited.

func (s *BaseArcParserListener) ExitVarDeclStmt(ctx *VarDeclStmtContext)
    ExitVarDeclStmt is called when production VarDeclStmt is exited.

func (s *BaseArcParserListener) ExitVariableDecl(ctx *VariableDeclContext)
    ExitVariableDecl is called when production variableDecl is exited.

func (s *BaseArcParserListener) ExitVectorLiteral(ctx *VectorLiteralContext)
    ExitVectorLiteral is called when production vectorLiteral is exited.

func (s *BaseArcParserListener) ExitVectorLiteralExpr(ctx *VectorLiteralExprContext)
    ExitVectorLiteralExpr is called when production VectorLiteralExpr is exited.

func (s *BaseArcParserListener) ExitVectorType(ctx *VectorTypeContext)
    ExitVectorType is called when production VectorType is exited.

func (s *BaseArcParserListener) VisitErrorNode(node antlr.ErrorNode)
    VisitErrorNode is called when an error node is visited.

func (s *BaseArcParserListener) VisitTerminal(node antlr.TerminalNode)
    VisitTerminal is called when a terminal node is visited.

type BaseArcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArcParserVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{}

func (v *BaseArcParserVisitor) VisitAddrOfExpr(ctx *AddrOfExprContext) interface{}

func (v *BaseArcParserVisitor) VisitAllocaExpr(ctx *AllocaExprContext) interface{}

func (v *BaseArcParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{}

func (v *BaseArcParserVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{}

func (v *BaseArcParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{}

func (v *BaseArcParserVisitor) VisitCastExpr(ctx *CastExprContext) interface{}

func (v *BaseArcParserVisitor) VisitCharLiteral(ctx *CharLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

func (v *BaseArcParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{}

func (v *BaseArcParserVisitor) VisitDerefExpr(ctx *DerefExprContext) interface{}

func (v *BaseArcParserVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{}

func (v *BaseArcParserVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitExternDecl(ctx *ExternDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternParameterList(ctx *ExternParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldAccessExpr(ctx *FieldAccessExprContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldInitList(ctx *FieldInitListContext) interface{}

func (v *BaseArcParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

func (v *BaseArcParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalNotExpr(ctx *LogicalNotExprContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{}

func (v *BaseArcParserVisitor) VisitMapEntry(ctx *MapEntryContext) interface{}

func (v *BaseArcParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitMapLiteralExpr(ctx *MapLiteralExprContext) interface{}

func (v *BaseArcParserVisitor) VisitMapType(ctx *MapTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitMulDivModExpr(ctx *MulDivModExprContext) interface{}

func (v *BaseArcParserVisitor) VisitParameter(ctx *ParameterContext) interface{}

func (v *BaseArcParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{}

func (v *BaseArcParserVisitor) VisitPointerType(ctx *PointerTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

func (v *BaseArcParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) interface{}

func (v *BaseArcParserVisitor) VisitProgram(ctx *ProgramContext) interface{}

func (v *BaseArcParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitStoreStmt(ctx *StoreStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitStructFieldList(ctx *StructFieldListContext) interface{}

func (v *BaseArcParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitStructLiteralExpr(ctx *StructLiteralExprContext) interface{}

func (v *BaseArcParserVisitor) VisitStructType(ctx *StructTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

func (v *BaseArcParserVisitor) VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitVectorLiteral(ctx *VectorLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitVectorLiteralExpr(ctx *VectorLiteralExprContext) interface{}

func (v *BaseArcParserVisitor) VisitVectorType(ctx *VectorTypeContext) interface{}

type BlockContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext

func NewEmptyBlockContext() *BlockContext

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BlockContext) AllStatement() []IStatementContext

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener)

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener)

func (s *BlockContext) GetParser() antlr.Parser

func (s *BlockContext) GetRuleContext() antlr.RuleContext

func (*BlockContext) IsBlockContext()

func (s *BlockContext) LBRACE() antlr.TerminalNode

func (s *BlockContext) RBRACE() antlr.TerminalNode

func (s *BlockContext) Statement(i int) IStatementContext

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BoolLiteralContext struct {
	PrimaryContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *BoolLiteralContext) FALSE() antlr.TerminalNode

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext

func (s *BoolLiteralContext) TRUE() antlr.TerminalNode

type CallExprContext struct {
	ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CallExprContext) ArgumentList() IArgumentListContext

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *CallExprContext) Expression() IExpressionContext

func (s *CallExprContext) GetRuleContext() antlr.RuleContext

func (s *CallExprContext) LPAREN() antlr.TerminalNode

func (s *CallExprContext) RPAREN() antlr.TerminalNode

type CastExprContext struct {
	ExpressionContext
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext

func (s *CastExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CastExprContext) CAST() antlr.TerminalNode

func (s *CastExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *CastExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *CastExprContext) Expression() IExpressionContext

func (s *CastExprContext) GT() antlr.TerminalNode

func (s *CastExprContext) GetRuleContext() antlr.RuleContext

func (s *CastExprContext) LPAREN() antlr.TerminalNode

func (s *CastExprContext) LT() antlr.TerminalNode

func (s *CastExprContext) RPAREN() antlr.TerminalNode

func (s *CastExprContext) Type_() ITypeContext

type CharLiteralContext struct {
	PrimaryContext
}

func NewCharLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharLiteralContext

func (s *CharLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CharLiteralContext) CHAR_LITERAL() antlr.TerminalNode

func (s *CharLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *CharLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *CharLiteralContext) GetRuleContext() antlr.RuleContext

type ComparisonExprContext struct {
	ExpressionContext
	// Has unexported fields.
}

func NewComparisonExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExprContext

func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ComparisonExprContext) AllExpression() []IExpressionContext

func (s *ComparisonExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ComparisonExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ComparisonExprContext) Expression(i int) IExpressionContext

func (s *ComparisonExprContext) GE() antlr.TerminalNode

func (s *ComparisonExprContext) GT() antlr.TerminalNode

func (s *ComparisonExprContext) GetOp() antlr.Token

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext

func (s *ComparisonExprContext) LE() antlr.TerminalNode

func (s *ComparisonExprContext) LT() antlr.TerminalNode

func (s *ComparisonExprContext) SetOp(v antlr.Token)

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext

func NewEmptyDeclarationContext() *DeclarationContext

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener)

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener)

func (s *DeclarationContext) ExternDecl() IExternDeclContext

func (s *DeclarationContext) FunctionDecl() IFunctionDeclContext

func (s *DeclarationContext) GetParser() antlr.Parser

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext

func (*DeclarationContext) IsDeclarationContext()

func (s *DeclarationContext) StructDecl() IStructDeclContext

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *DeclarationContext) VariableDecl() IVariableDeclContext

type DerefExprContext struct {
	ExpressionContext
}

func NewDerefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerefExprContext

func (s *DerefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DerefExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *DerefExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *DerefExprContext) Expression() IExpressionContext

func (s *DerefExprContext) GetRuleContext() antlr.RuleContext

func (s *DerefExprContext) STAR() antlr.TerminalNode

type EqualityExprContext struct {
	ExpressionContext
	// Has unexported fields.
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *EqualityExprContext) AllExpression() []IExpressionContext

func (s *EqualityExprContext) EQ() antlr.TerminalNode

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *EqualityExprContext) Expression(i int) IExpressionContext

func (s *EqualityExprContext) GetOp() antlr.Token

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext

func (s *EqualityExprContext) NE() antlr.TerminalNode

func (s *EqualityExprContext) SetOp(v antlr.Token)

type ExprStmtContext struct {
	StatementContext
}

func NewExprStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStmtContext

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExprStmtContext) Expression() IExpressionContext

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExpressionContext() *ExpressionContext

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext)

func (s *ExpressionContext) GetParser() antlr.Parser

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext

func (*ExpressionContext) IsExpressionContext()

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternDeclContext() *ExternDeclContext

func NewExternDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternDeclContext

func (s *ExternDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternDeclContext) AllExternFunctionDecl() []IExternFunctionDeclContext

func (s *ExternDeclContext) EXTERN() antlr.TerminalNode

func (s *ExternDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternDeclContext) ExternFunctionDecl(i int) IExternFunctionDeclContext

func (s *ExternDeclContext) GetParser() antlr.Parser

func (s *ExternDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternDeclContext) IsExternDeclContext()

func (s *ExternDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternFunctionDeclContext() *ExternFunctionDeclContext

func NewExternFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternFunctionDeclContext

func (s *ExternFunctionDeclContext) ARROW() antlr.TerminalNode

func (s *ExternFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternFunctionDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternFunctionDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternFunctionDeclContext) ExternParameterList() IExternParameterListContext

func (s *ExternFunctionDeclContext) FUNC() antlr.TerminalNode

func (s *ExternFunctionDeclContext) GetParser() antlr.Parser

func (s *ExternFunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternFunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternFunctionDeclContext) IsExternFunctionDeclContext()

func (s *ExternFunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternFunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternFunctionDeclContext) ReturnType() IReturnTypeContext

func (s *ExternFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternParameterListContext() *ExternParameterListContext

func NewExternParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternParameterListContext

func (s *ExternParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternParameterListContext) AllParameter() []IParameterContext

func (s *ExternParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ExternParameterListContext) ELLIPSIS() antlr.TerminalNode

func (s *ExternParameterListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternParameterListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternParameterListContext) GetParser() antlr.Parser

func (s *ExternParameterListContext) GetRuleContext() antlr.RuleContext

func (*ExternParameterListContext) IsExternParameterListContext()

func (s *ExternParameterListContext) Parameter(i int) IParameterContext

func (s *ExternParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FieldAccessExprContext struct {
	ExpressionContext
}

func NewFieldAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessExprContext

func (s *FieldAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FieldAccessExprContext) DOT() antlr.TerminalNode

func (s *FieldAccessExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FieldAccessExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FieldAccessExprContext) Expression() IExpressionContext

func (s *FieldAccessExprContext) GetRuleContext() antlr.RuleContext

func (s *FieldAccessExprContext) IDENTIFIER() antlr.TerminalNode

type FieldInitContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFieldInitContext() *FieldInitContext

func NewFieldInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitContext

func (s *FieldInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FieldInitContext) COLON() antlr.TerminalNode

func (s *FieldInitContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FieldInitContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FieldInitContext) Expression() IExpressionContext

func (s *FieldInitContext) GetParser() antlr.Parser

func (s *FieldInitContext) GetRuleContext() antlr.RuleContext

func (s *FieldInitContext) IDENTIFIER() antlr.TerminalNode

func (*FieldInitContext) IsFieldInitContext()

func (s *FieldInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FieldInitListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFieldInitListContext() *FieldInitListContext

func NewFieldInitListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitListContext

func (s *FieldInitListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FieldInitListContext) AllCOMMA() []antlr.TerminalNode

func (s *FieldInitListContext) AllFieldInit() []IFieldInitContext

func (s *FieldInitListContext) COMMA(i int) antlr.TerminalNode

func (s *FieldInitListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FieldInitListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FieldInitListContext) FieldInit(i int) IFieldInitContext

func (s *FieldInitListContext) GetParser() antlr.Parser

func (s *FieldInitListContext) GetRuleContext() antlr.RuleContext

func (*FieldInitListContext) IsFieldInitListContext()

func (s *FieldInitListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FloatLiteralContext struct {
	PrimaryContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext

func (s *FunctionDeclContext) ARROW() antlr.TerminalNode

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FunctionDeclContext) Block() IBlockContext

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode

func (s *FunctionDeclContext) GetParser() antlr.Parser

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*FunctionDeclContext) IsFunctionDeclContext()

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) ParameterList() IParameterListContext

func (s *FunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) ReturnType() IReturnTypeContext

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}
    IArgumentListContext is an interface to support dynamic dispatch.

type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}
    IBlockContext is an interface to support dynamic dispatch.

type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	ExternDecl() IExternDeclContext
	StructDecl() IStructDeclContext
	VariableDecl() IVariableDeclContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}
    IDeclarationContext is an interface to support dynamic dispatch.

type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}
    IExpressionContext is an interface to support dynamic dispatch.

type IExternDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternFunctionDecl() []IExternFunctionDeclContext
	ExternFunctionDecl(i int) IExternFunctionDeclContext

	// IsExternDeclContext differentiates from other interfaces.
	IsExternDeclContext()
}
    IExternDeclContext is an interface to support dynamic dispatch.

type IExternFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExternParameterList() IExternParameterListContext
	ReturnType() IReturnTypeContext
	ARROW() antlr.TerminalNode

	// IsExternFunctionDeclContext differentiates from other interfaces.
	IsExternFunctionDeclContext()
}
    IExternFunctionDeclContext is an interface to support dynamic dispatch.

type IExternParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsExternParameterListContext differentiates from other interfaces.
	IsExternParameterListContext()
}
    IExternParameterListContext is an interface to support dynamic dispatch.

type IFieldInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFieldInitContext differentiates from other interfaces.
	IsFieldInitContext()
}
    IFieldInitContext is an interface to support dynamic dispatch.

type IFieldInitListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldInit() []IFieldInitContext
	FieldInit(i int) IFieldInitContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldInitListContext differentiates from other interfaces.
	IsFieldInitListContext()
}
    IFieldInitListContext is an interface to support dynamic dispatch.

type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext
	ARROW() antlr.TerminalNode

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}
    IFunctionDeclContext is an interface to support dynamic dispatch.

type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}
    IIfStatementContext is an interface to support dynamic dispatch.

type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}
    IMapEntryContext is an interface to support dynamic dispatch.

type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}
    IMapLiteralContext is an interface to support dynamic dispatch.

type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}
    IParameterContext is an interface to support dynamic dispatch.

type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}
    IParameterListContext is an interface to support dynamic dispatch.

type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}
    IPrimaryContext is an interface to support dynamic dispatch.

type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTE() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}
    IPrimitiveTypeContext is an interface to support dynamic dispatch.

type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}
    IProgramContext is an interface to support dynamic dispatch.

type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}
    IReturnTypeContext is an interface to support dynamic dispatch.

type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}
    IStatementContext is an interface to support dynamic dispatch.

type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	StructFieldList() IStructFieldListContext
	RBRACE() antlr.TerminalNode

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}
    IStructDeclContext is an interface to support dynamic dispatch.

type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}
    IStructFieldContext is an interface to support dynamic dispatch.

type IStructFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext

	// IsStructFieldListContext differentiates from other interfaces.
	IsStructFieldListContext()
}
    IStructFieldListContext is an interface to support dynamic dispatch.

type IStructLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	FieldInitList() IFieldInitListContext

	// IsStructLiteralContext differentiates from other interfaces.
	IsStructLiteralContext()
}
    IStructLiteralContext is an interface to support dynamic dispatch.

type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}
    ITypeContext is an interface to support dynamic dispatch.

type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	CONST() antlr.TerminalNode

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}
    IVariableDeclContext is an interface to support dynamic dispatch.

type IVectorLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVectorLiteralContext differentiates from other interfaces.
	IsVectorLiteralContext()
}
    IVectorLiteralContext is an interface to support dynamic dispatch.

type IdentifierExprContext struct {
	PrimaryContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext

func (s *IdentifierExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext

func (s *IdentifierExprContext) IDENTIFIER() antlr.TerminalNode

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyIfStatementContext() *IfStatementContext

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IfStatementContext) AllBlock() []IBlockContext

func (s *IfStatementContext) Block(i int) IBlockContext

func (s *IfStatementContext) ELSE() antlr.TerminalNode

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener)

func (s *IfStatementContext) Expression() IExpressionContext

func (s *IfStatementContext) GetParser() antlr.Parser

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext

func (s *IfStatementContext) IF() antlr.TerminalNode

func (*IfStatementContext) IsIfStatementContext()

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type IfStmtContext struct {
	StatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext

func (s *IfStmtContext) IfStatement() IIfStatementContext

type IntLiteralContext struct {
	PrimaryContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext

func (s *IntLiteralContext) INTEGER_LITERAL() antlr.TerminalNode

type LogicalAndExprContext struct {
	ExpressionContext
}

func NewLogicalAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExprContext

func (s *LogicalAndExprContext) AND() antlr.TerminalNode

func (s *LogicalAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalAndExprContext) AllExpression() []IExpressionContext

func (s *LogicalAndExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LogicalAndExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *LogicalAndExprContext) Expression(i int) IExpressionContext

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext

type LogicalNotExprContext struct {
	ExpressionContext
}

func NewLogicalNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalNotExprContext

func (s *LogicalNotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalNotExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LogicalNotExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *LogicalNotExprContext) Expression() IExpressionContext

func (s *LogicalNotExprContext) GetRuleContext() antlr.RuleContext

func (s *LogicalNotExprContext) NOT() antlr.TerminalNode

type LogicalOrExprContext struct {
	ExpressionContext
}

func NewLogicalOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExprContext

func (s *LogicalOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalOrExprContext) AllExpression() []IExpressionContext

func (s *LogicalOrExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LogicalOrExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *LogicalOrExprContext) Expression(i int) IExpressionContext

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext

func (s *LogicalOrExprContext) OR() antlr.TerminalNode

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMapEntryContext() *MapEntryContext

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapEntryContext) AllExpression() []IExpressionContext

func (s *MapEntryContext) COLON() antlr.TerminalNode

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapEntryContext) Expression(i int) IExpressionContext

func (s *MapEntryContext) GetParser() antlr.Parser

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext

func (*MapEntryContext) IsMapEntryContext()

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMapLiteralContext() *MapLiteralContext

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode

func (s *MapLiteralContext) AllMapEntry() []IMapEntryContext

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapLiteralContext) GetParser() antlr.Parser

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext

func (*MapLiteralContext) IsMapLiteralContext()

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode

func (s *MapLiteralContext) MapEntry(i int) IMapEntryContext

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MapLiteralExprContext struct {
	PrimaryContext
}

func NewMapLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLiteralExprContext

func (s *MapLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapLiteralExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapLiteralExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapLiteralExprContext) GetRuleContext() antlr.RuleContext

func (s *MapLiteralExprContext) MapLiteral() IMapLiteralContext

type MapTypeContext struct {
	TypeContext
}

func NewMapTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapTypeContext

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapTypeContext) AllType_() []ITypeContext

func (s *MapTypeContext) COMMA() antlr.TerminalNode

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapTypeContext) GT() antlr.TerminalNode

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext

func (s *MapTypeContext) LT() antlr.TerminalNode

func (s *MapTypeContext) MAP() antlr.TerminalNode

func (s *MapTypeContext) Type_(i int) ITypeContext

type MulDivModExprContext struct {
	ExpressionContext
	// Has unexported fields.
}

func NewMulDivModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModExprContext

func (s *MulDivModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MulDivModExprContext) AllExpression() []IExpressionContext

func (s *MulDivModExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MulDivModExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MulDivModExprContext) Expression(i int) IExpressionContext

func (s *MulDivModExprContext) GetOp() antlr.Token

func (s *MulDivModExprContext) GetRuleContext() antlr.RuleContext

func (s *MulDivModExprContext) PERCENT() antlr.TerminalNode

func (s *MulDivModExprContext) SLASH() antlr.TerminalNode

func (s *MulDivModExprContext) STAR() antlr.TerminalNode

func (s *MulDivModExprContext) SetOp(v antlr.Token)

type ParameterContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyParameterContext() *ParameterContext

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ParameterContext) COLON() antlr.TerminalNode

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ParameterContext) GetParser() antlr.Parser

func (s *ParameterContext) GetRuleContext() antlr.RuleContext

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode

func (*ParameterContext) IsParameterContext()

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ParameterContext) Type_() ITypeContext

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyParameterListContext() *ParameterListContext

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ParameterListContext) AllParameter() []IParameterContext

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ParameterListContext) GetParser() antlr.Parser

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext

func (*ParameterListContext) IsParameterListContext()

func (s *ParameterListContext) Parameter(i int) IParameterContext

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ParenExprContext struct {
	PrimaryContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ParenExprContext) Expression() IExpressionContext

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext

func (s *ParenExprContext) LPAREN() antlr.TerminalNode

func (s *ParenExprContext) RPAREN() antlr.TerminalNode

type PointerTypeContext struct {
	TypeContext
}

func NewPointerTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PointerTypeContext

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext

func (s *PointerTypeContext) STAR() antlr.TerminalNode

func (s *PointerTypeContext) Type_() ITypeContext

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPrimaryContext() *PrimaryContext

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext)

func (s *PrimaryContext) GetParser() antlr.Parser

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext

func (*PrimaryContext) IsPrimaryContext()

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PrimaryExprContext struct {
	ExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext

func (s *PrimaryExprContext) Primary() IPrimaryContext

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimitiveTypeContext) BOOL() antlr.TerminalNode

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *PrimitiveTypeContext) FLOAT32() antlr.TerminalNode

func (s *PrimitiveTypeContext) FLOAT64() antlr.TerminalNode

func (s *PrimitiveTypeContext) GetParser() antlr.Parser

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext

func (s *PrimitiveTypeContext) INT32() antlr.TerminalNode

func (s *PrimitiveTypeContext) INT64() antlr.TerminalNode

func (*PrimitiveTypeContext) IsPrimitiveTypeContext()

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *PrimitiveTypeContext) VOID() antlr.TerminalNode

type PrimitiveTypeSpecContext struct {
	TypeContext
}

func NewPrimitiveTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimitiveTypeSpecContext

func (s *PrimitiveTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimitiveTypeSpecContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PrimitiveTypeSpecContext) ExitRule(listener antlr.ParseTreeListener)

func (s *PrimitiveTypeSpecContext) GetRuleContext() antlr.RuleContext

func (s *PrimitiveTypeSpecContext) PrimitiveType() IPrimitiveTypeContext

type ProgramContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyProgramContext() *ProgramContext

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext

func (s *ProgramContext) Declaration(i int) IDeclarationContext

func (s *ProgramContext) EOF() antlr.TerminalNode

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ProgramContext) GetParser() antlr.Parser

func (s *ProgramContext) GetRuleContext() antlr.RuleContext

func (*ProgramContext) IsProgramContext()

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ReferenceTypeContext struct {
	TypeContext
}

func NewReferenceTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceTypeContext

func (s *ReferenceTypeContext) AMP() antlr.TerminalNode

func (s *ReferenceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReferenceTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ReferenceTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ReferenceTypeContext) GetRuleContext() antlr.RuleContext

func (s *ReferenceTypeContext) Type_() ITypeContext

type ReturnStmtContext struct {
	StatementContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ReturnStmtContext) Expression() IExpressionContext

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyReturnTypeContext() *ReturnTypeContext

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext

func (s *ReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ReturnTypeContext) GetParser() antlr.Parser

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext

func (*ReturnTypeContext) IsReturnTypeContext()

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ReturnTypeContext) Type_() ITypeContext

type StatementContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStatementContext() *StatementContext

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext

func (s *StatementContext) CopyAll(ctx *StatementContext)

func (s *StatementContext) GetParser() antlr.Parser

func (s *StatementContext) GetRuleContext() antlr.RuleContext

func (*StatementContext) IsStatementContext()

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StoreStmtContext struct {
	StatementContext
}

func NewStoreStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StoreStmtContext

func (s *StoreStmtContext) ASSIGN() antlr.TerminalNode

func (s *StoreStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StoreStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StoreStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StoreStmtContext) Expression() IExpressionContext

func (s *StoreStmtContext) GetRuleContext() antlr.RuleContext

func (s *StoreStmtContext) IDENTIFIER() antlr.TerminalNode

func (s *StoreStmtContext) STAR() antlr.TerminalNode

type StringLiteralContext struct {
	PrimaryContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructDeclContext() *StructDeclContext

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructDeclContext) GetParser() antlr.Parser

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext

func (s *StructDeclContext) IDENTIFIER() antlr.TerminalNode

func (*StructDeclContext) IsStructDeclContext()

func (s *StructDeclContext) LBRACE() antlr.TerminalNode

func (s *StructDeclContext) RBRACE() antlr.TerminalNode

func (s *StructDeclContext) STRUCT() antlr.TerminalNode

func (s *StructDeclContext) StructFieldList() IStructFieldListContext

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructFieldContext() *StructFieldContext

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructFieldContext) COLON() antlr.TerminalNode

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructFieldContext) GetParser() antlr.Parser

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext

func (s *StructFieldContext) IDENTIFIER() antlr.TerminalNode

func (*StructFieldContext) IsStructFieldContext()

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *StructFieldContext) Type_() ITypeContext

type StructFieldListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructFieldListContext() *StructFieldListContext

func NewStructFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldListContext

func (s *StructFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructFieldListContext) AllStructField() []IStructFieldContext

func (s *StructFieldListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructFieldListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructFieldListContext) GetParser() antlr.Parser

func (s *StructFieldListContext) GetRuleContext() antlr.RuleContext

func (*StructFieldListContext) IsStructFieldListContext()

func (s *StructFieldListContext) StructField(i int) IStructFieldContext

func (s *StructFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructLiteralContext() *StructLiteralContext

func NewStructLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructLiteralContext

func (s *StructLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructLiteralContext) FieldInitList() IFieldInitListContext

func (s *StructLiteralContext) GetParser() antlr.Parser

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext

func (s *StructLiteralContext) IDENTIFIER() antlr.TerminalNode

func (*StructLiteralContext) IsStructLiteralContext()

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructLiteralExprContext struct {
	PrimaryContext
}

func NewStructLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructLiteralExprContext

func (s *StructLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructLiteralExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructLiteralExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructLiteralExprContext) GetRuleContext() antlr.RuleContext

func (s *StructLiteralExprContext) StructLiteral() IStructLiteralContext

type StructTypeContext struct {
	TypeContext
}

func NewStructTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeContext

func (s *StructTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext

func (s *StructTypeContext) IDENTIFIER() antlr.TerminalNode

type TypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTypeContext() *TypeContext

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext

func (s *TypeContext) CopyAll(ctx *TypeContext)

func (s *TypeContext) GetParser() antlr.Parser

func (s *TypeContext) GetRuleContext() antlr.RuleContext

func (*TypeContext) IsTypeContext()

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type UnaryMinusExprContext struct {
	ExpressionContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *UnaryMinusExprContext) Expression() IExpressionContext

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode

type VarDeclStmtContext struct {
	StatementContext
}

func NewVarDeclStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclStmtContext

func (s *VarDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VarDeclStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VarDeclStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VarDeclStmtContext) GetRuleContext() antlr.RuleContext

func (s *VarDeclStmtContext) VariableDecl() IVariableDeclContext

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyVariableDeclContext() *VariableDeclContext

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext

func (s *VariableDeclContext) ASSIGN() antlr.TerminalNode

func (s *VariableDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VariableDeclContext) COLON() antlr.TerminalNode

func (s *VariableDeclContext) CONST() antlr.TerminalNode

func (s *VariableDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VariableDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VariableDeclContext) Expression() IExpressionContext

func (s *VariableDeclContext) GetParser() antlr.Parser

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode

func (*VariableDeclContext) IsVariableDeclContext()

func (s *VariableDeclContext) LET() antlr.TerminalNode

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *VariableDeclContext) Type_() ITypeContext

type VectorLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyVectorLiteralContext() *VectorLiteralContext

func NewVectorLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorLiteralContext

func (s *VectorLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VectorLiteralContext) AllCOMMA() []antlr.TerminalNode

func (s *VectorLiteralContext) AllExpression() []IExpressionContext

func (s *VectorLiteralContext) COMMA(i int) antlr.TerminalNode

func (s *VectorLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralContext) Expression(i int) IExpressionContext

func (s *VectorLiteralContext) GetParser() antlr.Parser

func (s *VectorLiteralContext) GetRuleContext() antlr.RuleContext

func (*VectorLiteralContext) IsVectorLiteralContext()

func (s *VectorLiteralContext) LBRACE() antlr.TerminalNode

func (s *VectorLiteralContext) RBRACE() antlr.TerminalNode

func (s *VectorLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type VectorLiteralExprContext struct {
	PrimaryContext
}

func NewVectorLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorLiteralExprContext

func (s *VectorLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VectorLiteralExprContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralExprContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralExprContext) GetRuleContext() antlr.RuleContext

func (s *VectorLiteralExprContext) VectorLiteral() IVectorLiteralContext

type VectorTypeContext struct {
	TypeContext
}

func NewVectorTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorTypeContext

func (s *VectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VectorTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VectorTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VectorTypeContext) GT() antlr.TerminalNode

func (s *VectorTypeContext) GetRuleContext() antlr.RuleContext

func (s *VectorTypeContext) LT() antlr.TerminalNode

func (s *VectorTypeContext) Type_() ITypeContext

func (s *VectorTypeContext) VECTOR() antlr.TerminalNode

