package yaccTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/DefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/json"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStructWithTypes(t *testing.T) {

	verbose := false

	t.Run("WithBooleanType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			boolean a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.BooleanType{}, member.DefinedType)
	})

	t.Run("WithBooleanTypeWithIntAssignment", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			boolean a = 1234;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}

		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.BooleanType{}, member.DefinedType)
	})

	t.Run("WithBooleanTypeWithHexAssignment", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			boolean a = 0x1234;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.BooleanType{}, member.DefinedType)
	})

	t.Run("WithCharType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			char a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.NotNil(t, member.DefinedType)
		s := member.DefinedType.GetName()
		assert.Equal(t, "char", s)

		assert.IsType(t, &DefinedTypes.CharType{}, member.DefinedType)
	})
	t.Run("WithDoubleType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			double a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.DoubleType{}, member.DefinedType)
	})
	t.Run("WithFloatType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			float a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.FloatType{}, member.DefinedType)
	})
	t.Run("WithLongDoubleType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long double a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.LongDoubleType{}, member.DefinedType)
	})
	t.Run("WithOctetType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			octet a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.OctetType{}, member.DefinedType)
	})
	t.Run("WithSignedLongLongType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long long a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.SignedLongLongType{}, member.DefinedType)
	})
	t.Run("WithSignedLongType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.SignedLongType{}, member.DefinedType)
	})
	t.Run("WithSignedShortType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			short a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.SignedShortType{}, member.DefinedType)
	})
	t.Run("WithUnsignedLongLongType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			unsigned long long a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.UnsignedLongLongType{}, member.DefinedType)
	})
	t.Run("WithUnsignedLongType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			unsigned long a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.UnsignedLongType{}, member.DefinedType)
	})
	t.Run("WithUnsignedShortType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			unsigned short a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.IsType(t, &DefinedTypes.UnSignedShortType{}, member.DefinedType)
	})
	t.Run("WithWideCharType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			char a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
		assert.IsType(t, &yacc.StructDefinition{}, DeclaredTypes[0])
		structType := DeclaredTypes[0].(*yacc.StructDefinition)
		assert.Len(t, structType.Members, 1)
		member := structType.Members[0]
		assert.Equal(t, "a", member.Declarator.Identifier())
		assert.NotNil(t, member.DefinedType)
		assert.IsType(t, &DefinedTypes.CharType{}, member.DefinedType)
	})

	t.Run("02", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long a,b,c,d,e,f,g,h,i;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))

		jsonPublish, err := Publish.HasOutputType(Publish.Json)
		assert.NoError(t, err)
		assert.NotNil(t, jsonPublish)
	})

	t.Run("03", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long a,aa,aaa;
			long b, bb,bbb;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))

		jsonPublish, err := Publish.HasOutputType(Publish.Json)
		assert.NoError(t, err)
		assert.NotNil(t, jsonPublish)
	})

	t.Run("04", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long a,aa,aaa;
			long b, bb,bbb;
		};

		struct HelloWorld
		{
			long a,aa,aaa;
			long b, bb,bbb;
		};
	`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		assert.NotNil(t, DeclaredTypes)
		assert.Equal(t, 2, len(DeclaredTypes))

		jsonPublish, err := Publish.HasOutputType(Publish.Json)
		assert.NoError(t, err)
		assert.NotNil(t, jsonPublish)

	})
}
