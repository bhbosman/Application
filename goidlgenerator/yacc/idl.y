%{
package yacc
	//go:generate goyacc -o idl.go -p "IdlExpr" idl.y
	import (
		"github.com/bhbosman/Application/goidlgenerator/interfaces"
		"github.com/bhbosman/Application/goidlgenerator/TempleteTypes"
		)
%}


%token SingleLineComment
%token Identifier
%token Integer_literal
%token Hex_literal
%token Octet_literal
%token Floating_pt_literal
%token Fixed_pt_literal
%token Character_literal
%token Wide_character_literal
%token String_literal
%token Wide_string_literal
%token Scope
%token Shl
%token Shr
%token Whitespace

%token Rwboolean
%token Rwcase
%token Rwchar
%token Rwconst
%token Rwdefault
%token Rwdouble
%token Rwenum
%token RwFALSE
%token Rwfixed
%token Rwfloat
%token Rwlong
%token Rwmodule
%token Rwnative
%token Rwoctet
%token Rwsequence
%token Rwshort
%token Rwstring
%token Rwstruct
%token Rwswitch
%token RwTRUE
%token Rwtypedef
%token Rwunsigned
%token Rwunion
%token Rwvoid
%token Rwwchar
%token Rwwstring
%token RwMitchAlpha
%token RwMitchBitField
%token RwMitchByte
%token RwMitchDate
%token RwMitchTime
%token RwMitchPrice04
%token RwMitchPrice08
%token RwMitchUInt08
%token RwMitchUInt16
%token RwMitchUInt32
%token RwMitchUInt64
%token RwMitchMessageNumberType
%token RwMitchMessageLengthType





%union{
	Identifier 	string
	IntegerValue    int64
	StringValue string
	FloatValue float64
	ConstValue interfaces.IConstantValue
	BoolValue bool

	Member 		*Member
	Declarator 	interfaces.IDeclarator
	DefinedType 	interfaces.IDefinedType
	DefinitionDeclaration interfaces.IDefinitionDeclaration
	Specification []interfaces.IDefinitionDeclaration
}

%type	<StringValue>	String_literal Wide_string_literal
%type	<StringValue>	Character_literal Wide_character_literal
%type	<Identifier>	Identifier simple_declarator native_dcl
%type	<ConstValue>	literal primary_expr

%type 	<FloatValue> 	Floating_pt_literal
%type	<IntegerValue>	positive_int_const
%type	<IntegerValue>	const_expr
%type	<IntegerValue>	or_expr
%type	<IntegerValue>	xor_expr
%type	<IntegerValue>	and_expr
%type	<IntegerValue>	shift_expr

%type	<IntegerValue>	Integer_literal
%type	<IntegerValue>	add_expr
%type	<IntegerValue>	mult_expr
%type	<IntegerValue>	unary_expr

%type	<IntegerValue>	unary_operator

%type	<BoolValue>	boolean_literal

%type	<Member>	member
%type	<DefinedType>	type_spec
%type	<DefinedType>	signed_long_int
%type	<DefinedType>	simple_type_spec
%type	<DefinedType>	floating_pt_type
%type	<DefinedType>	base_type_spec
%type	<DefinedType>	mitch_type_spec
%type	<DefinedType>	signed_short_int
%type	<DefinedType>	char_type
%type	<DefinedType>	wide_char_type
%type	<DefinedType>	integer_type
%type	<DefinedType>	boolean_type
%type	<DefinedType>	octet_type
%type	<DefinedType>	unsigned_short_int
%type	<DefinedType>	unsigned_longlong_int
%type	<DefinedType>	unsigned_long_int
%type	<DefinedType>	signed_int
%type	<DefinedType>	signed_longlong_int
%type	<DefinedType>	unsigned_int
%type	<DefinedType>	scoped_name
%type	<DefinedType>	const_type
%type	<DefinedType>	fixed_pt_const_type
%type	<DefinedType>	sequence_type
%type	<DefinedType>	mitch_bit_field
%type	<DefinedType>	template_type_spec
%type	<DefinedType>	string_type
%type	<DefinedType>	wide_string_type
%type	<DefinedType>	unsigned_short_int
%type	<DefinedType>	fixed_pt_type
%type	<Specification>	specification
%type	<DefinitionDeclaration> definition
%type	<DefinitionDeclaration> definitions
%type	<DefinitionDeclaration> module_dcl
%type	<DefinitionDeclaration> type_dcl
%type	<DefinitionDeclaration> const_dcl
%type	<DefinitionDeclaration> struct_dcl

%type	<DefinitionDeclaration> struct_forward_dcl
%type	<DefinitionDeclaration>	struct_def constr_type_dcl enum_dcl

