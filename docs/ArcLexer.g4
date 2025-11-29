lexer grammar ArcLexer;

// Keywords
LET: 'let';
CONST: 'const';
FUNC: 'func';
EXTERN: 'extern'; // Added
STRUCT: 'struct';
RETURN: 'return';
IF: 'if';
ELSE: 'else';
ALLOCA: 'alloca';
CAST: 'cast';

// Primitive Types
INT32: 'int32';
INT64: 'int64';
FLOAT32: 'float32';
FLOAT64: 'float64';
BOOL: 'bool';
STRING: 'string';
BYTE: 'byte';
CHAR: 'char';
VOID: 'void'; // Added for C interop

// Collection Types
VECTOR: 'vector';
MAP: 'map';

// Boolean Literals
TRUE: 'true';
FALSE: 'false';

// Operators
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
PERCENT: '%';

EQ: '==';
NE: '!=';
LT: '<';
LE: '<=';
GT: '>';
GE: '>=';

AND: '&&';
OR: '||';
NOT: '!';

AMP: '&';
AT: '@';

ASSIGN: '=';
ARROW: '->'; // Added
ELLIPSIS: '...'; // Added for variadic functions

// Delimiters
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACK: '[';
RBRACK: ']';

COMMA: ',';
COLON: ':';
SEMICOLON: ';';
DOT: '.';

// Identifiers
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

// Literals
INTEGER_LITERAL: [0-9]+;
FLOAT_LITERAL: [0-9]+ '.' [0-9]+;
STRING_LITERAL: '"' (~["\r\n] | '\\' .)* '"';
CHAR_LITERAL: '\'' (~['\r\n] | '\\' .) '\'';

// Whitespace and Comments
WS: [ \t\r\n]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;