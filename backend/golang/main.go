package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func removeNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "\r\n") {
		return removeNewLineSuffixes(s[:len(s)-2])
	}
	if strings.HasSuffix(s, "\n") {
		return removeNewLineSuffixes(s[:len(s)-1])
	}
	return s
}
