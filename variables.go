package builder

import (
	"fmt"
)

// Define the core types - the most commonly used types.
const (
	String  coreType = "String"
	Integer coreType = "Integer"
	Int     coreType = "int"
	Bool    coreType = "boolean"
	Boolean coreType = "Boolean"
)

// Type defines a type of a variable.
// Type's single method returns the type.
type Type interface {
	Type() string
}

// Custom allows custom (usually object) types to be defined.
func Custom(c string) Type {
	return coreType(c)
}

// Variable defines a variable in the builder class.
// It has a type and a variable name.
type Variable interface {
	Type
	Name() string
}

type coreType string

// Type returns a coreType as a string.
// It implements the Type interface
func (c coreType) Type() string {
	return string(c)
}

// Var constructs a new variable from a type and name
func Var(t Type, name string) Variable {
	return namedVar{t: t, name: name}
}

// List constructs a new list variable from a type and name
func List(t Type, name string) Variable {
	return listVariable{Var(t, name)}
}

// Array constructs a new array variable from a type and name
func Array(t Type, name string) Variable {
	return arrayVariable{Var(t, name)}
}

type namedVar struct {
	t    Type
	name string
}

// Name returns the name of a namedVar
// It implements the Variable interface
func (n namedVar) Name() string {
	return n.name
}

// Name returns the type of a namedVar
// It implements the Variable interface
func (n namedVar) Type() string {
	return n.t.Type()
}

type listVariable struct {
	Variable
}

// Type returns the type of a listVariable
// It implements the Type interface and
// composes a variable adding list to the inner
// variables type definition.
func (l listVariable) Type() string {
	return fmt.Sprintf("List<%s>", l.Variable.Type())
}

type arrayVariable struct {
	Variable
}

// Type returns the type of an arrayVariable
// It implements the Type interface and
// composes a variable adding list to the inner
// variables type definition.
func (l arrayVariable) Type() string {
	return fmt.Sprintf("[]%s", l.Variable.Type())
}
