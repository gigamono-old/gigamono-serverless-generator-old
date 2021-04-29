package compiler

// JSServerlessCompiler compiles workflow to JavaScript code.
type JSServerlessCompiler struct{}

// Compile compiles workflow to JavaScript code.
func (compiler *JSServerlessCompiler) Compile(workflowString string, opts ...interface{}) (string, error) {
	return "", nil
}
