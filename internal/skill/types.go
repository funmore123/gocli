package skill

import "context"

// Flag 描述 Skill 的一个命令行参数
type Flag struct {
	Name     string
	Desc     string
	Default  string
	Required bool
}

// Skill 是一个声明式的快捷命令定义
// 所有 Skill 通过统一管线执行，不直接耦合 cobra
type Skill struct {
	Service     string // 挂载到哪个服务命令下（如 "demo"）
	Command     string // 命令名（如 "+greet"）
	Description string
	Flags       []Flag
	Execute     func(ctx context.Context, rt *Runtime) error
}
