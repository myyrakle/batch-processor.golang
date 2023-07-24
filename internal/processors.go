package internal

type Processor interface {
	run()
}

var Processors map[string]Processor = map[string]Processor{}
