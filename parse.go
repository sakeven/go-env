package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvSet map[string]string

// LoadEnvSet load environment variable to parse a new EnvSet.
func LoadEnvSet() EnvSet {
	envSet := make(EnvSet)

	envs := os.Environ()
	for _, env := range envs {
		k, v, err := parseEnv(env)
		if err != nil {
			continue
		}

		envSet[k] = v
	}

	return envSet
}

func loadEnvSet(envSet EnvSet) {
	envs := os.Environ()
	for _, env := range envs {
		k, v, err := parseEnv(env)
		if err != nil {
			continue
		}

		envSet[k] = v
	}
}

func parseEnv(env string) (key string, vaule string, err error) {

	splits := strings.SplitN(env, "=", 2)

	if len(splits) != 2 {
		return "", "", fmt.Errorf("parse environment %s error", env)
	}

	return splits[0], splits[1], nil
}

// Reload refresh envSet from environment.
func (e *EnvSet) Reload() {
	if *e != nil {
		for k, _ := range *e {
			delete(*e, k)
		}
	}
	loadEnvSet(*e)
}

// Int bind int value of specific key.
func (e EnvSet) Int(key string, defaultValue int) (value int) {
	return int(e.Int64(key, int64(defaultValue)))
}

// Int bind int value of specific key.
func (e EnvSet) Int64(key string, defaultValue int64) (value int64) {
	v, ok := e[key]
	if !ok {
		value = defaultValue
		return
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
		return
	}

	value = i
	return
}

// String bind string value of specific key.
func (e EnvSet) String(key string, defaultValue string) (value string) {
	v, ok := e[key]
	if !ok {
		value = defaultValue
		return
	}

	value = v
	return
}

// Bool bind bool value of specific key.
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