%type	<DefinitionDeclaration> type_declarator typedef_dcl


%type	<Declarator>	declarator any_declarators enumerator



%%
specification :
	definitions{
		$$ = AddDefinitions($1)
		context, _ := GetIdlExprContext(IdlExprlex)
		context.specification = $$
	}

definitions:
	definition{
		$$ = $1
	}
	|definitions definition{
		$$ = $1
		GetLast($$).SetNext($2)
	}

definition :
	module_dcl ';'{
		$$ = $1
	}
	| const_dcl ';'{
		$$ = $1
	}
	| type_dcl ';'{
		err := AddTypeDclToContext(IdlExprlex, $1)
		if err != nil {
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl
		}

		$$ = $1
	}
module_dcl :
	Rwmodule Identifier '{' definitions '}'{
		$$ = $4
	}
scoped_name :
	Identifier{
		lex, err := GetIdlExprContext(IdlExprlex)
		if err == nil {
			definitionDeclaration := lex.FindScopeName($1)
			if definitionDeclaration == nil{
				IdlExprlex.Error(__yyfmt__.Sprintf("Value %v is not declared", $1))
				return DefNotFound
			}else{
				$$ = definitionDeclaration
			}
		}else{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure. %v", $1))
			return NoLex
		}
	}
//	| Scope Identifier{
//		lex, err := GetIdlExprLex(IdlExprlex)
//		if err == nil && lex != nil {
//			$$ = NewScopeName($2)
//		} else {
//			return NoLex
//		}
//
//	}
//	| scoped_name Scope Identifier{
//		lex, err := GetIdlExprLex(IdlExprlex)
//		if err == nil && lex != nil {
//			$$ = NewScopeName($1.GetName()+$3)
//		} else{
//			return NoLex
//		}
//
//	}
const_dcl :
	Rwconst const_type Identifier '=' const_expr{
		$$ = Newconst_dcl()
	}
const_type :
	integer_type{
		$$ = $1
	}
	| floating_pt_type{
		$$ = $1
	}
	| fixed_pt_const_type{
		$$ = $1
	}
	| char_type{
		$$ = $1
	}
	| wide_char_type{
		$$ = $1
	}
	| boolean_type{
		$$ = $1
	}
	| octet_type{
		$$ = $1
	}
	| string_type{
		$$ = $1
	}
	| wide_string_type{
		$$ = $1
	}
	| scoped_name{
		$$ = $1
	}


const_expr :
	or_expr{
		$$ = $1
	}

or_expr :
	xor_expr{
		$$ = $1
	}
	| or_expr '|' xor_expr
	{
		$$ = $1
	}

xor_expr :
	and_expr {
		$$ = $1
	}
	| xor_expr '^' and_expr{
		$$ = $1
	}

and_expr :
	shift_expr {
		$$ = $1
	}
	| and_expr '&' shift_expr{
		$$ = $1
	}


shift_expr :
	add_expr{
		$$ = $1
	}
	| shift_expr Shr add_expr{
		$$ = $1
	}
	| shift_expr Shl add_expr{
		$$ = $1
	}

add_expr :
	mult_expr {
		$$ = $1
	}
	| add_expr '+' mult_expr{
		$$ = $1 + $3
	}
	| add_expr '-' mult_expr{
		$$ = $1 - $3
	}

mult_expr :
	unary_expr{
		$$ = $1
	}
	| mult_expr '*' unary_expr{
		$$ = $1 * $3
	}
	| mult_expr '/' unary_expr{
		$$ = $1 / $3
	}
	| mult_expr '%' unary_expr{
		$$ = $1 % $3
	}

unary_expr :
	unary_operator primary_expr {
		value, ok := $2.Value().(int64)
		if ok{
			$$ = value
		}else{
			SendError(IdlExprlex, "Value must be an integer (int64)")
			return ErrorMustbeAnInt
		}
	}
	| primary_expr{
		value, ok := $1.Value().(int64)
		if ok{
			$$ = value
		}else{
			SendError(IdlExprlex, "Value must be an integer (int64)")
			return ErrorMustbeAnInt
		}
	}

unary_operator :
	'-' {
		$$ = -1
	}
	| '+'{
		$$ = 0
	}
	| '~'
	{
		$$ = 99
	}

primary_expr :
	scoped_name{
		lex, err := GetIdlExprContext(IdlExprlex)
		if err == nil {
			data := lex.FindScopeName($1.GetName())
			if data == nil{
				IdlExprlex.Error(__yyfmt__.Sprintf("Could not find defined value %v", $1.GetName()))
				return 10003
			}
		} else {
			IdlExprlex.Error("Could not find lex")
			return NoLex
		}
	}
	| literal{
		$$ = $1
	}
	| '(' const_expr ')'{
		$$ = newConstantValueWithNoLength($2, interfaces.Int64)
	}

