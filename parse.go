package env

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type EnvSet map[string]string

func Load() EnvSet {

    envSet := make(EnvSet)

    envs := os.Environ()
    for _, env := range envs {
        k, v, err := parse(env)
        if err != nil {
            continue
        }

        envSet[k] = v
    }

    return envSet
}

func load(envSet EnvSet) {
    envs := os.Environ()
    for _, env := range envs {
        k, v, err := parse(env)
        if err != nil {
            continue
        }

        envSet[k] = v
    }
}

func parse(env string) (key string, vaule string, err error) {

    splits := strings.SplitN(env, "=", 2)

    if len(splits) != 2 {
        return "", "", fmt.Errorf("parse environment %s error", env)
    }

    return splits[0], splits[1], nil
}

func (e *EnvSet) Reload() {
    if *e != nil {
        for k, _ := range *e {
            delete(*e, k)
        }
    }
    load(*e)
}

func (e EnvSet) Int(key string, defaultValue int) (value int) {
    v, ok := e[key]
    if !ok {
        value = defaultValue
        return
    }

    i, err := strconv.Atoi(v)
    if err != nil {
        panic(err)
        return
    }

    value = i
    return
}

func (e EnvSet) String(key string, defaultValue string) (value string) {
    v, ok := e[key]
    if !ok {
        value = defaultValue
        return
    }

    value = v
    return
}

func (e EnvSet) Bool(key string, defaultValue bool) (value bool) {
    v, ok := e[key]
    if !ok {
        value = defaultValue
        return
    }

    i, err := strconv.ParseBool(v)
    if err != nil {
        panic(err)
        return
    }

    value = i
    return
}
