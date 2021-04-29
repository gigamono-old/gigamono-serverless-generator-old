package generator

// RustServerlessGenerator generates Rust code from workflow specification.
type RustServerlessGenerator struct{}

// Generate generates Rust code from workflow specification.
func (generator *RustServerlessGenerator) Generate(workflowString string, opts ...interface{}) (string, error) {
	return "", nil
}
