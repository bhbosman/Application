%{
package yacc
	//go:generate goyacc -o idl.go -p "IdlExpr" idl.y
	import (
		"github.com/bhbosman/gofintech/goidlgenerator/interfaces"
		"github.com/bhbosman/gofintech/goidlgenerator/DefinedTypes"
		"github.com/bhbosman/gofintech/goidlgenerator/TempleteTypes")
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


%union{
	Identifier 	string
	GenericValue	interface{}
	Member 		*Member
	Declarator 	interfaces.IDeclarator
	DefinedType 	interfaces.IDefinedType
	DefinitionDeclaration interfaces.IDefinitionDeclaration
	ScopeName interfaces.IScopeName
	Specification []interfaces.IDefinitionDeclaration
}


%type	<Identifier>	Identifier
%type	<Identifier>	simple_declarator
%type	<Identifier> native_dcl

%type	<GenericValue>	positive_int_const
%type	<GenericValue>	const_expr
%type	<GenericValue>	or_expr
%type	<GenericValue>	xor_expr
%type	<GenericValue>	and_expr
%type	<GenericValue>	shift_expr
%type	<GenericValue>	literal
%type	<GenericValue>	Integer_literal
%type	<GenericValue>	add_expr
%type	<GenericValue>	mult_expr
%type	<GenericValue>	unary_expr
%type	<GenericValue>	primary_expr
%type	<GenericValue>	unary_operator




%type	<Member>	member
%type	<DefinedType>	type_spec
%type	<DefinedType>	signed_long_int
%type	<DefinedType>	simple_type_spec
%type	<DefinedType>	floating_pt_type
%type	<DefinedType>	base_type_spec
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
%type	<ScopeName>	scoped_name
%type	<DefinedType>	const_type
%type	<DefinedType>	fixed_pt_const_type
%type	<DefinedType>	sequence_type
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
%type	<DefinitionDeclaration>	struct_def constr_type_dcl, enum_dcl

%type	<DefinitionDeclaration> type_declarator typedef_dcl


%type	<Declarator>	declarator any_declarators


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
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure", $1))
			return NoLex
		}
	}
	| Scope Identifier{
		lex, err := GetIdlExprLex(IdlExprlex)
		if err == nil && lex != nil {
			$$ = NewScopeName($2)
		} else {
			return NoLex
		}

	}
	| scoped_name Scope Identifier{
		lex, err := GetIdlExprLex(IdlExprlex)
		if err == nil && lex != nil {
			$$ = NewScopeName($1.GetName()+$3)
		} else{
			return NoLex
		}

	}
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
	or_expr

or_expr :
	xor_expr
	| or_expr '|' xor_expr

xor_expr : and_expr | xor_expr '^' and_expr

and_expr : shift_expr | and_expr '&' shift_expr

shift_expr : add_expr | shift_expr Shr add_expr | shift_expr Shl add_expr

add_expr : mult_expr | add_expr '+' mult_expr | add_expr '-' mult_expr

mult_expr : unary_expr | mult_expr '*' unary_expr | mult_expr '/' unary_expr | mult_expr '%' unary_expr

unary_expr :
	unary_operator primary_expr {
		$$ = $2
	}
	| primary_expr{
		$$=$1
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
		$$ = 0
	}

literal :
	Integer_literal{
		$$ = $1
	}
	| Floating_pt_literal{
		$$ = nil
	}
	| Fixed_pt_literal{
		$$ = nil
	}
	| Character_literal{
		$$ = nil
	}
	| Wide_character_literal{
		$$ = nil
	}
	| boolean_literal{
		$$ = nil
	}
	| String_literal{
		$$ = nil
	}
	| Wide_string_literal{
		$$ = nil
	}

boolean_literal :
	RwTRUE
	| RwFALSE

positive_int_const : const_expr

type_dcl :
	constr_type_dcl{
		$$ = $1
	}
	| native_dcl{
		nativeDecl := NewNativeDeclaration($1)
		err := AddNativeDcl(IdlExprlex, nativeDecl)
		if err != nil {
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl
		}
		$$ = nativeDecl
	}
	| typedef_dcl{

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
simple_type_spec :
	base_type_spec{
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
		$$ = DefinedTypes.NewFloatType()
	}
	| Rwdouble{
        	$$ = DefinedTypes.NewDoubleType()
     	}
	| Rwlong Rwdouble{
               	$$ = DefinedTypes.NewLongDoubleType()
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
		$$ = DefinedTypes.NewSignedShortType()
	}
signed_long_int :
	Rwlong{
		$$ = DefinedTypes.NewSignedLongType()
	}
signed_longlong_int :
	Rwlong Rwlong{
		$$ = DefinedTypes.NewSignedLongLongType()
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
		$$ = DefinedTypes.NewUnSignedShortType()
	}
unsigned_long_int :
	Rwunsigned Rwlong{
		$$ = DefinedTypes.NewUnsignedLongType()
	}
unsigned_longlong_int :
	Rwunsigned Rwlong Rwlong{
		$$ = DefinedTypes.NewUnsignedLongLongType()
	}

char_type :
	Rwchar{
		$$ = DefinedTypes.NewCharType()
	}

wide_char_type :
	Rwwchar{
		$$ = DefinedTypes.NewWideCharType()
	}
boolean_type :
	Rwboolean{
		$$ = DefinedTypes.NewBooleanType()
	}

octet_type :
	Rwoctet{
		$$ = DefinedTypes.NewOctetType()
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

sequence_type :
	Rwsequence '<' type_spec ',' positive_int_const '>'{
		switch  val := $5.(type){
			case int:
				$$ = TempleteTypes.NewSequenceType($3, int64(val))
				break
			case int64:
				$$ = TempleteTypes.NewSequenceType($3, val)
				break
			default:
				IdlExprlex.Error(__yyfmt__.Sprintf("Unexpected type. wanted int or int64. Got %v", val))
				return 10001
		}
	}
	| Rwsequence '<' type_spec '>'{
		$$ = TempleteTypes.NewSequenceType($3, -1)
	}

string_type :
	Rwstring '<' positive_int_const '>'{
		val, ok := $3.(int64)
		if ok{
			$$ = DefinedTypes.NewStringType(val)
		}
	}
	|
	Rwstring{
		$$ = DefinedTypes.NewStringType(-1)
	}

wide_string_type :
	Rwwstring '<' positive_int_const '>'{
		val, ok := $3.(int64)
		if ok{
			$$ = DefinedTypes.NewWideStringType(val)
		} else {
			return 9898
		}
	}
	| Rwwstring{
		$$ = DefinedTypes.NewWideStringType(-1)
	}


fixed_pt_type :
	Rwfixed '<' positive_int_const ',' positive_int_const '>'{
		$$ = DefinedTypes.NewFixedType()
	}

fixed_pt_const_type :
	Rwfixed{
		$$ = DefinedTypes.NewFixedType()
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
		$$ = Newenum_dcl()
	}
enumerator : declarator| enumerator ',' enumerator

enumerator_value: positive_int_const
array_declarator : Identifier fixed_array_sizes
fixed_array_sizes: fixed_array_size | fixed_array_sizes fixed_array_size
fixed_array_size : '[' positive_int_const ']'
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
	if $1 == nil{
		SendError(IdlExprlex, "simple_type_spec can not be nil")
		return ErrorTypeisNull
	}
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
	simple_declarator '=' positive_int_const {
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


