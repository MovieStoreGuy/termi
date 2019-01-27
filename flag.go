package termi

type Flag interface {
	SetDescription(description string)

	SetDefault(value interface{}) Flag

	GetValue() interface{}

	SetName(name string) Flag

	Set(value string) error

	IsFlag(name string) bool
}