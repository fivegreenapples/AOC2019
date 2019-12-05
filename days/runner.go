package days

import "fmt"

import "reflect"

type Runner struct {
	verbose bool
}

func NewRunner(verboseMode bool) *Runner {
	return &Runner{
		verbose: verboseMode,
	}
}

func (r *Runner) Run(day int, part int, input string) (string, error) {

	// use reflection to determine if we have a suitable method
	methodName := fmt.Sprintf("Day%dPart%d", day, part)

	rValue := reflect.ValueOf(r)
	rMethod := rValue.MethodByName(methodName)
	if !rMethod.IsValid() {
		return "", fmt.Errorf("implementation for day %d, part %d is not available", day, part)
	}

	rResult := rMethod.Call([]reflect.Value{reflect.ValueOf(input)})
	// rResult should represent a slice of a single string
	return rResult[0].String(), nil
}
