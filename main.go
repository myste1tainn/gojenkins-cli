package main

import "github.com/myste1tainn/gojenkins-cli/helpers"

func main() {
	// TODO:1 Now to make the parser have the ability to "map command to jobUrl"
	// TODO:10 Maybe add the ability to "map positional arguments to Params with fixed name"
	parser := helpers.NewMainArgsParser()
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	// TODO:2 Add the ability to "core" to print out consoleText after execute a job also
	//   and make sure it can "recognized FINISH or ERROR" of a job
	// TODO:3 Add ability to send teams message after it finishes / error
}
