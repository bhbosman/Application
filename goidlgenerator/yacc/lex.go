package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/DFA"
	"github.com/bhbosman/Application/goidlgenerator/IdlDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/MitchDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
	"log"
	"strconv"
)

const eof = 0

func StringToIDlBaseType(s string) (interfaces.IBaseTypeInformation, error) {
	switch s {
	case "IdlNative":
		return IdlDefinedTypes.NewIdlNativeTypeInformation(), nil
	case "Mitch":
		return MitchDefinedTypes.NewMitchTypeInformation(), nil
	}

	return nil, errors.New("invalid type")

}

type IdlExprLex struct {
	InputStream     io.ByteScanner
	CollDfa         []DFA.IDfa
	TokenToIgnore   map[int]int
	Verbose         bool
	VerboseEachChar bool
	Col             int
	Row             int
	idlExprContext  *IdlExprContext
	IDlBaseType     interfaces.IBaseTypeInformation
}

func (x *IdlExprLex) Lex(yylval *IdlExprSymType) int {
	for {
		collection := DFA.NewDfaCollection(x.CollDfa, x.Row, x.Col)
		token, lexValue, err := func() (int, string, error) {
			byteBuffer := bytes.NewBuffer(nil)
			doFinish := func() (int, string, error) {
				lexem := byteBuffer.String()
				return collection.Token(lexem)
			}
			for {
				byteRead, err := x.InputStream.ReadByte()
				if err != nil {
					if err == io.EOF {
						i, s, e := doFinish()
						if e != nil {
							return eof, "", e
						}
						return i, s, e
					}
					return eof, "", err
				}
				if x.VerboseEachChar {
					fmt.Println(string(byteRead), x.Row, x.Col)
				}

				b, err := collection.Walk(byteRead)
				if err != nil {

				}
				if !b {
					err := x.InputStream.UnreadByte()
					if err != nil {
						return eof, "", err
					}
					return doFinish()
				}
				byteBuffer.WriteByte(byteRead)
				if byteRead == '\n' {
					x.Row++
					x.Col = 0
				} else {
					x.Col++
				}
			}
		}()
		if err != nil {
			return 0
		}
		if x.Verbose {
			idx := token - IdlExprPrivate + 2

			fmt.Println(IdlExprTokname(idx), collection.Row(), collection.Col(), token, lexValue)
		}

		if _, ok := x.TokenToIgnore[token]; ok {
			continue
		}

		switch token {
		case String_literal:
			yylval.StringValue = lexValue
		case Character_literal:
			yylval.StringValue = lexValue

		case Identifier:
			yylval.Identifier = lexValue
			break
		case Integer_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 10, 64)
			yylval.IntegerValue = integerValue
			break
		case Hex_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 0, 64)
			yylval.IntegerValue = integerValue
			token = Integer_literal
			break
		}
		return token
	}
}

func (x *IdlExprLex) Error(s string) {
	log.Printf("parse error: %s at (%d,%d).", s, x.Row, x.Col)
}

