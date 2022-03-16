package CliArgsParser

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

var group = flags.Group{
	ShortDescription: "Group Short Description",
	LongDescription:  "Group Long Description",
	Namespace:        "Group namespace",
	EnvNamespace:     "Group EnvNamespace",
	Hidden:           false,
}

var command = flags.Command{
	Group:               &group,
	Name:                "build-with-parameters",
	Active:              nil,
	SubcommandsOptional: false,
	Aliases:             []string{"bwp"},
	ArgsRequired:        true,
}

var BuildWithParametersArgsParser = flags.Parser{
	Command:               &command,
	Usage:                 "For the specify URL path run the job as `buildWithParameter` with the specify parameters",
	Options:               0,
	NamespaceDelimiter:    "",
	EnvNamespaceDelimiter: "",
	UnknownOptionHandler: func(option string, arg flags.SplitArgument, args []string) ([]string, error) {
		fmt.Println("UnknownOptionHandler", option, arg, args)
		return args, nil
	},
	CompletionHandler: func(items []flags.Completion) {
		fmt.Println("CompletionHandler", items)
	},
	CommandHandler: func(command flags.Commander, args []string) error {
		fmt.Println("CommandHandler", command, args)
		return nil
	},
}

func Asdf() {
}
