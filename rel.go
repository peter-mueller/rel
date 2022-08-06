package rel

import "reflect"

type Var[T any] struct {
	attributeNames []string
	typ            reflect.Type
}

func NewVar[T any]() (v *Var[T]) {
	v = new(Var[T])

	var typ *T
	v.typ = reflect.TypeOf(typ).Elem()
	v.attributeNames = make([]string, v.typ.NumField())
	for i := 0; i < v.typ.NumField(); i++ {
		v.attributeNames[i] = v.typ.Field(i).Tag.Get("db")
	}

	return v
}

func (v *Var[T]) Tuple(value T) (tuple []any) {
	tuple = make([]any, v.typ.NumField())
	r := reflect.ValueOf(value)
	for i := 0; i < v.typ.NumField(); i++ {
		tuple[i] = r.Field(i).Interface()
	}
	return tuple
}

func (v *Var[T]) Dests(value *T) (dests []any) {
	dests = make([]any, v.typ.NumField())
	r := reflect.ValueOf(value).Elem()
	for i := 0; i < v.typ.NumField(); i++ {
		dests[i] = r.Field(i).Addr().Interface()
	}
	return dests
}

func (v *Var[T]) AttributeRef() (attributeRefs []string) {
	return v.attributeNames
}

func (v *Var[T]) Params(paramToken string) (params []string) {
	params = make([]string, v.typ.NumField())
	for i := 0; i < v.typ.NumField(); i++ {
		params[i] = paramToken
	}
	return params
}
