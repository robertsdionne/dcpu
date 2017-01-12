package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/robertsdionne/dcpu/parser"
)

func main() {

	text, _ := ioutil.ReadAll(os.Stdin)
	input := antlr.NewInputStream(string(text))
	lexer := parser.NewDCPULexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDCPUParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	stream.Fill()
	fmt.Println(stream.GetAllTokens())
	fmt.Println()
	tree := p.Program()
	fmt.Println(tree.ToStringTree(p.RuleNames, p))

}