func CreateIdlTokens() ([]DFA.IDfa, error) {
	reservedWords := make(map[string]int)
	reservedWords["boolean"] = Rwboolean
	reservedWords["case"] = Rwcase
	reservedWords["char"] = Rwchar
	reservedWords["const"] = Rwconst
	reservedWords["default"] = Rwdefault
	reservedWords["double"] = Rwdouble
	reservedWords["enum"] = Rwenum
	reservedWords["FALSE"] = RwFALSE
	reservedWords["fixed"] = Rwfixed
	reservedWords["float"] = Rwfloat
	reservedWords["long"] = Rwlong
	reservedWords["module"] = Rwmodule
	reservedWords["native"] = Rwnative
	reservedWords["octet"] = Rwoctet
	reservedWords["sequence"] = Rwsequence
	reservedWords["short"] = Rwshort
	reservedWords["string"] = Rwstring
	reservedWords["struct"] = Rwstruct
	reservedWords["switch"] = Rwswitch
	reservedWords["TRUE"] = RwTRUE
	reservedWords["typedef"] = Rwtypedef
	reservedWords["unsigned"] = Rwunsigned
	reservedWords["union"] = Rwunion
	reservedWords["void"] = Rwvoid
	reservedWords["wchar"] = Rwwchar
	reservedWords["wstring"] = Rwwstring
	reservedWords["MitchAlpha"] = RwMitchAlpha
	reservedWords["MitchBitField"] = RwMitchBitField
	reservedWords["MitchByte"] = RwMitchByte
	reservedWords["MitchDate"] = RwMitchDate
	reservedWords["MitchTime"] = RwMitchTime
	reservedWords["MitchPrice04"] = RwMitchPrice04
	reservedWords["MitchPrice08"] = RwMitchPrice08
	reservedWords["MitchUInt08"] = RwMitchUInt08
	reservedWords["MitchUInt16"] = RwMitchUInt16
	reservedWords["MitchUInt32"] = RwMitchUInt32
	reservedWords["MitchUInt64"] = RwMitchUInt64
	reservedWords["MitchMessageNumberType"] = RwMitchMessageNumberType
	reservedWords["MitchMessageLengthType"] = RwMitchMessageLengthType

	collDfaFunctions := []func() (DFA.IDfa, error){
		func() (DFA.IDfa, error) { return DFA.NewIdentifier(Identifier, reservedWords) },
		func() (DFA.IDfa, error) { return DFA.NewDfaInteger(Integer_literal) },
		func() (DFA.IDfa, error) { return DFA.NewHexDfa(Hex_literal) },
		func() (DFA.IDfa, error) { return DFA.NewDfaWhiteSpace(Whitespace), nil },
		func() (DFA.IDfa, error) { return DFA.NewSingleLineComment(SingleLineComment) },
		func() (DFA.IDfa, error) { return DFA.NewStringNode(String_literal), nil },
		func() (DFA.IDfa, error) { return DFA.NewCharNode(Character_literal) },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("{", '{', '{') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("}", '}', '}') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("<", '<', '<') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken(">", '>', '>') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("[", '[', '[') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("]", ']', ']') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken("=", '=', '=') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken(";", ';', ';') },
		func() (DFA.IDfa, error) { return DFA.NewGenericSingleCharToken(",", ',', ',') },
	}

	collDfa := make([]DFA.IDfa, len(collDfaFunctions), len(collDfaFunctions))
	err := Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		for i, f := range collDfaFunctions {
			dfa, err := f()
			if err != nil {
				errorList.Add(err)
			}
			collDfa[i] = dfa
		}
	})
	if err != nil {
		return nil, err
	}

	return collDfa, nil
}

type NewIdlExprLexParams struct {
	IdlExprContext *IdlExprContext
	Verbose        bool
}

func NewIdlExprLex(
	inputStream io.ByteScanner,
	IDlBaseType interfaces.IBaseTypeInformation,
	params NewIdlExprLexParams) (*IdlExprLex, error) {

	IdlExprErrorVerbose = true
	tokenToIgnore := []int{
		Whitespace,
		SingleLineComment,
	}
	tokenToIgnoreMap := make(map[int]int)
	for _, token := range tokenToIgnore {
		tokenToIgnoreMap[token] = token
	}

	createIdlTokens, err := CreateIdlTokens()
	if err != nil {
		return nil, err
	}

	return &IdlExprLex{
		InputStream:     inputStream,
		CollDfa:         createIdlTokens,
		TokenToIgnore:   tokenToIgnoreMap,
		Verbose:         params.Verbose,
		VerboseEachChar: false,
		Col:             1,
		Row:             1,
		idlExprContext:  params.IdlExprContext,
		IDlBaseType:     IDlBaseType,
	}, nil

}
