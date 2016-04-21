// The MIT License (MIT)
//
// Copyright (c) 2015 Chong Jiang

/*
	go-env is a go library for using environment variable to configure your project
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
*/
package env
