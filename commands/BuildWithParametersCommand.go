package commands

import (
	"github.com/myste1tainn/gojenkins-cli/models/requests"
	"github.com/myste1tainn/gojenkins-core/controllers"
	coreRequests "github.com/myste1tainn/gojenkins-core/models/requests"
)

type BuildWithParametersCommand interface {
	Execute(req requests.BuildWithParametersCommandRequest)
}

type DefaultBuildWithParametersCommand struct {
	JobController controllers.JobController
}

func (this DefaultBuildWithParametersCommand) Execute(jobUrl string, params map[string]string) {
	this.JobController.BuildWithParameters(coreRequests.BuildWithParametersRequest{
		JobUrl: jobUrl,
		Params: params,
	})
}
