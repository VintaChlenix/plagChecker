package checker

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func GetSum(file *os.File) (string, error) {
	sum := md5.New()
	_, err := io.Copy(sum, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", sum.Sum(nil)), nil
}

func SumCheck(sum1, sum2 string) bool {
	return sum1 == sum2
}

func DiffCheck(norm1, norm2 string, window int) float64 {
	var matches float64
	var mismatches float64
	n := min(norm1, norm2)
	for i := range n {
		if i+window >= len(n) {
			break
		}
		if norm1[i:i+window] == norm2[i:i+window] {
			matches++
		}
		if norm1[i:i+window] != norm2[i:i+window] {
			mismatches++
		}
	}
	return matches / (matches + mismatches)
}

func min(x, y string) string {
	if len(x) <= len(y) {
		return x
	}
	return y
}
