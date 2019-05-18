// Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved.

package main

import (
	"strconv"
	"unicode"
)

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func arrAtoi(s []string) (t []int, err error) {
	var j int

	for _, i := range s {
		if j, err = strconv.Atoi(i); err != nil {
			return
		}
		t = append(t, j)
	}
	return
}
