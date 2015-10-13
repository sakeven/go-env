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
        switch structField.Type.Kind() {
        case reflect.Int:
            tag := structField.Tag.Get("env")

            v.Field(i).SetInt(int64(env.Int(tag, 0)))
        case reflect.Bool:
            tag := structField.Tag.Get("env")

            v.Field(i).SetBool(env.Bool(tag, false))
        case reflect.String:
            tag := structField.Tag.Get("env")

            v.Field(i).SetString(env.String(tag, ""))
        }
    }

}
