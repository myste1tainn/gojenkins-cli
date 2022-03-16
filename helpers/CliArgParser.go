package helpers

type CliArgParser interface {
	Parse()
}

type DefaultCliArgParser struct {
}

func (this DefaultCliArgParser) Parse() {

}

func NewCliArgParser() *DefaultCliArgParser {
	return &DefaultCliArgParser{}
}
