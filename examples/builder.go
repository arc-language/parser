package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/builder"
	"github.com/arc-language/ir"
	"github.com/arc-language/parser"
	"github.com/arc-language/types"
)

type IRGenerator struct {
	*parser.BaseArcParserVisitor
	builder      *builder.Builder
	currentFunc  *ir.Function
	symbolTable  map[string]*builder.Value
	functionMap  map[string]*ir.Function
	breakTargets []*ir.BasicBlock
	contTargets  []*ir.BasicBlock
}

func NewIRGenerator() *IRGenerator {
	gen := &IRGenerator{
		builder:     builder.New(),
		symbolTable: make(map[string]*builder.Value),
		functionMap: make(map[string]*ir.Function),
	}
	gen.BaseArcParserVisitor = &parser.BaseArcParserVisitor{}
	return gen
}

// VisitProgram handles the top-level program
func (g *IRGenerator) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("DEBUG: VisitProgram called")
	declarations := ctx.AllDeclaration()
	fmt.Printf("DEBUG: Found %d declarations\n", len(declarations))
	
	// First pass: declare all functions
	for i, declCtx := range declarations {
		fmt.Printf("DEBUG: Processing declaration %d\n", i)
		decl := declCtx.(*parser.DeclarationContext)
		if funcDecl := decl.FunctionDecl(); funcDecl != nil {
			fmt.Printf("DEBUG: Found function declaration\n")
			g.declareFunctionSignature(funcDecl.(*parser.FunctionDeclContext))
		}
	}

	// Second pass: generate function bodies
	for _, declCtx := range declarations {
		declCtx.Accept(g) // Use Accept instead of Visit
	}

	return nil
}

func (g *IRGenerator) declareFunctionSignature(ctx *parser.FunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	fmt.Printf("DEBUG: Declaring function: %s\n", funcName)
	
	// Get return type
	var returnType types.Type
	if ctx.ReturnType() != nil {
		returnType = g.getType(ctx.ReturnType().Type_())
		fmt.Printf("DEBUG: Return type: %v\n", returnType)
	} else {
		returnType = types.Int32 // default
		fmt.Printf("DEBUG: Using default return type: %v\n", returnType)
	}

	// Get parameter types
	var paramTypes []types.Type
	if ctx.ParameterList() != nil {
		paramList := ctx.ParameterList().(*parser.ParameterListContext)
		fmt.Printf("DEBUG: Found %d parameters\n", len(paramList.AllParameter()))
		for _, paramCtx := range paramList.AllParameter() {
			param := paramCtx.(*parser.ParameterContext)
			paramType := g.getType(param.Type_())
			paramTypes = append(paramTypes, paramType)
		}
	}

	fn := g.builder.CreateFunction(funcName, returnType, paramTypes)
	g.functionMap[funcName] = fn
	fmt.Printf("DEBUG: Function %s created successfully\n", funcName)
}

// VisitDeclaration handles declarations
func (g *IRGenerator) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	if funcDecl := ctx.FunctionDecl(); funcDecl != nil {
		return g.VisitFunctionDecl(funcDecl.(*parser.FunctionDeclContext))
	}
	if varDecl := ctx.VariableDecl(); varDecl != nil {
		return g.VisitVariableDecl(varDecl.(*parser.VariableDeclContext))
	}
	if structDecl := ctx.StructDecl(); structDecl != nil {
		// Handle struct declarations later
		return nil
	}
	return nil
}

