package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {

// 	testCases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}

// 	for _, tc := range testCases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }

/* Now will change the unit test into a fuzz test */
/* Add a fuzz test */

/*
1. The unit test has limitations,namely that each input must be added to the test by the developer.
2. One benefit of fuzzing is that it comes up with inputs for your code,
and may identify edge cases that the test cases you came up with didn't reach.

3. In this section you will convert the unit test to a fuzz test so that  you can generate more inputs with less work.

4. Note that, you can keep unit tests, benchmarks, and fuzz tests in the same `*_test.go` file.

*/

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {

		// rev := Reverse(orig)
		// doubleRev := Reverse(rev)
		rev, err := FixReverse(orig)
		if err != nil {
			return
		}
		doubleRev, doubleErr := FixReverse(rev)
		if doubleErr != nil {
			return
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

/*
Fuzzing has a few limitations as well.
In your unit test, you could predict the expected output of the Reverse function, and verify that the actual output met those expectations.

For example, in the testcase Reverse("Hello, world") the unit test specifies the return as "dlrow ,olleH".

When fuzzing, you can't predict the expected output, since you don't have control over the inputs.

 However, there are a few properties of the Reverse function that you can verify in a funzz test. The two properties being checked in this funzz test are:

  1. Reversing a string twice preserves the original value.
  2. The reversed string preserves its state as valid UTF-8

Note the syntax differences between the unit test and the fuzz test:

1. The function begins with FuzzXxx instead of TestXxx, and takes *testing.F instead of *testing.T
2. Where you would expect to see a t.Run execution, you instead see f.Fuzz which takes a fuzz target function whose
parameters are *testing.T and the types to be fuzzed. The inputs from your unit test are provided as seed corpus inputs using f.Add
*/

/*
logs

    reverse_test.go:50: Number of runes: orig=1, rev=2, doubleRev=1
        reverse_test.go:55: Reverse produced invalid UTF-8 string "\xb2\xdf"


The entire seed corpus used strings in which every character was a single byte.

However, characters such as symbols can require several bytes. Thus, reversing the string
byte-by-byte will invalidate multi-byte characters.


with a better understanding of the bug, correct the error in the Reverse function.

*/

/* Fix the error */

// to correct the Reverse function, let's traverse the string by runes,
// instead of by bytes
