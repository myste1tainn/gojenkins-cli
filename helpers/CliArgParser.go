package helpers

import (
	"github.com/jessevdk/go-flags"
	"github.com/myste1tainn/gojenkins-cli/commands"
	gojenkins_core "github.com/myste1tainn/gojenkins-core/wire"
)

func NewCliArgParser() *flags.Parser {
	var parser = flags.NewParser(&struct{}{}, flags.Default)
	_, err := parser.AddCommand(
		"build-with-parameters",
		"BWP Short Description",
		"BWP Long Description",
		&commands.DefaultBuildWithParametersCommand{gojenkins_core.InitializeJobController()},
	)
	if err != nil {
		panic(err)
	}
	return parser
}