// VisitFunctionDecl handles function declarations
func (g *IRGenerator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	funcName := ctx.IDENTIFIER().GetText()
	
	fn := g.functionMap[funcName]
	g.currentFunc = fn

	// Note: CreateFunction already created the entry block and set it as current
	// No need to create it again here

	// Create new scope
	prevSymbols := g.symbolTable
	g.symbolTable = make(map[string]*builder.Value)

	// Map parameters
	if ctx.ParameterList() != nil {
		paramList := ctx.ParameterList().(*parser.ParameterListContext)
		for i, paramCtx := range paramList.AllParameter() {
			param := paramCtx.(*parser.ParameterContext)
			paramName := param.IDENTIFIER().GetText()
			paramValue := g.builder.GetParam(i)
			g.symbolTable[paramName] = paramValue
		}
	}

	// Generate function body
	if ctx.Block() != nil {
		ctx.Block().Accept(g)
	}

	// Restore scope
	g.symbolTable = prevSymbols
	g.currentFunc = nil

	return nil
}

// VisitVariableDecl handles variable declarations
func (g *IRGenerator) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	varName := ctx.IDENTIFIER().GetText()
	varType := g.getType(ctx.Type_())

	// Allocate space
	alloca := g.builder.CreateAlloca(varType)

	// Initialize if present
	if ctx.Expression() != nil {
		initValueInterface := ctx.Expression().Accept(g)
		if initValueInterface == nil {
			fmt.Fprintf(os.Stderr, "Error: expression returned nil for variable %s (expr type: %T)\n", 
				varName, ctx.Expression())
			// Create a default value based on type
			initValueInterface = g.getDefaultValue(varType)
		}
		initValue := initValueInterface.(*builder.Value)
		g.builder.CreateStore(alloca, initValue)
	}

	g.symbolTable[varName] = alloca
	return nil
}

// VisitBlock handles code blocks
func (g *IRGenerator) VisitBlock(ctx *parser.BlockContext) interface{} {
	// Create new scope
	prevSymbols := g.symbolTable
	newSymbols := make(map[string]*builder.Value)
	for k, v := range prevSymbols {
		newSymbols[k] = v
	}
	g.symbolTable = newSymbols

	// Visit all statements
	for _, stmtCtx := range ctx.AllStatement() {
		stmtCtx.Accept(g)
	}

	// Restore scope
	g.symbolTable = prevSymbols
	return nil
}

// VisitReturnStmt handles return statements
func (g *IRGenerator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if ctx.Expression() != nil {
		valueInterface := ctx.Expression().Accept(g)
		if valueInterface == nil {
			fmt.Fprintf(os.Stderr, "Error: return expression returned nil\n")
			// Return a default value
			valueInterface = g.builder.CreateInt32("", 0)
		}
		value := valueInterface.(*builder.Value)
		g.builder.CreateReturn(value)
	} else {
		g.builder.CreateReturn(nil)
	}
	return nil
}

// VisitVarDeclStmt handles variable declaration statements
func (g *IRGenerator) VisitVarDeclStmt(ctx *parser.VarDeclStmtContext) interface{} {
	return g.VisitVariableDecl(ctx.VariableDecl().(*parser.VariableDeclContext))
}

// VisitAssignStmt handles assignments
func (g *IRGenerator) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	// Get variable name and value
	varName := ctx.IDENTIFIER().GetText()
	valueInterface := ctx.Expression().Accept(g)
	
	if valueInterface == nil {
		fmt.Fprintf(os.Stderr, "Error: assignment expression returned nil for variable %s\n", varName)
		valueInterface = g.builder.CreateInt32("", 0)
	}
	
	value := valueInterface.(*builder.Value)

	ptr, exists := g.symbolTable[varName]
	if !exists {
		fmt.Fprintf(os.Stderr, "Undefined variable: %s\n", varName)
		return nil
	}

	g.builder.CreateStore(ptr, value)
	return nil
}

