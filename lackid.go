package lackid

import (
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

var Analyzer = &gqlanalysis.Analyzer{
	Name: "lackid",
	Doc:  "lackid finds a selection for a type which has id field but the selection does not have id",
	Run:  run,
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	for _, q := range pass.Queries {
		checkOperations(pass, q.Operations)
		checkFragments(pass, q.Fragments)
	}
	return nil, nil
}

func checkOperations(pass *gqlanalysis.Pass, ops ast.OperationList) {
	for _, op := range ops {
		for _, sel := range op.SelectionSet {
			checkField(pass, sel)
			checkInlineFragment(pass, sel)
		}
	}
}

func checkFragments(pass *gqlanalysis.Pass, fs ast.FragmentDefinitionList) {
	for _, f := range fs {
		checkFragment(pass, f)
	}
}

func checkFragment(pass *gqlanalysis.Pass, f *ast.FragmentDefinition) {
	if !hasID(f.Definition.Fields) {
		return
	}

	var hasID bool
	for _, s := range f.SelectionSet {
		hasID = hasID || isID(s)
		checkField(pass, s)
	}

	if hasID {
		return
	}

	name := f.Definition.Name
	if name == "" {
		name = "fragment"
	}

	pass.Reportf(f.Position, "type %s has id field but %s does not have id field", name, f.Name)
}

func checkField(pass *gqlanalysis.Pass, sel ast.Selection) {
	field, _ := sel.(*ast.Field)
	if field == nil || allFragmentSpread(field.SelectionSet) {
		return
	}

	for _, s := range field.SelectionSet {
		checkField(pass, s)
		checkInlineFragment(pass, s)
	}

	ft := pass.Schema.Types[field.Definition.Type.Name()]
	if ft == nil || !hasID(ft.Fields) {
		return
	}

	var hasID bool
	for _, s := range field.SelectionSet {
		hasID = hasID || isID(s)
	}

	if !hasID && !allInlineFragment(field.SelectionSet) {
		pass.Reportf(field.Position, "type %s has id field but selection %s does not have id field", ft.Name, field.Name)
	}
}

func hasID(fields ast.FieldList) bool {
	for _, field := range fields {
		if field.Name == "id" {
			return true
		}
	}
	return false
}

func isID(sel ast.Selection) bool {
	field, _ := sel.(*ast.Field)
	return field != nil && field.Name == "id"
}

func checkInlineFragment(pass *gqlanalysis.Pass, sel ast.Selection) {
	f, _ := sel.(*ast.InlineFragment)
	if f == nil || allFragmentSpread(f.SelectionSet) {
		return
	}

	for _, s := range f.SelectionSet {
		checkField(pass, s)
		checkInlineFragment(pass, s)
	}

	if !hasID(f.ObjectDefinition.Fields) {
		return
	}

	var hasID bool
	for _, s := range f.SelectionSet {
		hasID = hasID || isID(s)
	}

	if !hasID && !allInlineFragment(f.SelectionSet) {
		pass.Reportf(f.Position, "type %s has id field but fragment does not have id field", f.ObjectDefinition.Name)
	}
}

func allFragmentSpread(set ast.SelectionSet) bool {
	for _, sel := range set {
		if _, ok := sel.(*ast.FragmentSpread); !ok {
			return false
		}
	}
	return true
}

func allInlineFragment(set ast.SelectionSet) bool {
	for _, sel := range set {
		if _, ok := sel.(*ast.InlineFragment); !ok {
			return false
		}
	}
	return true
}
