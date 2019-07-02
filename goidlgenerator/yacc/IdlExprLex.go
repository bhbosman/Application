package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import (
	"bytes"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/DFA"
	"io"
	"log"
	"strconv"
)

const eof = 0

type IdlExprLex struct {
	InputStream     io.ByteScanner
	CollDfa         []DFA.IDfa
	TokenToIgnore   map[int]int
	Verbose         bool
	VerboseEachChar bool
	Col             int
	Row             int
	idlExprContext  *IdlExprContext
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

func NewIdlExprLex(
	inputStream io.ByteScanner,
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
	}, nil
}
