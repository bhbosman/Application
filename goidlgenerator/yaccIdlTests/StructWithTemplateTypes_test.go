package yaccIdlTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/IdlDefinedTypes"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/json"
	"github.com/bhbosman/Application/goidlgenerator/TempleteTypes"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
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
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		assert.IsType(t, &IdlDefinedTypes.StringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "a", typeDeclaration.GetDeclarator().Identifier())
	})
	t.Run("TypeDefStringTwoIdentifier", func(t *testing.T) {
		data := `typedef string a, b;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		assert.IsType(t, &IdlDefinedTypes.StringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "a", typeDeclaration.GetDeclarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &IdlDefinedTypes.StringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "b", typeDeclaration.GetDeclarator().Identifier())
	})
	t.Run("TypeDefWideStringOneIdentifier", func(t *testing.T) {
		data := `typedef wstring a;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		assert.IsType(t, &IdlDefinedTypes.WideStringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "a", typeDeclaration.GetDeclarator().Identifier())
	})
	t.Run("TypeDefWideStringTwoIdentifier", func(t *testing.T) {
		data := `
			// some comment
			typedef wstring a, b;
			`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		assert.IsType(t, &IdlDefinedTypes.WideStringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "a", typeDeclaration.GetDeclarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &IdlDefinedTypes.WideStringType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "b", typeDeclaration.GetDeclarator().Identifier())
	})
	t.Run("TypeDefLongTwoIdentifier", func(t *testing.T) {
		data := `typedef long a, b;`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        verbose})
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
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "a", typeDeclaration.GetDeclarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[1])
		typeDeclaration = DeclaredTypes[1].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "b", typeDeclaration.GetDeclarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[2])
		typeDeclaration = DeclaredTypes[2].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "c", typeDeclaration.GetDeclarator().Identifier())

		assert.IsType(t, &yacc.TypeDeclarator{}, DeclaredTypes[3])
		typeDeclaration = DeclaredTypes[3].(*yacc.TypeDeclarator)
		assert.IsType(t, &TempleteTypes.SequenceType{}, typeDeclaration.GetDefinedTyped())
		assert.Equal(t, "d", typeDeclaration.GetDeclarator().Identifier())
	})
}
