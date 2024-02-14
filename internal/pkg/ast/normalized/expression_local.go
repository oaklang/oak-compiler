package normalized

import (
	"fmt"
	"nar-compiler/internal/pkg/ast"
	"nar-compiler/internal/pkg/ast/typed"
	"nar-compiler/internal/pkg/common"
)

type Local struct {
	*expressionBase
	name   ast.Identifier
	target Pattern
}

func NewLocal(loc ast.Location, name ast.Identifier, target Pattern) Expression {
	return &Local{
		expressionBase: newExpressionBase(loc),
		name:           name,
		target:         target,
	}
}

func (e *Local) flattenLambdas(parentName ast.Identifier, m *Module, locals map[ast.Identifier]Pattern) Expression {
	if lp, ok := locals[e.name]; ok {
		e.target = lp
	}
	return e
}

func (e *Local) replaceLocals(replace map[ast.Identifier]Expression) Expression {
	if r, ok := replace[e.name]; ok {
		return r
	}
	return e
}

func (e *Local) extractUsedLocalsSet(definedLocals map[ast.Identifier]Pattern, usedLocals map[ast.Identifier]struct{}) {
	if _, ok := definedLocals[e.name]; ok {
		usedLocals[e.name] = struct{}{}
	}
}

func (e *Local) annotate(ctx *typed.SolvingContext, typeParams typeParamsMap, modules map[ast.QualifiedIdentifier]*Module, typedModules map[ast.QualifiedIdentifier]*typed.Module, moduleName ast.QualifiedIdentifier, stack []*typed.Definition) (typed.Expression, error) {
	if e.target == nil {
		return nil, common.Error{
			Location: e.location, Message: fmt.Sprintf("local variable `%s` not resolved", e.name),
		}
	}
	return e.setSuccessor(typed.NewLocal(ctx, e.location, e.name, e.target.Successor()))
}
