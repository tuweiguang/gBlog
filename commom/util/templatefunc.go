package util

import (
	"strconv"
	"strings"
	"time"
)

func IndexForOne(i int, p, limit int) int {
	s := strconv.Itoa(i)
	index, _ := strconv.ParseInt(s, 10, 0)
	return (p-1)*limit + int(index) + 1
}

func IndexDecrOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index - 1
}

func IndexAddOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index + 1
}

func Dateformat(t time.Time, format string) string {
	return t.Format(format)
}

func StringReplace(s, old, new string) string {
	return strings.Replace(s, old, new, 1)
}
