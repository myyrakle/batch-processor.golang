package internal

import (
	example "batch/internal/example"
)

type Processor interface {
	Run()
}

var Processors map[string]Processor = map[string]Processor{
	"EXAMPLE": &example.ExampleProcessor{},
}
