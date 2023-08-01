// Copyright (c) 2016-2017 ByteDance, Inc. All rights reserved.
package errors

type ProgMissErr struct {
	Name string
}

func (p ProgMissErr) Error() string {
	return p.Name + " is no find."
}

type MapMissErr struct {
	Name string
}

func (p MapMissErr) Error() string {
	return p.Name + " is no find."
}

var _ error = &ProgMissErr{}
var _ error = &MapMissErr{}

func IsIgnore(err error) bool {
	switch err.(type) {
	case *ProgMissErr, *MapMissErr:
		return true
	default:
		return false
	}
}
