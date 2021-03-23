package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abc", "ade"))
	fmt.Println(strings.Contains("abc", "ab"))
	fmt.Println(strings.Count("abcabcab", "ab"))
	fmt.Println(strings.EqualFold("ab", "AB"))
	fmt.Printf("%q\n", strings.Fields("a b\tc\r\nd\fe"))
	fmt.Println(strings.HasPrefix("abcdef.go", "abc"))
	fmt.Println(strings.HasSuffix("abcdef.go", ".exe"))
	fmt.Println(strings.Index("1abcabc", "ab"))
	fmt.Println(strings.LastIndex("1abcabc", "ab"))
	fmt.Println(strings.Index("1abcabc", "abd"))
	fmt.Println(strings.LastIndex("1abcabc", "abd"))

	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Printf("%q\n", strings.Split("a-b-c", "-"))
	fmt.Printf("%q\n", strings.SplitN("a-b-c", "-", 2))
	fmt.Printf("%q\n", strings.SplitN("abc", "-", 2))
	fmt.Printf("%q\n", strings.Repeat("*", 20))

	fmt.Println(strings.Replace("abc,abc,abcd,abd", "abc", "xyz", 2))
	fmt.Println(strings.Replace("abc,abc,abcd,abd", "abc", "xyz", -1))
	fmt.Println(strings.ReplaceAll("abc,abc,abcd,abd", "abc", "xyz"))
	fmt.Println(strings.Title("abc Def gH"))
	fmt.Println(strings.ToLower("abc Def gH"))
	fmt.Println(strings.ToUpper("abc Def gH"))

	fmt.Println(strings.Trim("acbAbcdeacb", "abc"))
	fmt.Println(strings.TrimLeft("acbAbcdeacb", "abc"))
	fmt.Println(strings.TrimRight("acbAbcdeacb", "abc"))
	fmt.Println(strings.TrimSuffix("acbAbcdeacb", "abc"))
	fmt.Println(strings.TrimPrefix("acbAbcdeacb", "abc"))
	fmt.Println(strings.TrimSuffix("abcacbAbcdeacbabc", "abc"))
	fmt.Println(strings.TrimPrefix("abcacbAbcdeacbabc", "abc"))
	fmt.Printf("%q\n", strings.TrimSpace(" abc abc\r\n\f"))
}
