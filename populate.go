package env

import (
	"errors"
	// "log"
	"reflect"
	"strconv"
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
	obj.EnvSet = LoadEnvSet()

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
	value     reflect.Value
	tp        reflect.Type
	prefixTag string
	EnvSet    EnvSet
}

func (obj *object) decode() {
	v := obj.value
	tp := obj.tp
	env := obj.EnvSet

	n := tp.NumField()
	for i := 0; i < n; i++ {
		structField := tp.Field(i)
		feild := indirect(v.Field(i))

		rawStructTag := structField.Tag.Get("env")
		tag := structTag{Name: strings.ToUpper(structField.Name)}
		tag.parseTag(obj.prefixTag, rawStructTag)
		// log.Println(rawStructTag, structField.Name, tag.Name)

		if tag.Skip {
			continue
		}

		switch feild.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			defaultValue := int64(0)
			if tag.Omitempty != true {
				var err error
				defaultValue, err = strconv.ParseInt(tag.Default, 10, 64)
				if err != nil {
					break
				}
			}

			n := env.Int64(tag.Name, defaultValue)
			feild.OverflowInt(n)
			feild.SetInt(n)
		case reflect.Bool:
			defaultValue := false
			if tag.Omitempty != true {
				var err error
				defaultValue, err = strconv.ParseBool(tag.Default)
				if err != nil {
					break
				}
			}
			feild.SetBool(env.Bool(tag.Name, defaultValue))
		case reflect.String:
			defaultValue := ""
			if tag.Omitempty != true {
				defaultValue = tag.Default
			}

			feild.SetString(env.String(tag.Name, defaultValue))
		case reflect.Struct:
			_obj := new(object)
			_obj.EnvSet = obj.EnvSet
			_obj.value = feild
			_obj.tp = feild.Type()
			_obj.prefixTag = tag.Name

			_obj.decode()
		default:
		}
	}

}

type structTag struct {
	Name      string
	Omitempty bool
	Skip      bool
	Default   string
}

func (t *structTag) parseTag(prefixTag, rawStructTag string) {
	list := strings.Split(rawStructTag, ",")

	var options [2]string

	for i, op := range list {
		options[i] = fix(op)
	}

	//  tag name
	switch options[0] {
	case "-":
		t.Skip = true
	case "":
		// use origin field name
	default:
		t.Name = options[0]
	}

	if len(prefixTag) > 0 {
		t.Name = prefixTag + "_" + t.Name
	}

	// tag default value
	switch options[1] {
	case "":
		// use omit empty value
		t.Omitempty = true
	default:
		t.Default = options[1]
	}
}

func fix(s string) string {
	s = strings.Trim(s, " ")
	s = strings.Trim(s, "\t")

	return s
}