// VisitIfStatement handles if statements
func (g *IRGenerator) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	conditionInterface := ctx.Expression().Accept(g)
	if conditionInterface == nil {
		fmt.Fprintf(os.Stderr, "Error: if condition returned nil\n")
		conditionInterface = g.builder.CreateBool("", false)
	}
	condition := conditionInterface.(*builder.Value)

	thenBlock := g.builder.CreateBasicBlock("if.then")
	elseBlock := g.builder.CreateBasicBlock("if.else")
	mergeBlock := g.builder.CreateBasicBlock("if.merge")

	g.builder.CreateBranch(condition, thenBlock, elseBlock)

	// Then block
	g.currentFunc.AddBlock(thenBlock)
	g.builder.SetInsertPoint(thenBlock)
	ctx.Block(0).Accept(g)
	g.builder.CreateUnconditionalBranch(mergeBlock)

	// Else block
	g.currentFunc.AddBlock(elseBlock)
	g.builder.SetInsertPoint(elseBlock)
	if ctx.Block(1) != nil {
		ctx.Block(1).Accept(g)
	}
	g.builder.CreateUnconditionalBranch(mergeBlock)

	// Merge block
	g.currentFunc.AddBlock(mergeBlock)
	g.builder.SetInsertPoint(mergeBlock)

	return nil
}

// Visit expression contexts
func (g *IRGenerator) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	text := ctx.INTEGER_LITERAL().GetText()
	var value int64
	fmt.Sscanf(text, "%d", &value)
	return g.builder.CreateInt64("", value)
}

func (g *IRGenerator) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	text := ctx.FLOAT_LITERAL().GetText()
	var value float64
	fmt.Sscanf(text, "%f", &value)
	return g.builder.CreateFloat64("", value)
}

func (g *IRGenerator) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	if ctx.TRUE() != nil {
		return g.builder.CreateBool("", true)
	}
	return g.builder.CreateBool("", false)
}

func (g *IRGenerator) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	text := ctx.STRING_LITERAL().GetText()
	// Remove quotes
	if len(text) >= 2 {
		text = text[1 : len(text)-1]
	}
	return g.builder.CreateString("", text)
}

func (g *IRGenerator) VisitIdentifierExpr(ctx *parser.IdentifierExprContext) interface{} {
	varName := ctx.IDENTIFIER().GetText()
	ptr, exists := g.symbolTable[varName]
	if !exists {
		fmt.Fprintf(os.Stderr, "Undefined variable: %s\n", varName)
		return g.builder.CreateInt32("", 0) // error fallback
	}
	return g.builder.CreateLoad(ptr)
}

func (g *IRGenerator) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "add/sub left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "add/sub right operand")
	
	op := "+"
	if ctx.MINUS() != nil {
		op = "-"
	}
	
	return g.builder.CreateBinaryOp(op, left, right)
}

func (g *IRGenerator) VisitMulDivModExpr(ctx *parser.MulDivModExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "mul/div/mod left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "mul/div/mod right operand")
	
	op := "*"
	if ctx.SLASH() != nil {
		op = "/"
	} else if ctx.PERCENT() != nil {
		op = "%"
	}
	
	return g.builder.CreateBinaryOp(op, left, right)
}

func (g *IRGenerator) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "comparison left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "comparison right operand")
	
	op := "<"
	if ctx.LE() != nil {
		op = "<="
	} else if ctx.GT() != nil {
		op = ">"
	} else if ctx.GE() != nil {
		op = ">="
	}
	
	return g.builder.CreateBinaryOp(op, left, right)
}

func (g *IRGenerator) VisitEqualityExpr(ctx *parser.EqualityExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "equality left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "equality right operand")
	
	op := "=="
	if ctx.NE() != nil {
		op = "!="
	}
	
	return g.builder.CreateBinaryOp(op, left, right)
}

func (g *IRGenerator) VisitLogicalAndExpr(ctx *parser.LogicalAndExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "logical and left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "logical and right operand")
	return g.builder.CreateBinaryOp("&&", left, right)
}

func (g *IRGenerator) VisitLogicalOrExpr(ctx *parser.LogicalOrExprContext) interface{} {
	left := g.safeAcceptExpr(ctx.Expression(0), "logical or left operand")
	right := g.safeAcceptExpr(ctx.Expression(1), "logical or right operand")
	return g.builder.CreateBinaryOp("||", left, right)
}

