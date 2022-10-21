package checker

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
	"path/filepath"
	c "plagChecker/pkg/parser/c"
	cpp "plagChecker/pkg/parser/cpp"
	"strings"
)

func GetTokens(file *os.File) (string, error) {
	file.Seek(0, 0)
	fs, err := antlr.NewFileStream(file.Name())
	if err != nil {
		return "", err
	}

	tokens := make([]string, 0)
	switch filepath.Ext(file.Name()) {
	case ".c":
		lexer := c.NewCLexer(fs)
		for {
			token := lexer.NextToken()
			if token.GetTokenType() == antlr.TokenEOF {
				break
			}
			tokens = append(tokens, lexer.SymbolicNames[token.GetTokenType()])
		}
	case ".cpp":
		lexer := cpp.NewCPP14Lexer(fs)
		for {
			token := lexer.NextToken()
			if token.GetTokenType() == antlr.TokenEOF {
				break
			}
			tokens = append(tokens, lexer.SymbolicNames[token.GetTokenType()])
		}
	}

	res := strings.Join(tokens, "|")

	return res, nil
}

func TokensCheck(tokens1, tokens2 string) float64 {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(tokens1, tokens2, false)

	distance := dmp.DiffLevenshtein(diffs)
	length := maxString(tokens1, tokens2)

	return (1.00 - float64(distance)/float64(length)) * 100.00
}
