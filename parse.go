package env

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

var envSet map[string]string

func Load() {
    if envSet == nil {
        envSet = make(map[string]string)
    } else {
        for k, _ := range envSet {
            delete(envSet, k)
        }
    }

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

func Int(key string, defaultValue int) (value int) {
    v, ok := envSet[key]
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

func String(key string, defaultValue string) (value string) {
    v, ok := envSet[key]
    if !ok {
        value = defaultValue
        return
    }

    value = v
    return
}

func Bool(key string, defaultValue bool) (value bool) {
    v, ok := envSet[key]
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
