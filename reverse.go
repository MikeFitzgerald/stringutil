// Package stringutil contains utility functions for strings.
package stringutil

// Reverse returns a string reversed rune-wise left to right.
func Reverse(s string) string {
    r := []rune(s)

    // [0] [1] [2] [3] [4]
    //  F   A   R   T   \0
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}