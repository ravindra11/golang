package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main() {
	input := "The quick fox jumped over the lazy dog"
	rev := Reverse(input)
	fmt.Printf("Original input, %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	// fmt.Printf("reversed again, %q\n", Reverse(rev))
	reverse, err := FixReverse(rev)
	doubleRev, doubleRevErr := FixReverse(reverse)
	fmt.Printf("reversed , %q\n, err:= %q", reverse, err)
	fmt.Printf("reversed again, %q\n, err=%q", doubleRev, doubleRevErr)
}

func FixReverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not a valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {

		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
