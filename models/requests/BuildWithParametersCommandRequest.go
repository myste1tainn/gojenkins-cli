package requests

type BuildWithParametersCommandRequest struct {
	JobUrl string
	Params map[string]string
}
