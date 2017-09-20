// The MIT License (MIT)
//
// Copyright (c) 2015 Chong Jiang

/*
Package env is a go library for using environment variable to configure your project
Usage:

	// Field appears in env as key "FEILD" and
	// the field is omitted from the object if its value is empty,
	// as defined
	Field int `env:"FEILD"`
	// Field appears in env as key "FEILD" and
	// the field is omitted 1 if its value is empty.
	Field int `env:"FEILD,1"`
	// Field is ignored by this package.
	Field int `env:"-"`
	// The upper key name will be used if it's a non-empty string consisting of
	// only Unicode letters, digits, dollar signs, percent signs, hyphens, underscores and slashes.
	Field int `env:",1"`
*/
package env
