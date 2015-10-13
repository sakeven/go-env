package env

import (
	"errors"
	"reflect"
)

// Decode parse the os environment variables and stores the result in the value pointed to by i.
// the value pointed to by i must be string, int, bool, or struct based on string, int, bool.
func Decode(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.IsValid() == false {
		return errors.New("not valid value")
	}

	v = indirect(v)

	obj := new(object)
	obj.tp = v.Type()
	obj.value = v
	obj.EnvSet = Load()

	switch v.Kind() {
	case reflect.Struct:
		obj.decode()
	default:
		return nil
	}
	return nil

}

// indirect get the real value that v points to or stores.
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
	value  reflect.Value
	tp     reflect.Type
	EnvSet EnvSet
}

func (obj *object) decode() {

	v := obj.value
	tp := obj.tp
	env := obj.EnvSet

	n := tp.NumField()
	for i := 0; i < n; i++ {
		structField := tp.Field(i)

		feild := indirect(v.Field(i))

		switch feild.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

			tag := structField.Tag.Get("env")
			n := int64(env.Int(tag, 0))
			feild.OverflowInt(n)
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
			_obj.value = feild
			_obj.tp = feild.Type()
			_obj.decode()
		default:
		}
	}

}
