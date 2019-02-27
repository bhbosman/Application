package yaccTests

import (
	"bufio"
	"github.com/bhbosman/gofintech/goidlgenerator/DefinedTypes"
	_ "github.com/bhbosman/gofintech/goidlgenerator/Publish/json"
	"github.com/bhbosman/gofintech/goidlgenerator/TempleteTypes"
	"github.com/bhbosman/gofintech/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTypeDefWithTemplateTypes(t *testing.T) {
	verbose := false

	t.Run("TypeDefStringOneIdentifier", func(t *testing.T) {
		data := `typedef string a;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[0])
		typeDeclaration := DeclaredTypes[0].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.StringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "a", typeDeclaration.Declarator().Identifier())
	})
	t.Run("TypeDefStringTwoIdentifier", func(t *testing.T) {
		data := `typedef string a, b;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 2)

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[0])
		typeDeclaration := DeclaredTypes[0].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.StringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "a", typeDeclaration.Declarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.StringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "b", typeDeclaration.Declarator().Identifier())
	})
	t.Run("TypeDefWideStringOneIdentifier", func(t *testing.T) {
		data := `typedef wstring a;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[0])
		typeDeclaration := DeclaredTypes[0].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.WideStringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "a", typeDeclaration.Declarator().Identifier())
	})
	t.Run("TypeDefWideStringTwoIdentifier", func(t *testing.T) {
		data := `
			// some comment
			typedef wstring a, b;
			`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 2)

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[0])
		typeDeclaration := DeclaredTypes[0].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.WideStringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "a", typeDeclaration.Declarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &DefinedTypes.WideStringType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "b", typeDeclaration.Declarator().Identifier())
	})
	t.Run("TypeDefLongTwoIdentifier", func(t *testing.T) {
		data := `typedef long a, b;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 2)
	})

	t.Run("TypeDefSequence<Long>TwoIdentifier", func(t *testing.T) {
		data := `
			typedef sequence<long> a, b;
			typedef sequence<long> c, d;
		`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex := yacc.NewIdlExprLex(
			reader, IdlExprContext,
			verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 4)

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[0])
		typeDeclaration := DeclaredTypes[0].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "a", typeDeclaration.Declarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "b", typeDeclaration.Declarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[2])
		typeDeclaration = DeclaredTypes[2].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "c", typeDeclaration.Declarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[3])
		typeDeclaration = DeclaredTypes[3].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.DefinedTyped())
		assert.Equal(t, "d", typeDeclaration.Declarator().Identifier())
	})
}
