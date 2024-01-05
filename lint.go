package mycustomrule

import (
	"github.com/golangci/golangci-lint/pkg/lint"
)

func init() {
	lint.RegisterLinter(&MyCustomRule{})
}
