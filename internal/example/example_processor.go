package example

import "fmt"

type ExampleProcessor struct {
}

func (p *ExampleProcessor) Run() {
	fmt.Println("Example Processor")
}
