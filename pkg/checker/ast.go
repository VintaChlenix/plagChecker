package checker

import (
	parser "plagChecker/pkg/parser/c"
)

type TreeShapeListener struct {
	*parser.BaseCListener
}

/*func GetAST(file *os.File) {
	input, err := antlr.NewFileStream(file.Name())
	if err != nil{
		return
	}

	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&TreeShapeListener{}, ..)
}*/

func ASTCheck() float64 {
	return float64(0)
}
