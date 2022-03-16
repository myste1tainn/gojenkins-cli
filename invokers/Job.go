package invokers

import (
	"github.com/myste1tainn/gojenkins-cli/commands"
	"github.com/myste1tainn/gojenkins-cli/models/requests"
)

type JobInvoker interface {
	BuildWithParameter(req requests.BuildWithParametersRequest)
	ConsoleText(req requests.ConsoleTextRequest)
}

type DefaultJobInvoker struct {
	BuildWithParameterCommand commands.BuildWithParametersCommand
	ConsoleTextCommand        commands.ConsoleTextCommand
}

func (this DefaultJobInvoker) BuildWithParameter(req requests.BuildWithParametersRequest) {
	this.BuildWithParameterCommand.Execute(requests.BuildWithParametersCommandRequest(req))
}

func (this DefaultJobInvoker) ConsoleText(req requests.ConsoleTextRequest) {
	this.ConsoleTextCommand.Execute(req.JobId)
}
