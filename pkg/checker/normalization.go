package checker

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Normalize(file *os.File) (string, error) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	re1 := regexp.MustCompile("//.*|/\\*.*|\\{|\\}|#.*")
	re2 := regexp.MustCompile(`\s+`)
	var res string
	for scanner.Scan() {
		strings.ToLower(scanner.Text())
		tmp := re1.ReplaceAllString(scanner.Text(), "")
		tmp = re2.ReplaceAllString(tmp, " ")
		res += tmp
	}

	return res, nil
}
