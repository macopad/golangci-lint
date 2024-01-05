package mycustomrule

import (
	"github.com/macopad/golangci-lint/lint"
)

type MyCustomRule struct {
	// 可以在这里定义规则需要的任何配置参数或状态
}

func (r *MyCustomRule) Name() string {
	return "mycustomrule"
}

func (r *MyCustomRule) Description() string {
	return "This is a custom lint rule."
}

func (r *MyCustomRule) Apply(file *lint.File, _ lint.Arguments) ([]lint.Issue, error) {
	// 在这里实现规则的逻辑，检查文件并返回问题列表
}

func (r *MyCustomRule) DefaultEnabled() bool {
	return true
}
