
/*
Generated by GoGLL. Do not edit
*/

package token

import(
	"fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
	Type       Type
	Lext, Rext int
	Literal    []rune
}

// Type is the token type
type Type int
const(
	Error  Type = iota  // Error 
	EOF  // EOF 
	Type0  // a 
	Type1  // b 
	Type2  // rule1 
)

var TypeToString = []string{ 
	"Error",
	"EOF",
	"Type0",
	"Type1",
	"Type2",
}

var StringToType = map[string] Type { 
	"Error" : Error, 
	"EOF" : EOF, 
	"Type0" : Type0, 
	"Type1" : Type1, 
	"Type2" : Type2, 
}

var TypeToID = []string { 
	"Error", 
	"EOF", 
	"a", 
	"b", 
	"rule1", 
}

func New(t Type, lext, rext int, lit []rune) *Token {
	return &Token{
		Type: t,
		Lext: lext,
		Rext: rext,
		Literal: lit,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s (%d,%d) %s",
		t.TypeID(), t.Lext, t.Rext, string(t.Literal))
}

func (t Type) String() string {
	return TypeToString[t]
}

func (t *Token) TypeID() string {
	return t.Type.ID()
}

func (t Type) ID() string {
	return TypeToID[t]
}