func (g *IRGenerator) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) interface{} {
	operand := g.safeAcceptExpr(ctx.Expression(), "unary minus operand")
	return g.builder.CreateUnaryOp("-", operand)
}

func (g *IRGenerator) VisitLogicalNotExpr(ctx *parser.LogicalNotExprContext) interface{} {
	operand := g.safeAcceptExpr(ctx.Expression(), "logical not operand")
	return g.builder.CreateUnaryOp("!", operand)
}

func (g *IRGenerator) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return g.safeAcceptExpr(ctx.Expression(), "parenthesized expression")
}

func (g *IRGenerator) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	// PrimaryExpr wraps the actual primary (literals, identifiers, etc.)
	// Delegate to the primary's Accept method
	if ctx.Primary() != nil {
		return ctx.Primary().Accept(g)
	}
	fmt.Fprintf(os.Stderr, "Error: PrimaryExpr has no primary\n")
	return g.builder.CreateInt32("", 0)
}

func (g *IRGenerator) VisitCastExpr(ctx *parser.CastExprContext) interface{} {
	value := g.safeAcceptExpr(ctx.Expression(), "cast expression")
	targetType := g.getType(ctx.Type_())
	return g.builder.CreateCast(value, targetType)
}

func (g *IRGenerator) VisitDerefExpr(ctx *parser.DerefExprContext) interface{} {
	ptr := g.safeAcceptExpr(ctx.Expression(), "dereference expression")
	return g.builder.CreateLoad(ptr)
}

func (g *IRGenerator) VisitAddrOfExpr(ctx *parser.AddrOfExprContext) interface{} {
	// For address-of, we need to handle this differently
	// This is a placeholder - you may need to adjust based on your needs
	operand := g.safeAcceptExpr(ctx.Expression(), "address-of expression")
	return operand
}

func (g *IRGenerator) VisitFieldAccessExpr(ctx *parser.FieldAccessExprContext) interface{} {
	structPtr := g.safeAcceptExpr(ctx.Expression(), "field access expression")
	fieldName := ctx.IDENTIFIER().GetText()
	return g.builder.CreateStructAccess(structPtr, fieldName)
}

func (g *IRGenerator) VisitVectorLiteralExpr(ctx *parser.VectorLiteralExprContext) interface{} {
	// Placeholder - implement based on your vector literal structure
	fmt.Fprintf(os.Stderr, "Warning: Vector literals not yet implemented\n")
	return g.builder.CreateInt32("", 0)
}

func (g *IRGenerator) VisitMapLiteralExpr(ctx *parser.MapLiteralExprContext) interface{} {
	// Placeholder - implement based on your map literal structure
	fmt.Fprintf(os.Stderr, "Warning: Map literals not yet implemented\n")
	return g.builder.CreateInt32("", 0)
}

func (g *IRGenerator) VisitStructLiteralExpr(ctx *parser.StructLiteralExprContext) interface{} {
	// Placeholder - implement based on your struct literal structure
	fmt.Fprintf(os.Stderr, "Warning: Struct literals not yet implemented\n")
	return g.builder.CreateInt32("", 0)
}

func (g *IRGenerator) VisitAllocaExpr(ctx *parser.AllocaExprContext) interface{} {
	allocType := g.getType(ctx.Type_())
	return g.builder.CreateAlloca(allocType)
}

func (g *IRGenerator) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	funcName := ctx.IDENTIFIER().GetText()
	fn, exists := g.functionMap[funcName]
	if !exists {
		fmt.Fprintf(os.Stderr, "Undefined function: %s\n", funcName)
		return g.builder.CreateInt32("", 0) // error fallback
	}

	var args []*builder.Value
	if ctx.ArgumentList() != nil {
		argList := ctx.ArgumentList().(*parser.ArgumentListContext)
		for i, exprCtx := range argList.AllExpression() {
			arg := g.safeAcceptExpr(exprCtx, fmt.Sprintf("call argument %d", i))
			args = append(args, arg)
		}
	}

	return g.builder.CreateCall(fn, args)
}

