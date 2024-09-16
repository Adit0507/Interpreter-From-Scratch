package evaluator

import (
	"monkey/object"
	"monkey/ast")

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}