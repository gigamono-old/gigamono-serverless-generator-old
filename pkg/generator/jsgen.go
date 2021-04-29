package generator

// JSServerlessGenerator compiles workflow to JavaScript code.
type JSServerlessGenerator struct{}

// Generate generates JavaScript code from workflow specification.
func (generator *JSServerlessGenerator) Generate(workflowString string, opts ...interface{}) (string, error) {
	return "", nil
}
