package builder

import (
	"fmt"
	"html/template"
)

const (
	String  coreType = "String"
	Integer coreType = "Integer"
	Int     coreType = "int"
	Bool    coreType = "boolean"
	Boolean coreType = "Boolean"
)

type Type interface {
	Type() template.HTML
}

func Custom(c string) Type {
	return coreType(c)
}

type Variable interface {
	Type
	Name() string
}

type coreType string

func (c coreType) Type() template.HTML {
	return template.HTML(c)
}

func Var(t Type, name string) Variable {
	return namedVar{t: t, name: name}
}

func List(t Type, name string) Variable {
	return listVariable{Var(t, name)}
}

func Array(t Type, name string) Variable {
	return arrayVariable{Var(t, name)}
}

type namedVar struct {
	t    Type
	name string
}

func (n namedVar) Name() string {
	return n.name
}

func (n namedVar) Type() template.HTML {
	return n.t.Type()
}

func NewList(v Variable) Variable {
	return listVariable{v}
}

type listVariable struct {
	Variable
}

func (l listVariable) Type() template.HTML {
	t := fmt.Sprintf("List<%s>", l.Variable.Type())
	return template.HTML(t)
}

type arrayVariable struct {
	Variable
}

func (l arrayVariable) Type() template.HTML {
	t := fmt.Sprintf("[]%s", l.Variable.Type())
	return template.HTML(t)
}
