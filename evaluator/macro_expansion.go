package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitons := []int{}

	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitons = append(definitons, i)
		}

		for i := len(definitons) - 1; i >= 0; i = i - 1 {
			definitionIndex := definitons[i]
			program.Statements = append(program.Statements[:definitionIndex], program.Statements[definitionIndex+1:]...)
		}
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	return true
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env: env,
		Body: macroLiteral.Body,
	}

	env.Set(letStatement.Name.Value, macro)
}