literal :
	Integer_literal{
		$$ = newConstantValueWithNoLength($1, interfaces.Int64)
	}
	| Floating_pt_literal{
		$$ = nil
	}
	| Fixed_pt_literal{
		$$ = nil
	}
	| Character_literal{
		$$ = newConstantValue($1, interfaces.Char,1)
	}
	| Wide_character_literal{
		$$ = newConstantValue($1, interfaces.WChar, 1)
	}
	| boolean_literal{
		$$ = newConstantValueWithNoLength($1,interfaces.Bool)
	}
	| String_literal{
		$$ = newConstantValueWithNoLength($1, interfaces.String)
	}
	| Wide_string_literal{
		$$ = newConstantValueWithNoLength($1, interfaces.WideString)
	}

boolean_literal :
	RwTRUE{
		$$ = true
	}
	| RwFALSE{
		$$ = false
	}

positive_int_const : const_expr

type_dcl :
	constr_type_dcl{
		$$ = $1
	}
	| native_dcl{
		nativeDecl := NewNativeDeclaration($1)
		$$ = nativeDecl
	}
	| typedef_dcl{
		if $1 == nil{
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl

		}
		if typeDecl, ok := $1.(interfaces.ITypeDeclaration); ok{
			err := AddTypedefDcl(IdlExprlex, typeDecl)
			if err != nil {
				SendError(IdlExprlex, "AddTypedefDcl error")
				return ErrorOnAddTypedefDcl
			}
		}else{
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl
		}

	}

type_spec : simple_type_spec{
	$$ = $1
}

