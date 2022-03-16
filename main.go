package main

import "github.com/myste1tainn/gojenkins-cli/helpers"

func main() {
	parser := helpers.NewCliArgParser()
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}
}
