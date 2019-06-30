package yacc

import (
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/DFA"
)

func CreateIdlTokens() ([]DFA.IDfa, error) {
	reservedWords := make(map[string]int)
	reservedWords["enum"] = Rwenum
	reservedWords["struct"] = Rwstruct
	reservedWords["typedef"] = Rwtypedef
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
	//reservedWords["MitchMessageNumberType"] = RwMitchMessageNumberType
	//reservedWords["MitchMessageLengthType"] = RwMitchMessageLengthType

	collDfaFunctions := []func() (DFA.IDfa, error){
		func() (DFA.IDfa, error) { return DFA.NewIdentifier(Identifier, reservedWords) },
		func() (DFA.IDfa, error) { return DFA.NewDfaInteger(Integer_literal) },
		func() (DFA.IDfa, error) { return DFA.NewHexDfa(Hex_literal) },
		func() (DFA.IDfa, error) { return DFA.NewDfaWhiteSpace(Whitespace) },
		func() (DFA.IDfa, error) { return DFA.NewSingleLineComment(SingleLineComment) },
		//func() (DFA.IDfa, error) { return DFA.NewStringNode(String_literal), nil },
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

