package generator

// JSServerlessGenerator compiles workflow to JavaScript code.
type JSServerlessGenerator struct{}

// Generate generates JavaScript code from workflow specification.
func (compiler *JSServerlessGenerator) Generate(workflowString string, opts ...interface{}) (string, error) {
	return "", nil
}
