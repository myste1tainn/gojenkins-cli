package main

import (
	"github.com/myste1tainn/gojenkins-cli/helpers"
)

func main() {
	parser := helpers.NewCliArgParser()
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	//fmt.Println("#### Parser", CliArgsParser.BuildWithParametersArgsParser)
	//args, err := CliArgsParser.BuildWithParametersArgsParser.Parse()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("args received", args)
}
