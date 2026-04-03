package skill

import (
	"github.com/example/gocli/internal/config"
	"github.com/example/gocli/internal/output"
	"github.com/spf13/cobra"
)

// Runtime 是 Skill 执行时的上下文，提供 Flag 访问、配置读取、输出能力
type Runtime struct {
	cmd    *cobra.Command
	Config *config.Config
}

// Str 读取 string 类型 Flag
func (r *Runtime) Str(name string) string {
	v, _ := r.cmd.Flags().GetString(name)
	return v
}

// Bool 读取 bool 类型 Flag
func (r *Runtime) Bool(name string) bool {
	v, _ := r.cmd.Flags().GetBool(name)
	return v
}

// Out 以 JSON 格式输出数据
func (r *Runtime) Out(data interface{}) error {
	return output.JSON(data)
}

func newRuntime(cmd *cobra.Command) (*Runtime, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return &Runtime{cmd: cmd, Config: cfg}, nil
}
