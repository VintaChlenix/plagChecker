package checker

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Normalize(file *os.File) (string, error) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[[:space:]]|//.*|/\\*.*")
	var res string
	for scanner.Scan() {
		strings.ToLower(scanner.Text())
		res += re.ReplaceAllString(scanner.Text(), "")
	}

	return res, nil
}