// Helper function to safely get expression value with nil check
func (g *IRGenerator) safeAcceptExpr(exprCtx parser.IExpressionContext, defaultMsg string) *builder.Value {
	if exprCtx == nil {
		fmt.Fprintf(os.Stderr, "Error: %s - expression context is nil\n", defaultMsg)
		return g.builder.CreateInt32("", 0)
	}
	
	result := exprCtx.Accept(g)
	if result == nil {
		fmt.Fprintf(os.Stderr, "Error: %s - expression returned nil (type was %T)\n", defaultMsg, exprCtx)
		return g.builder.CreateInt32("", 0)
	}
	
	return result.(*builder.Value)
}

// Helper function to get default value for a type
func (g *IRGenerator) getDefaultValue(t types.Type) interface{} {
	switch t {
	case types.Int8, types.Int16, types.Int32:
		return g.builder.CreateInt32("", 0)
	case types.Int64:
		return g.builder.CreateInt64("", 0)
	case types.Float32:
		return g.builder.CreateFloat32("", 0.0)
	case types.Float64:
		return g.builder.CreateFloat64("", 0.0)
	case types.Bool:
		return g.builder.CreateBool("", false)
	case types.String:
		return g.builder.CreateString("", "")
	default:
		return g.builder.CreateInt32("", 0)
	}
}

// Helper function to convert type contexts to IR types
func (g *IRGenerator) getType(typeCtx parser.ITypeContext) types.Type {
	if typeCtx == nil {
		return types.Int32
	}

	// Type assertion to concrete types
	switch t := typeCtx.(type) {
	case *parser.TypeContext:
		// Get the text representation and parse it
		typeText := t.GetText()
		switch typeText {
		case "int32":
			return types.Int32
		case "int64":
			return types.Int64
		case "int8", "byte":
			return types.Int8
		case "int16":
			return types.Int16
		case "float32":
			return types.Float32
		case "float64":
			return types.Float64
		case "bool":
			return types.Bool
		case "string":
			return types.String
		default:
			// Struct type or other custom type
			return types.Int32 // TODO: handle struct types
		}
	default:
		return types.Int32 // default
	}
}

func (g *IRGenerator) Module() *ir.Module {
	return g.builder.Module()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <source.arc>\n", os.Args[0])
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Read source file
	source, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create lexer and parser
	input := antlr.NewInputStream(string(source))
	lexer := parser.NewArcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewArcParser(stream)

	// Parse the program
	tree := p.Program()
	
	fmt.Println("DEBUG: Parse tree created")
	fmt.Printf("DEBUG: Tree type: %T\n", tree)

	// Generate IR
	gen := NewIRGenerator()
	fmt.Println("DEBUG: Starting IR generation...")
	result := tree.Accept(gen) // Use Accept instead of Visit
	fmt.Printf("DEBUG: Accept returned: %v\n", result)

	// Get the module
	module := gen.Module()

	// Print the IR (for debugging)
	fmt.Println("=== Generated IR ===")
	for _, fn := range module.Functions {
		fmt.Printf("\nFunction: %s\n", fn.Name)
		fmt.Printf("Return Type: %v\n", fn.ReturnType)
		fmt.Printf("Parameters: %v\n", fn.ParamTypes)

		for _, block := range fn.Blocks {
			fmt.Printf("\n  Block: %s\n", block.Name)
			fmt.Printf("  Predecessors: %v\n", len(block.Predecessors))
			fmt.Printf("  Successors: %v\n", len(block.Successors))

			for _, inst := range block.Instructions {
				fmt.Printf("    %s\n", inst.String())
			}
		}
	}

	fmt.Println("\n=== IR Generation Complete ===")
	fmt.Println("IR tree is ready for code generation")
}