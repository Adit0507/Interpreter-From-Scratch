package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var res object.Object

	for _, statment := range stmts {
		res = Eval(statment)
	}

	return res
}
