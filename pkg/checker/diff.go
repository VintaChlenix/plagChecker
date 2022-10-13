package checker

import (
	"crypto/md5"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io"
	"os"
)

func GetSum(file *os.File) (string, error) {
	file.Seek(0, 0)
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

func DiffCheck(norm1, norm2 string) float64 {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(norm1, norm2, false)

	distance := dmp.DiffLevenshtein(diffs)
	length := maxString(norm1, norm2)

	return (1.00 - float64(distance)/float64(length)) * 100.00
}

func maxString(x, y string) int {
	if len(x) >= len(y) {
		return len(x)
	}
	return len(y)
}
