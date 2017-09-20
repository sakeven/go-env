package env

import (
	"errors"
	"reflect"
	"strings"
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
	obj.set = LoadSet()

	if v.Kind() == reflect.Struct {
		return obj.decode()
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
	value     reflect.Value
	tp        reflect.Type
	prefixTag string
	set       Set
}

func (obj *object) decode() error {
	v := obj.value
	tp := obj.tp
	env := obj.set

	n := tp.NumField()
	for i := 0; i < n; i++ {
		structField := tp.Field(i)
		feild := indirect(v.Field(i))

		rawStructTag := structField.Tag.Get("env")
		fieldName := strings.ToUpper(structField.Name)
		tag := newStructTag(fieldName, rawStructTag).
			withPrefix(obj.prefixTag)

		// Skip this filed
		if tag.skip ||
			(feild.CanSet() == false) {
			continue
		}

		switch feild.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			defaultValue, err := tag.defaultInt64()
			if err != nil {
				return err
			}
			n := env.Int64(tag.name, defaultValue)
			feild.OverflowInt(n)
			feild.SetInt(n)
		case reflect.Bool:
			defaultValue, err := tag.defaultBool()
			if err != nil {
				return err
			}
			feild.SetBool(env.Bool(tag.name, defaultValue))
		case reflect.String:
			defaultValue, err := tag.defaultString()
			if err != nil {
				return err
			}
			feild.SetString(env.String(tag.name, defaultValue))
		case reflect.Struct:
			_obj := &object{
				set:       obj.set,
				value:     feild,
				tp:        feild.Type(),
				prefixTag: tag.name,
			}
			err := _obj.decode()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
