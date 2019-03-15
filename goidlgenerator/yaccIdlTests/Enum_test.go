package yaccIdlTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/IdlDefinedTypes"
	. "github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEnums(t *testing.T) {
	verbose := false

	createContext := func() *IdlExprContext {
		return NewIdlExprContext()
	}

	t.Run("One Enum Value", func(t *testing.T) {
		data := `enum ABC {AA};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})
	t.Run("Two Enum Value", func(t *testing.T) {
		data := `enum ABC {AA,BB};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})
	t.Run("One Enum Value assigned hex", func(t *testing.T) {
		data := `enum ABC {AA=0x12};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})

	t.Run("One Enum Value assigned int", func(t *testing.T) {
		data := `enum ABC {AA=12};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})
	t.Run("Four Enum Value assigned int", func(t *testing.T) {
		data := `enum ABC {A=1,A=0x2,A=-3,A=0x4};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})

	t.Run("One Enum Value assigned string", func(t *testing.T) {
		data := `enum ABC {A='1'};`
		//data := `"1"`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})

}