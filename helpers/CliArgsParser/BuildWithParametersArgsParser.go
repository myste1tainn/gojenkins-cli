package CliArgsParser

import "github.com/jessevdk/go-flags"

var command = flags.Command{
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
	UnknownOptionHandler:  nil,
	CompletionHandler:     nil,
	CommandHandler:        nil,
}