mitch_type_spec:
	RwMitchAlpha '<' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchAlpha, $3)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchByte{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchByte, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchDate{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchDate, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchTime{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchTime, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchPrice04{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchPrice04, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchPrice08{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchPrice08, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchUInt08{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchUInt08, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchUInt16{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchUInt16, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchUInt32{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchUInt32, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchUInt64{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchUInt64, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchMessageNumberType '<' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchMessageNumber, $3)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|RwMitchMessageLengthType '<' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.MitchMessageLength, $3)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}

	}

simple_type_spec :
	mitch_type_spec{
		$$ = $1
	}
	|base_type_spec{
		$$ = $1
	}
	| scoped_name{
		$$ = $1
	}
base_type_spec :
	integer_type{
		$$ = $1
	}
	| floating_pt_type {
		$$ = $1
	}
	| char_type {
		$$ = $1
	}
	| wide_char_type {
		$$ = $1
	}
	| boolean_type {
		$$ = $1
	}
	| octet_type {
		$$ = $1
	}
floating_pt_type :
	Rwfloat {
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Float, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	| Rwdouble{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Double, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
     	}
	| Rwlong Rwdouble{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.LongDouble, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
       	}
integer_type :
	signed_int{
		$$ = $1
	}
	| unsigned_int{
		$$ = $1
	}
signed_int :
	signed_short_int{
		$$ = $1
	}
	| signed_long_int{
		$$ = $1
	}
	| signed_longlong_int{
		$$ = $1
	}
signed_short_int :
	Rwshort{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Int16, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
signed_long_int :
	Rwlong{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Int32, 0)
		if $$ == nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
signed_longlong_int :
	Rwlong Rwlong{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Int64, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

unsigned_int :
	unsigned_short_int{
		$$ = $1
	}
	| unsigned_long_int{
		$$ = $1
	}
	| unsigned_longlong_int{
		$$ = $1
	}

unsigned_short_int :
	Rwunsigned Rwshort{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Uint16, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
unsigned_long_int :
	Rwunsigned Rwlong{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Uint32, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
unsigned_longlong_int :
	Rwunsigned Rwlong Rwlong{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Uint64, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

char_type :
	Rwchar{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Char, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

wide_char_type :
	Rwwchar{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.WChar, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
boolean_type :
	Rwboolean{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Bool, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

octet_type :
	Rwoctet{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Octet, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

template_type_spec :
	sequence_type{
		$$ = $1
	}
	| string_type{
		$$ = $1
	}
	| wide_string_type{
		$$ = $1
	}
	| fixed_pt_type{
		$$ = $1
	}
	|mitch_bit_field{
		$$ = $1
	}

mitch_bit_field:
	RwMitchBitField '<' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}

		$$, err = context.IDlBaseType.CreateType(interfaces.MitchBitField, []string{$3, $5, $7, $9, $11, $13, $15, $17})
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}



sequence_type :
	Rwsequence '<' type_spec ',' positive_int_const '>'{
		$$ = TempleteTypes.NewSequenceType($3, $5)
	}
	| Rwsequence '<' type_spec '>'{
		$$ = TempleteTypes.NewSequenceType($3, -1)
	}

string_type :
	Rwstring '<' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.String, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	|
	Rwstring{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.String, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

wide_string_type :
	Rwwstring '<' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.WideString, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}
	| Rwwstring{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.WideString, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}


fixed_pt_type :
	Rwfixed '<' positive_int_const ',' positive_int_const '>'{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Fixed, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

fixed_pt_const_type :
	Rwfixed{
		context, err := GetIdlExprLex(IdlExprlex)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure"))
			return NoLex
		}
		$$, err = context.IDlBaseType.CreateType(interfaces.Fixed, 0)
		if err != nil{
			IdlExprlex.Error(__yyfmt__.Sprintf("%v", err))
			return TypeNotAvailable
		}
	}

constr_type_dcl :
	struct_dcl{
		$$ = $1
	}
	| union_dcl{
		$$ = Newunion_dcl()
	}
	| enum_dcl{
		$$ = $1
	}

/*
 * (45) <struct_dcl> ::= <struct_def> | <struct_forward_dcl>
 */
struct_dcl :
	struct_def{
		$$ = $1

    	}
    	| struct_forward_dcl{
    		$$ = $1
    	}

/* (46) <struct_def> ::= 'struct' <Identifier> '{' <member>+ '}'*/
struct_def :
	Rwstruct Identifier '{' member '}'{
		def := NewStructDefinition($2)
		member := $4
		for member != nil {
			decl := member.Declarator
			for decl != nil  {
				def.AddMember(member.DefinedType, decl)
			decl = decl.Next()
			}
			member = member.Next
		}
		$$ = def
	}
	|Rwstruct Identifier '{'  '}'{
		def := NewStructDefinition($2)
		$$ = def
	}

///*(47) <member> ::= <type_spec> <declarators> ';'*/
member :
	type_spec declarator ';'{
 		$$ = NewMember($1, $2, nil)
	}
	|member member{
		$$ = NewMember($1.DefinedType, $1.Declarator, $2)
	}



struct_forward_dcl :
	Rwstruct Identifier{
		$$ = Newstruct_forward_dcl()
	}
union_dcl : union_def | union_foRward_dcl
union_def : Rwunion Identifier Rwswitch '(' switch_type_spec ')' '{' switch_body '}'
switch_type_spec : integer_type | char_type | boolean_type | scoped_name

switch_body : cases
cases: case |case case
case : case_labels element_spec ';'
case_labels : case_label|case_label case_label
case_label : Rwcase const_expr ':' | Rwdefault ':'
element_spec : type_spec declarator
union_foRward_dcl : Rwunion Identifier
enum_dcl :
	Rwenum Identifier '{' enumerator  '}'{
		def := NewEnumDcl($2)
		decl := $4
		for decl != nil  {
			def.AddMember(decl)
			decl = decl.Next()
		}
		$$ = def
	}
enumerator :
	declarator{}
	| enumerator ',' enumerator{
		$1.SetNext($3)
	}


array_declarator :
	Identifier fixed_array_sizes{

	}

fixed_array_sizes:
	fixed_array_size{

	}
	| fixed_array_sizes fixed_array_size{

	}

fixed_array_size :
	'[' positive_int_const ']'{
	}

native_dcl :
	Rwnative simple_declarator{
		$$ = $2
	}

simple_declarator :
	Identifier{
		$$ =$1
	}

typedef_dcl :
	Rwtypedef type_declarator{
		$$ = $2
	}


/*
(64) <type_declarator> ::= { <simple_type_spec> | <template_type_spec> | <constr_type_dcl> } <any_declarators>
*/
type_declarator :
	simple_type_spec any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}
	| template_type_spec any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}
	| constr_type_dcl  any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}

/*
// (65) <any_declarators> ::= <any_declarator> { ',' <any_declarator> }*
// (66) <any_declarator> ::= <simple_declarator> | <array_declarator>
// */
any_declarators:
	simple_declarator{
		$$ = NewDeclarator($1, nil)
	}
	|array_declarator{

	}
	|any_declarators ',' any_declarators{
		$1.SetNext($3)
		$$ = $1
	}
/*
// (67) <declarators> ::= <declarator> { ',' <declarator> }*
// (68) <declarator> ::= <simple_declarator>
// */
declarator :
	simple_declarator '=' literal {
		$$ = NewDeclarator($1, $3)
	}
	|simple_declarator{
		$$ = NewDeclarator($1, nil)
	}
	|declarator ',' declarator{
		$1.SetNext($3)
		$$ = $1
	}
%%


