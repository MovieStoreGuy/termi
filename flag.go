package termi

// Flag is a neat little abstraction that allows for simplified
// Command line arguments
type Flag interface {
	// SetDescription allows you to set the usage of the flag
	SetDescription(description string) Flag

	// SetValue allows you to pass the reference to the value you wish to
	// update with when parse is called.
	SetValue(value interface{}) Flag

	// SetName allows you to define what argument is associated with this flag.
	// Can be called multiple times
	SetName(name string) Flag

	// Set is called when parsing the argument within the FlagSet
	Set(value string) error

	// IsFlag allows the flag set to parse the command line arg
	IsFlag(name string) bool
}
