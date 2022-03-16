package commands

type ConsoleTextCommand interface {
	Execute(jobId string)
}

type DefaultConsoleTextCommand struct {
}

func (this DefaultConsoleTextCommand) Execute(params map[string]string) {

}
