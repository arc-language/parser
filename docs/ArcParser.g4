parser grammar ArcParser;

options {
    tokenVocab = ArcLexer;
}

// Entry point
program
    : (declaration)* EOF
    ;

// Top-level declarations
declaration
    : functionDecl
    | externDecl        // Added
    | structDecl
    | variableDecl
    ;

// Variable declarations
variableDecl
    : LET IDENTIFIER (COLON type)? ASSIGN expression
    | CONST IDENTIFIER (COLON type)? ASSIGN expression
    ;

// Function declaration
// Added optional ARROW ('->') to support 'func name() -> int' style if desired
functionDecl
    : FUNC IDENTIFIER LPAREN parameterList? RPAREN (ARROW? returnType)? block
    ;

// Extern declarations (Added)
externDecl
    : EXTERN IDENTIFIER LBRACE externFunctionDecl* RBRACE
    ;

externFunctionDecl
    : FUNC IDENTIFIER LPAREN externParameterList? RPAREN (ARROW? returnType)?
    ;

parameterList
    : parameter (COMMA parameter)*
    ;

// Externs support variadic arguments (...)
externParameterList
    : parameter (COMMA parameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

parameter
    : IDENTIFIER COLON type
    ;

returnType
    : type
    ;

// Struct declaration
structDecl
    : STRUCT IDENTIFIER LBRACE structFieldList RBRACE
    ;

structFieldList
    : structField (structField)*
    ;

structField
    : IDENTIFIER COLON type
    ;

// Block of statements
block
    : LBRACE (statement)* RBRACE
    ;

// Statements
statement
    : variableDecl                                          # VarDeclStmt
    | IDENTIFIER ASSIGN expression                          # AssignStmt
    | STAR IDENTIFIER ASSIGN expression                     # StoreStmt
    | RETURN expression?                                    # ReturnStmt
    | ifStatement                                           # IfStmt
    | expression                                            # ExprStmt
    ;

ifStatement
    : IF expression block (ELSE block)?
    ;

// Types
type
    : primitiveType                                         # PrimitiveTypeSpec
    | STAR type                                             # PointerType
    | AMP type                                              # ReferenceType
    | VECTOR LT type GT                                     # VectorType
    | MAP LT type COMMA type GT                             # MapType
    | IDENTIFIER                                            # StructType
    ;

primitiveType
    : INT32
    | INT64
    | FLOAT32
    | FLOAT64
    | BOOL
    | STRING
    | BYTE
    | CHAR
    | VOID  // Added
    ;

// Expressions (with precedence)
expression
    : primary                                               # PrimaryExpr
    | STAR expression                                       # DerefExpr
    | AMP expression                                        # AddrOfExpr
    | MINUS expression                                      # UnaryMinusExpr
    | NOT expression                                        # LogicalNotExpr
    | ALLOCA LPAREN type RPAREN                             # AllocaExpr
    | CAST LT type GT LPAREN expression RPAREN              # CastExpr
    | expression op=(STAR | SLASH | PERCENT) expression     # MulDivModExpr
    | expression op=(PLUS | MINUS) expression               # AddSubExpr
    | expression op=(LT | LE | GT | GE) expression          # ComparisonExpr
    | expression op=(EQ | NE) expression                    # EqualityExpr
    | expression AND expression                             # LogicalAndExpr
    | expression OR expression                              # LogicalOrExpr
    // Moved CallExpr here and changed to 'expression' to support io.printf(...)
    | expression LPAREN argumentList? RPAREN                # CallExpr 
    | expression DOT IDENTIFIER                             # FieldAccessExpr
    ;

primary
    : INTEGER_LITERAL                                       # IntLiteral
    | FLOAT_LITERAL                                         # FloatLiteral
    | STRING_LITERAL                                        # StringLiteral
    | CHAR_LITERAL                                          # CharLiteral
    | TRUE                                                  # BoolLiteral
    | FALSE                                                 # BoolLiteral
    | IDENTIFIER                                            # IdentifierExpr
    | structLiteral                                         # StructLiteralExpr
    | vectorLiteral                                         # VectorLiteralExpr
    | mapLiteral                                            # MapLiteralExpr
    | LPAREN expression RPAREN                              # ParenExpr
    ;

// Literals
structLiteral
    : IDENTIFIER LBRACE fieldInitList? RBRACE
    ;

fieldInitList
    : fieldInit (COMMA fieldInit)*
    ;

fieldInit
    : IDENTIFIER COLON expression
    ;

vectorLiteral
    : LBRACE (expression (COMMA expression)*)? RBRACE
    ;

mapLiteral
    : LBRACE (mapEntry (COMMA mapEntry)*)? RBRACE
    ;

mapEntry
    : expression COLON expression
    ;

argumentList
    : expression (COMMA expression)*
    ;