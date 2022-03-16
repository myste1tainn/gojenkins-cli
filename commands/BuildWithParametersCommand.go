package commands

import (
	"errors"
	"fmt"
	"github.com/myste1tainn/gojenkins-cli/models/requests"
	"github.com/myste1tainn/gojenkins-core/controllers"
	coreRequests "github.com/myste1tainn/gojenkins-core/models/requests"
	"strings"
)

type BuildWithParametersCommand interface {
	execute(req requests.BuildWithParametersCommandRequest)
}

type DefaultBuildWithParametersCommand struct {
	JobController controllers.JobController
}

func (this *DefaultBuildWithParametersCommand) execute(jobUrl string, params map[string]string) {
	fmt.Printf("build-with-parameters: jobUrl:%s params:%v\n", jobUrl, params)
	this.JobController.BuildWithParameters(coreRequests.BuildWithParametersRequest{
		JobUrl: jobUrl,
		Params: params,
	})
}

func (this DefaultBuildWithParametersCommand) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("build-with-parameters command requires at least 2 arguments, provided: `" + strings.Join(args, ",") + "`")
	}
	params := map[string]string{}
	strs := strings.Split(args[1], ",")
	for _, str := range strs {
		vals := strings.Split(str, "=")
		if len(vals) > 1 {
			params[vals[0]] = vals[1]
		}
	}
	this.execute(args[0], params)
	return nil
}
