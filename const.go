package termi

import (
	"errors"
)

const (
	EnvironmentTag = "env"
	DescriptionTag = "description"
	DefaultDescription = "expects %v as the value"
)

var (
	ErrInvalidType = errors.New("incorrect type being used, expected struct or pointer to struct")
)