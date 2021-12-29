package main

import (
	"fmt"
	"strings"
)

// count
func count(s, sep string) int {
	return strings.Count(s, sep)
}

// HasPrefix
func Hasprefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// Index
func Index(s, sep string) int {
	return strings.Index(s, sep)
}

// join
func JoinStr(a []string, sep string) string {
	return strings.Join(a, sep)
}

// repeat
func RepeatStr(s string, count int) string {
	return strings.Repeat(s, count)
}

// replace
func ReplaceStr(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// split
func SplitStr(s, sep string) []string {
	return strings.Split(s, sep)
}

// Lower
func LowerStr(s string) string {
	return strings.ToLower(s)
}

// Upper
func UpperStr(s string) string {
	return strings.ToUpper(s)
}

func main() {
	// contains(s , substr string)bool
	fmt.Println(strings.Contains("test", "es"))

	// call func
	c := count("test", "t")
	fmt.Println(c)

	h := Hasprefix("test", "te")
	fmt.Println(h)

	i := Index("test", "e")
	fmt.Println(i)

	j := JoinStr([]string{"a", "b"}, "-")
	fmt.Println(j)

	r := RepeatStr("a", 5)
	fmt.Println("Repeat : ", r)

	re := ReplaceStr("aaaa", "a", "b", 2)
	fmt.Println("Replace : ", re)

	s := SplitStr("a-b-c-d-e", "-")
	fmt.Println("Split '-' :", s)
	fmt.Printf("Type : %T\n", s)

	lo := LowerStr("TEST")
	fmt.Println("TEST to lower : ", lo)

	up := UpperStr("test")
	fmt.Println("test to upper : ", up)
}
