package env

import (
	"errors"
	// "log"
	"reflect"
	// "strings"
)

func Decode(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.IsValid() == false {
		return errors.New("not valid value")
	}

	if v.IsNil() {
		return errors.New("can't decode nil interface")
	}

	obj := new(object)
	obj.src = i
	obj.tp = reflect.Indirect(v).Type()
	obj.value = reflect.Indirect(v)
	obj.EnvSet = Load()

	switch reflect.Indirect(v).Kind() {
	case reflect.Struct:
		decode(obj)
	default:
		return nil
	}
	return nil

}

func indirect(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr && v.CanAddr() {
		v = v.Addr()
	}

	for {
		if v.Kind() != reflect.Ptr {
			break
		}

		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		v = v.Elem()
	}

	return v
}

type object struct {
	src    interface{}
	value  reflect.Value
	tp     reflect.Type
	EnvSet EnvSet
}

func decode(obj *object) {

	v := obj.value
	tp := obj.tp
	env := obj.EnvSet

	n := tp.NumField()
	for i := 0; i < n; i++ {
		structField := tp.Field(i)

		feild := indirect(v.Field(i))

		switch feild.Kind() {
		case reflect.Int:
			tag := structField.Tag.Get("env")

			feild.SetInt(int64(env.Int(tag, 0)))
		case reflect.Bool:
			tag := structField.Tag.Get("env")

			feild.SetBool(env.Bool(tag, false))
		case reflect.String:
			tag := structField.Tag.Get("env")

			feild.SetString(env.String(tag, ""))
		case reflect.Struct:
			_obj := new(object)
			_obj.EnvSet = obj.EnvSet
			_obj.src = obj.src
			_obj.value = feild
			_obj.tp = feild.Type()
			decode(_obj)
		default:
		}
	}

}
