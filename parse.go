package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Set stores all environment variables <key, valuer> pair.
type Set map[string]string

// LoadSet load environment variables to parse a new EnvSet.
func LoadSet() Set {
	envSet := make(Set)
	loadEnvSet(envSet)
	return envSet
}

func loadEnvSet(envSet Set) {
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
func (e *Set) Reload() {
	if *e != nil {
		for k := range *e {
			delete(*e, k)
		}
	}
	loadEnvSet(*e)
}

// Int bind int value of specific key.
func (e Set) Int(key string, defaultValue int) int {
	return int(e.Int64(key, int64(defaultValue)))
}

// Int64 bind int value of specific key.
func (e Set) Int64(key string, defaultValue int64) int64 {
	v, ok := e[key]
	if !ok {
		return defaultValue
	}

	realValue, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return realValue
}

// String bind string value of specific key.
func (e Set) String(key string, defaultValue string) string {
	realValue, ok := e[key]
	if !ok {
		return defaultValue
	}
	return realValue
}

// Bool bind bool value of specific key.
func (e Set) Bool(key string, defaultValue bool) bool {
	v, ok := e[key]
	if !ok {
		return defaultValue
	}

	realValue, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return realValue
}
